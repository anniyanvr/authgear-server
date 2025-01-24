package password

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/authgear/authgear-server/pkg/lib/config"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/authgear/authgear-server/pkg/api/apierrors"
)

func TestPasswordCheckingFuncs(t *testing.T) {
	Convey("check password length", t, func() {
		So(checkPasswordLength("", 0), ShouldEqual, true)
		So(checkPasswordLength("", 1), ShouldEqual, false)
		So(checkPasswordLength("a", 1), ShouldEqual, true)
		So(checkPasswordLength("ab", 1), ShouldEqual, true)
	})
	Convey("check password uppercase", t, func() {
		So(checkPasswordUppercase("A"), ShouldEqual, true)
		So(checkPasswordUppercase("Z"), ShouldEqual, true)
		So(checkPasswordUppercase("a"), ShouldEqual, false)
	})
	Convey("check password lowercase", t, func() {
		So(checkPasswordLowercase("A"), ShouldEqual, false)
		So(checkPasswordLowercase("a"), ShouldEqual, true)
		So(checkPasswordLowercase("z"), ShouldEqual, true)
	})
	Convey("check password alphabet", t, func() {
		So(checkPasswordAlphabet("A"), ShouldEqual, true)
		So(checkPasswordAlphabet("a"), ShouldEqual, true)
		So(checkPasswordAlphabet("1"), ShouldEqual, false)
	})
	Convey("check password digit", t, func() {
		So(checkPasswordDigit("a"), ShouldEqual, false)
		So(checkPasswordDigit("0"), ShouldEqual, true)
		So(checkPasswordDigit("9"), ShouldEqual, true)
	})
	Convey("check password symbol", t, func() {
		So(checkPasswordSymbol("azAZ09"), ShouldEqual, false)
		So(checkPasswordSymbol("~"), ShouldEqual, true)
	})
	Convey("check password excluded keywords", t, func() {
		p := ".+[]{}^$QuoteRegexMetaCorrectly"
		kws := []string{".", "+", "[", "]", "{", "}", "^", "$"}
		So(checkPasswordExcludedKeywords(p, kws), ShouldEqual, false)

		p = "ADminIsEmbedded"
		kws = []string{"admin"}
		So(checkPasswordExcludedKeywords(p, kws), ShouldEqual, false)

		p = "user"
		kws = []string{"admin", "user"}
		So(checkPasswordExcludedKeywords(p, kws), ShouldEqual, false)

		So(checkPasswordExcludedKeywords(p, nil), ShouldEqual, true)

		p = "a_good_password"
		kws = []string{"bad"}
		So(checkPasswordExcludedKeywords(p, kws), ShouldEqual, true)
	})
	Convey("check password guessable level", t, func() {
		p := "nihongo-wo-manabimashou" // 日本語を学びましょう
		_, ok := checkPasswordGuessableLevel(p, 5)
		So(ok, ShouldEqual, true)
	})
}

func TestValidateNewPassword(t *testing.T) {
	// fixture
	authID := "chima"
	phData := map[string][]History{
		authID: []History{
			{
				ID:             "1",
				UserID:         authID,
				HashedPassword: []byte("$2a$10$EazYxG5cUdf99wGXDU1fguNxvCe7xQLEgr/Ay6VS9fkkVjHZtpJfm"), // "chima"
				CreatedAt:      time.Date(2017, 11, 3, 0, 0, 0, 0, time.UTC),
			},
			{
				ID:             "2",
				UserID:         authID,
				HashedPassword: []byte("$2a$10$8Z0zqmCZ3pZUlvLD8lN.B.ecN7MX8uVcZooPUFnCcB8tWR6diVc1a"), // "faseng"
				CreatedAt:      time.Date(2017, 11, 2, 0, 0, 0, 0, time.UTC),
			},
			{
				ID:             "3",
				UserID:         authID,
				HashedPassword: []byte("$2a$10$qzmi8TkYosj66xHvc9EfEulKjGoZswJSyNVEmmbLDxNGP/lMm6UXC"), // "coffee"
				CreatedAt:      time.Date(2017, 11, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}
	phStore := &mockPasswordHistoryStoreImpl{
		Data:    phData,
		TimeNow: func() time.Time { return time.Date(2017, 11, 4, 0, 0, 0, 0, time.UTC) },
	}

	test := func(pc *Checker, userID string, password string, expected string) {
		ctx := context.Background()
		err := pc.ValidateNewPassword(ctx, userID, password)
		if err != nil {
			e := apierrors.AsAPIError(err)
			b, _ := json.Marshal(e)
			So(string(b), ShouldEqualJSON, expected)
		} else {
			So(expected, ShouldBeEmpty)
		}
	}

	Convey("validate short password", t, func() {
		password := "1"
		pc := &Checker{
			PwMinLength: 2,
		}
		test(pc, "", password, `
		{
			"name": "Invalid",
			"reason": "PasswordPolicyViolated",
			"message": "password policy violated",
			"code": 400,
			"info": {
				"causes": [
					{
						"Name": "PasswordTooShort",
						"Info": {
							"min_length": 2,
							"pw_length": 1
						}
					}
				]
			}
		}
		`)
	})

	Convey("validate uppercase password", t, func() {
		password := "a"
		pc := &Checker{
			PwUppercaseRequired: true,
		}
		test(pc, "", password, `
		{
			"name": "Invalid",
			"reason": "PasswordPolicyViolated",
			"message": "password policy violated",
			"code": 400,
			"info": {
				"causes": [
					{
						"Name": "PasswordUppercaseRequired"
					}
				]
			}
		}
		`)
	})
	Convey("validate lowercase password", t, func() {
		password := "A"
		pc := &Checker{
			PwLowercaseRequired: true,
		}
		test(pc, "", password, `
		{
			"name": "Invalid",
			"reason": "PasswordPolicyViolated",
			"message": "password policy violated",
			"code": 400,
			"info": {
				"causes": [
					{
						"Name": "PasswordLowercaseRequired"
					}
				]
			}
		}
		`)
	})
	Convey("validate alphabet password", t, func() {
		password := "123"
		pc := &Checker{
			PwAlphabetRequired: true,
		}
		test(pc, "", password, `
		{
			"name": "Invalid",
			"reason": "PasswordPolicyViolated",
			"message": "password policy violated",
			"code": 400,
			"info": {
				"causes": [
					{
						"Name": "PasswordAlphabetRequired"
					}
				]
			}
		}
		`)
	})
	Convey("validate digit password", t, func() {
		password := "-"
		pc := &Checker{
			PwDigitRequired: true,
		}
		test(pc, "", password, `
		{
			"name": "Invalid",
			"reason": "PasswordPolicyViolated",
			"message": "password policy violated",
			"code": 400,
			"info": {
				"causes": [
					{
						"Name": "PasswordDigitRequired"
					}
				]
			}
		}
		`)
	})
	Convey("validate symbol password", t, func() {
		password := "azAZ09"
		pc := &Checker{
			PwSymbolRequired: true,
		}
		test(pc, "", password, `
		{
			"name": "Invalid",
			"reason": "PasswordPolicyViolated",
			"message": "password policy violated",
			"code": 400,
			"info": {
				"causes": [
					{
						"Name": "PasswordSymbolRequired"
					}
				]
			}
		}
		`)
	})
	Convey("validate excluded keywords password", t, func() {
		password := "useradmin1"
		pc := &Checker{
			PwExcludedKeywords: []string{"user"},
		}
		test(pc, "", password, `
		{
			"name": "Invalid",
			"reason": "PasswordPolicyViolated",
			"message": "password policy violated",
			"code": 400,
			"info": {
				"causes": [
					{
						"Name": "PasswordContainingExcludedKeywords"
					}
				]
			}
		}
		`)
	})
	Convey("validate guessable password", t, func() {
		password := "abcde123456"
		pc := &Checker{
			PwMinGuessableLevel: 5,
		}
		test(pc, "", password, `
		{
			"name": "Invalid",
			"reason": "PasswordPolicyViolated",
			"message": "password policy violated",
			"code": 400,
			"info": {
				"causes": [
					{
						"Name": "PasswordBelowGuessableLevel",
						"Info": {
							"min_level": 5,
							"pw_level": 2
						}
					}
				]
			}
		}
		`)
	})
	Convey("validate strong password", t, func() {
		// nolint:gosec
		password := "N!hon-no-tsuk!-wa-seka!-1ban-k!re!desu" // 日本の月は世界一番きれいです
		pc := &Checker{
			PwMinLength:         8,
			PwUppercaseRequired: true,
			PwLowercaseRequired: true,
			PwDigitRequired:     true,
			PwSymbolRequired:    true,
			PwMinGuessableLevel: 5,
			PwExcludedKeywords:  []string{"user", "admin"},
		}
		ctx := context.Background()
		So(
			pc.ValidateNewPassword(ctx, "", password),
			ShouldBeNil,
		)
	})

	Convey("validate password history", t, func(c C) {
		historySize := 12
		historyDays := config.DurationDays(365)

		pc := &Checker{
			PwHistorySize:          historySize,
			PwHistoryDays:          historyDays,
			PasswordHistoryEnabled: true,
			PasswordHistoryStore:   phStore,
		}

		test(pc, authID, "chima", `
		{
			"name": "Invalid",
			"reason": "PasswordPolicyViolated",
			"message": "password policy violated",
			"code": 400,
			"info": {
				"causes": [
					{
						"Name": "PasswordReused",
						"Info": {
							"history_size": 12,
							"history_days": 365
						}
					}
				]
			}
		}
		`)

		test(pc, authID, "coffee", `
		{
			"name": "Invalid",
			"reason": "PasswordPolicyViolated",
			"message": "password policy violated",
			"code": 400,
			"info": {
				"causes": [
					{
						"Name": "PasswordReused",
						"Info": {
							"history_size": 12,
							"history_days": 365
						}
					}
				]
			}
		}
		`)

		test(pc, authID, "milktea", "")
	})

	Convey("validate password history by size", t, func(c C) {
		historySize := 2
		historyDays := config.DurationDays(0)

		pc := &Checker{
			PwHistorySize:          historySize,
			PwHistoryDays:          historyDays,
			PasswordHistoryEnabled: true,
			PasswordHistoryStore:   phStore,
		}

		test(pc, authID, "chima", `
		{
			"name": "Invalid",
			"reason": "PasswordPolicyViolated",
			"message": "password policy violated",
			"code": 400,
			"info": {
				"causes": [
					{
						"Name": "PasswordReused",
						"Info": {
							"history_size": 2,
							"history_days": 0
						}
					}
				]
			}
		}
		`)

		test(pc, authID, "coffee", "")
	})

	Convey("validate password history by days", t, func(c C) {
		historySize := 0
		historyDays := config.DurationDays(2)

		pc := &Checker{
			PwHistorySize:          historySize,
			PwHistoryDays:          historyDays,
			PasswordHistoryEnabled: true,
			PasswordHistoryStore:   phStore,
		}

		test(pc, authID, "chima", `
		{
			"name": "Invalid",
			"reason": "PasswordPolicyViolated",
			"message": "password policy violated",
			"code": 400,
			"info": {
				"causes": [
					{
						"Name": "PasswordReused",
						"Info": {
							"history_size": 0,
							"history_days": 2
						}
					}
				]
			}
		}
		`)

		test(pc, authID, "coffee", "")
	})
}

func TestPasswordPolicy(t *testing.T) {
	Convey("PasswordPolicy", t, func() {
		Convey("empty", func() {
			pc := &Checker{}
			So(pc.PasswordPolicy(), ShouldBeEmpty)
			So(pc.PasswordPolicy(), ShouldNotBeNil)
		})
		Convey("length", func() {
			pc := &Checker{
				PwMinLength: 8,
			}
			So(pc.PasswordPolicy(), ShouldResemble, []Policy{
				Policy{
					Name: PasswordTooShort,
					Info: map[string]interface{}{
						"min_length": 8,
					},
				},
			})
		})
		Convey("guessable level", func() {
			pc := &Checker{
				PwMinGuessableLevel: 3,
			}
			So(pc.PasswordPolicy(), ShouldResemble, []Policy{
				Policy{
					Name: PasswordBelowGuessableLevel,
					Info: map[string]interface{}{
						"min_level": 3,
					},
				},
			})
		})
		Convey("history", func() {
			pc := &Checker{
				PasswordHistoryEnabled: true,
				PwHistorySize:          10,
				PwHistoryDays:          90,
			}
			So(pc.PasswordPolicy(), ShouldResemble, []Policy{
				Policy{
					Name: PasswordReused,
					Info: map[string]interface{}{
						"history_size": 10,
						"history_days": 90,
					},
				},
			})
		})
		Convey("only output effective policies", func() {
			pc := &Checker{
				PwUppercaseRequired: true,
				PwDigitRequired:     true,
			}
			So(pc.PasswordPolicy(), ShouldResemble, []Policy{
				Policy{
					Name: PasswordUppercaseRequired,
				},
				Policy{
					Name: PasswordDigitRequired,
				},
			})
		})
	})
}

func TestPasswordRules(t *testing.T) {
	Convey("PasswordRules", t, func() {
		Convey("validate passwordrules", func() {
			pc := &Checker{
				PwMinLength:         20,
				PwUppercaseRequired: true,
				PwDigitRequired:     true,
				PwSymbolRequired:    true,
			}
			So(pc.PasswordRules(), ShouldEqual, "minlength: 20; required: upper; required: digit; required: special;")
		})
	})
}
