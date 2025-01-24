package password

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/authgear/authgear-server/pkg/lib/config"

	"github.com/trustelem/zxcvbn"

	"github.com/authgear/authgear-server/pkg/api/apierrors"
	corepassword "github.com/authgear/authgear-server/pkg/util/password"
)

func isUpperRune(r rune) bool {
	// unicode.IsUpper is not used intentionally because it takes other languages into account.
	return r >= 'A' && r <= 'Z'
}

func isLowerRune(r rune) bool {
	// unicode.IsLower is not used intentionally because it takes other languages into account.
	return r >= 'a' && r <= 'z'
}

func isDigitRune(r rune) bool {
	// unicode.IsDigit is not used intentionally because it takes other languages into account.
	return r >= '0' && r <= '9'
}

func isSymbolRune(r rune) bool {
	// We define symbol as non-alphanumeric character
	return !isUpperRune(r) && !isLowerRune(r) && !isDigitRune(r)
}

func checkPasswordLength(password string, minLength int) bool {
	if minLength <= 0 {
		return true
	}
	// Count number of code points.
	return len([]rune(password)) >= minLength
}

func checkPasswordUppercase(password string) bool {
	for _, r := range password {
		if isUpperRune(r) {
			return true
		}
	}
	return false
}

func checkPasswordLowercase(password string) bool {
	for _, r := range password {
		if isLowerRune(r) {
			return true
		}
	}
	return false
}

func checkPasswordAlphabet(password string) bool {
	for _, r := range password {
		if isLowerRune(r) || isUpperRune(r) {
			return true
		}
	}
	return false
}

func checkPasswordDigit(password string) bool {
	for _, r := range password {
		if isDigitRune(r) {
			return true
		}
	}
	return false
}

func checkPasswordSymbol(password string) bool {
	for _, r := range password {
		if isSymbolRune(r) {
			return true
		}
	}
	return false
}

func checkPasswordExcludedKeywords(password string, keywords []string) bool {
	if len(keywords) <= 0 {
		return true
	}
	words := []string{}
	for _, w := range keywords {
		words = append(words, regexp.QuoteMeta(w))
	}
	re, err := regexp.Compile("(?i)" + strings.Join(words, "|"))
	if err != nil {
		return false
	}
	loc := re.FindStringIndex(password)
	return loc == nil
}

func checkPasswordGuessableLevel(password string, minLevel int) (int, bool) {
	if minLevel <= 0 {
		return 0, true
	}
	minScore := minLevel - 1
	if minScore > 4 {
		minScore = 4
	}
	result := zxcvbn.PasswordStrength(password, nil)
	ok := result.Score >= minScore
	return result.Score + 1, ok
}

type CheckerHistoryStore interface {
	GetPasswordHistory(ctx context.Context, userID string, historySize int, historyDays config.DurationDays) ([]History, error)
}

type Checker struct {
	PwMinLength            int
	PwUppercaseRequired    bool
	PwLowercaseRequired    bool
	PwAlphabetRequired     bool
	PwDigitRequired        bool
	PwSymbolRequired       bool
	PwMinGuessableLevel    int
	PwExcludedKeywords     []string
	PwHistorySize          int
	PwHistoryDays          config.DurationDays
	PasswordHistoryEnabled bool
	PasswordHistoryStore   CheckerHistoryStore
}

func (pc *Checker) policyPasswordLength() Policy {
	return Policy{
		Name: PasswordTooShort,
		Info: map[string]interface{}{
			"min_length": pc.PwMinLength,
		},
	}
}

func (pc *Checker) checkPasswordLength(password string) *Policy {
	v := pc.policyPasswordLength()
	minLength := pc.PwMinLength
	if minLength > 0 && !checkPasswordLength(password, minLength) {
		v.Info["pw_length"] = len(password)
		return &v
	}
	return nil
}

func (pc *Checker) checkPasswordUppercase(password string) *Policy {
	if pc.PwUppercaseRequired && !checkPasswordUppercase(password) {
		return &Policy{Name: PasswordUppercaseRequired}
	}
	return nil
}

func (pc *Checker) checkPasswordLowercase(password string) *Policy {
	if pc.PwLowercaseRequired && !checkPasswordLowercase(password) {
		return &Policy{Name: PasswordLowercaseRequired}
	}
	return nil
}

func (pc *Checker) checkPasswordAlphabet(password string) *Policy {
	if pc.PwAlphabetRequired && !checkPasswordAlphabet(password) {
		return &Policy{Name: PasswordAlphabetRequired}
	}
	return nil
}

func (pc *Checker) checkPasswordDigit(password string) *Policy {
	if pc.PwDigitRequired && !checkPasswordDigit(password) {
		return &Policy{Name: PasswordDigitRequired}
	}
	return nil
}

func (pc *Checker) checkPasswordSymbol(password string) *Policy {
	if pc.PwSymbolRequired && !checkPasswordSymbol(password) {
		return &Policy{Name: PasswordSymbolRequired}
	}
	return nil
}

func (pc *Checker) checkPasswordExcludedKeywords(password string) *Policy {
	keywords := pc.PwExcludedKeywords
	if len(keywords) > 0 && !checkPasswordExcludedKeywords(password, keywords) {
		return &Policy{Name: PasswordContainingExcludedKeywords}
	}
	return nil
}

func (pc *Checker) policyPasswordGuessableLevel() Policy {
	return Policy{
		Name: PasswordBelowGuessableLevel,
		Info: map[string]interface{}{
			"min_level": pc.PwMinGuessableLevel,
		},
	}
}

func (pc *Checker) checkPasswordGuessableLevel(password string) *Policy {
	v := pc.policyPasswordGuessableLevel()
	minLevel := pc.PwMinGuessableLevel
	if minLevel > 0 {
		level, ok := checkPasswordGuessableLevel(password, minLevel)
		if !ok {
			v.Info["pw_level"] = level
			return &v
		}
	}
	return nil
}

func (pc *Checker) policyPasswordHistory() Policy {
	return Policy{
		Name: PasswordReused,
		Info: map[string]interface{}{
			"history_size": pc.PwHistorySize,
			"history_days": int(pc.PwHistoryDays),
		},
	}
}

func (pc *Checker) checkPasswordHistory(ctx context.Context, password, authID string) (*Policy, error) {
	v := pc.policyPasswordHistory()
	if pc.shouldCheckPasswordHistory() && authID != "" {
		history, err := pc.PasswordHistoryStore.GetPasswordHistory(
			ctx,
			authID,
			pc.PwHistorySize,
			pc.PwHistoryDays,
		)
		if err != nil {
			return nil, err
		}
		for _, ph := range history {
			if IsSamePassword(ph.HashedPassword, password) {
				return &v, nil
			}
		}
	}
	return nil, nil
}

func (pc *Checker) checkCommonPolicies(plainPassword string) []apierrors.Cause {
	var violations []apierrors.Cause
	check := func(v *Policy) {
		if v != nil {
			violations = append(violations, *v)
		}
	}

	check(pc.checkPasswordLength(plainPassword))
	check(pc.checkPasswordUppercase(plainPassword))
	check(pc.checkPasswordLowercase(plainPassword))
	check(pc.checkPasswordAlphabet(plainPassword))
	check(pc.checkPasswordDigit(plainPassword))
	check(pc.checkPasswordSymbol(plainPassword))
	check(pc.checkPasswordExcludedKeywords(plainPassword))
	check(pc.checkPasswordGuessableLevel(plainPassword))

	return violations
}

// ValidateNewPassword should be used when the user changes their password.
func (pc *Checker) ValidateNewPassword(ctx context.Context, userID string, plainPassword string) error {
	var violations []apierrors.Cause

	violations = pc.checkCommonPolicies(plainPassword)

	p, err := pc.checkPasswordHistory(ctx, plainPassword, userID)
	if err != nil {
		return err
	}
	if p != nil {
		violations = append(violations, p)
	}

	if len(violations) == 0 {
		return nil
	}

	return PasswordPolicyViolated.NewWithCauses("password policy violated", violations)
}

// ValidateCurrentPassword should be used when the user authenticates.
func (pc *Checker) ValidateCurrentPassword(plainPassword string) error {
	var violations []apierrors.Cause = pc.checkCommonPolicies(plainPassword)

	if len(violations) == 0 {
		return nil
	}

	return PasswordPolicyViolated.NewWithCauses("password policy violated", violations)
}

func (pc *Checker) PasswordRules() string {
	var builder strings.Builder

	if pc.PwMinLength > 0 {
		builder.WriteString(fmt.Sprintf("minlength: %v; ", pc.PwMinLength))
	}
	if pc.PwUppercaseRequired {
		builder.WriteString("required: upper; ")
	}
	if pc.PwLowercaseRequired {
		builder.WriteString("required: lower; ")
	}
	if pc.PwAlphabetRequired {
		builder.WriteString("required: [abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ]; ")
	}
	if pc.PwDigitRequired {
		builder.WriteString("required: digit; ")
	}
	if pc.PwSymbolRequired {
		builder.WriteString("required: special;")
	}

	return strings.TrimSpace(builder.String())
}

// PasswordPolicy outputs a list of PasswordPolicy to reflect the password policy.
func (pc *Checker) PasswordPolicy() (out []Policy) {
	if pc.PwMinLength > 0 {
		out = append(out, pc.policyPasswordLength())
	}
	if pc.PwUppercaseRequired {
		out = append(out, Policy{Name: PasswordUppercaseRequired})
	}
	if pc.PwLowercaseRequired {
		out = append(out, Policy{Name: PasswordLowercaseRequired})
	}
	if pc.PwDigitRequired {
		out = append(out, Policy{Name: PasswordDigitRequired})
	}
	if pc.PwAlphabetRequired {
		out = append(out, Policy{Name: PasswordAlphabetRequired})
	}
	if pc.PwSymbolRequired {
		out = append(out, Policy{Name: PasswordSymbolRequired})
	}
	if len(pc.PwExcludedKeywords) > 0 {
		out = append(out, Policy{Name: PasswordContainingExcludedKeywords})
	}
	if pc.PwMinGuessableLevel > 0 {
		out = append(out, pc.policyPasswordGuessableLevel())
	}
	if pc.shouldCheckPasswordHistory() {
		out = append(out, pc.policyPasswordHistory())
	}
	if out == nil {
		out = []Policy{}
	}
	return
}

func (pc *Checker) shouldCheckPasswordHistory() bool {
	return pc.PasswordHistoryEnabled
}

func IsSamePassword(hashedPassword []byte, password string) bool {
	return corepassword.Compare([]byte(password), hashedPassword) == nil
}
