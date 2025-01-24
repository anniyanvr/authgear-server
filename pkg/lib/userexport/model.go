package userexport

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/iawaknahc/jsonschema/pkg/jsonpointer"

	"github.com/authgear/authgear-server/pkg/api/apierrors"
	"github.com/authgear/authgear-server/pkg/api/model"
	"github.com/authgear/authgear-server/pkg/lib/infra/redisqueue"
	"github.com/authgear/authgear-server/pkg/util/duration"
)

// PresignGetExpiresForUserExport is how long the presign GET request remains valid for user export.
const PresignGetExpiresForUserExport time.Duration = 1 * duration.PerMinute

const BatchSize = 1000

var defaultCSVExportFields = []*FieldPointer{
	&FieldPointer{Pointer: "/sub"},
	&FieldPointer{Pointer: "/preferred_username"},
	&FieldPointer{Pointer: "/email"},
	&FieldPointer{Pointer: "/phone_number"},
	&FieldPointer{Pointer: "/email_verified"},
	&FieldPointer{Pointer: "/phone_number_verified"},
	&FieldPointer{Pointer: "/name"},
	&FieldPointer{Pointer: "/given_name"},
	&FieldPointer{Pointer: "/middle_name"},
	&FieldPointer{Pointer: "/nickname"},
	&FieldPointer{Pointer: "/profile"},
	&FieldPointer{Pointer: "/picture"},
	&FieldPointer{Pointer: "/website"},
	&FieldPointer{Pointer: "/gender"},
	&FieldPointer{Pointer: "/birthdate"},
	&FieldPointer{Pointer: "/zoneinfo"},
	&FieldPointer{Pointer: "/locale"},
	&FieldPointer{Pointer: "/address/formatted"},
	&FieldPointer{Pointer: "/address/street_address"},
	&FieldPointer{Pointer: "/address/locality"},
	&FieldPointer{Pointer: "/address/region"},
	&FieldPointer{Pointer: "/address/postal_code"},
	&FieldPointer{Pointer: "/address/country"},
	&FieldPointer{Pointer: "/roles"},
	&FieldPointer{Pointer: "/groups"},
	&FieldPointer{Pointer: "/disabled"},
	&FieldPointer{Pointer: "/delete_at"},
	&FieldPointer{Pointer: "/identities"},
	&FieldPointer{Pointer: "/mfa/emails"},
	&FieldPointer{Pointer: "/mfa/phone_numbers"},
	&FieldPointer{Pointer: "/mfa/totps"},
	&FieldPointer{Pointer: "/biometric_count"},
	&FieldPointer{Pointer: "/passkey_count"},
}

type FieldPointer struct {
	Pointer   string `json:"pointer,omitempty"`
	FieldName string `json:"field_name,omitempty"`
}

type CSVField struct {
	Fields []*FieldPointer `json:"fields,omitempty"`
}

type Request struct {
	Format string    `json:"format,omitempty"`
	CSV    *CSVField `json:"csv,omitempty"`
}

type Response struct {
	ID          string                `json:"id,omitempty"`
	CreatedAt   *time.Time            `json:"created_at,omitempty"`
	CompletedAt *time.Time            `json:"completed_at,omitempty"`
	FailedAt    *time.Time            `json:"failed_at,omitempty"`
	Status      redisqueue.TaskStatus `json:"status,omitempty"`
	Request     *Request              `json:"request,omitempty"`
	DownloadUrl string                `json:"download_url,omitempty"`
	Error       *apierrors.APIError   `json:"error,omitempty"`
}

type Result struct {
	Filename string              `json:"file_name,omitempty"`
	Error    *apierrors.APIError `json:"error,omitempty"`
}

type Address struct {
	Formatted     string `json:"formatted,omitempty"`
	StreetAddress string `json:"street_address,omitempty"`
	Locality      string `json:"locality,omitempty"`
	Region        string `json:"region,omitempty"`
	PostalCode    string `json:"postal_code,omitempty"`
	Country       string `json:"country,omitempty"`
}

type Identity struct {
	Type    model.IdentityType     `json:"type"`
	LoginID *IdentityLoginID       `json:"login_id,omitempty"`
	OAuth   *IdentityOAuth         `json:"oauth,omitempty"`
	LDAP    *IdentityLDAP          `json:"ldap,omitempty"`
	Claims  map[string]interface{} `json:"claims,omitempty"`
}

type IdentityLoginID struct {
	Key           string `json:"key,omitempty"`
	Type          string `json:"type,omitempty"`
	Value         string `json:"value,omitempty"`
	OriginalValue string `json:"original_value,omitempty"`
}

type IdentityLDAP struct {
	ServerName           string                 `json:"server_name,omitempty"`
	LastLoginUsername    string                 `json:"last_login_username,omitempty"`
	UserIDAttributeName  string                 `json:"user_id_attribute_name,omitempty"`
	UserIDAttributeValue string                 `json:"user_id_attribute_value,omitempty"`
	Attributes           map[string]interface{} `json:"attributes,omitempty"`
}

type IdentityOAuth struct {
	ProviderAlias     string                 `json:"provider_alias,omitempty"`
	ProviderType      string                 `json:"provider_type,omitempty"`
	ProviderSubjectID string                 `json:"provider_subject_id,omitempty"`
	UserProfile       map[string]interface{} `json:"user_profile,omitempty"`
}

type MFATOTP struct {
	Secret string `json:"secret,omitempty"`
	URI    string `json:"uri,omitempty"`
}

type MFA struct {
	Emails       []string   `json:"emails,omitempty"`
	PhoneNumbers []string   `json:"phone_numbers,omitempty"`
	TOTPs        []*MFATOTP `json:"totps,omitempty"`
}

type Record struct {
	Sub string `json:"sub,omitempty"`

	PreferredUsername string `json:"preferred_username,omitempty"`
	Email             string `json:"email,omitempty"`
	PhoneNumber       string `json:"phone_number,omitempty"`

	EmailVerified       bool `json:"email_verified"`
	PhoneNumberVerified bool `json:"phone_number_verified"`

	Name       string   `json:"name,omitempty"`
	GivenName  string   `json:"given_name,omitempty"`
	FamilyName string   `json:"family_name,omitempty"`
	MiddleName string   `json:"middle_name,omitempty"`
	Nickname   string   `json:"nickname,omitempty"`
	Profile    string   `json:"profile,omitempty"`
	Picture    string   `json:"picture,omitempty"`
	Website    string   `json:"website,omitempty"`
	Gender     string   `json:"gender,omitempty"`
	Birthdate  string   `json:"birthdate,omitempty"`
	Zoneinfo   string   `json:"zoneinfo,omitempty"`
	Locale     string   `json:"locale,omitempty"`
	Address    *Address `json:"address,omitempty"`

	CustomAttributes map[string]interface{} `json:"custom_attributes,omitempty"`

	Roles  []string `json:"roles,omitempty"`
	Groups []string `json:"groups,omitempty"`

	Disabled bool       `json:"disabled"`
	DeleteAt *time.Time `json:"delete_at,omitempty"`

	Identities []*Identity `json:"identities,omitempty"`

	Mfa *MFA `json:"mfa,omitempty"`

	BiometricCount int `json:"biometric_count"`
	PasskeyCount   int `json:"passkey_count"`
}

func NewResponseFromTask(task *redisqueue.Task) (*Response, error) {
	response := &Response{
		ID:        task.ID,
		CreatedAt: task.CreatedAt,
		Status:    task.Status,
	}

	if task.Input != nil {
		var request Request
		err := json.Unmarshal(task.Input, &request)
		if err != nil {
			return nil, err
		}
		response.Request = &request
	}

	if task.Output != nil {
		var result Result
		err := json.Unmarshal(task.Output, &result)
		if err != nil {
			return nil, err
		}

		if result.Error != nil {
			response.FailedAt = task.CompletedAt
			response.Error = result.Error
		} else {
			response.CompletedAt = task.CompletedAt
			response.DownloadUrl = result.Filename
		}
	}

	return response, nil
}

func ExtractCSVHeaderField(fieldPointer []*FieldPointer) (headerFields []string, err error) {
	isDuplicated := false
	fields := make([]string, 0)
	fieldsMap := map[string]bool{}
	for _, pointer := range fieldPointer {
		var fieldName string
		if pointer.FieldName == "" {
			ptr, err := jsonpointer.Parse(pointer.Pointer)
			if err != nil {
				return nil, err
			}
			fieldName = strings.Join(ptr, ".")
		} else {
			fieldName = pointer.FieldName
		}

		if fieldsMap[fieldName] {
			isDuplicated = true
		}

		fieldsMap[fieldName] = true
		fields = append(fields, fieldName)
	}

	if isDuplicated {
		info := apierrors.Details{
			"field_names": fields,
		}
		return nil, ErrUserExportDuplicateField.NewWithInfo("field names are not unique", info)
	}

	return fields, nil
}

func TraverseRecordValue(jsonMap interface{}, pointer string) (fieldValue string, err error) {
	ptr, err := jsonpointer.Parse(pointer)
	if err != nil {
		return "", err
	}
	value, err := ptr.Traverse(jsonMap)
	if err != nil {
		return "", err
	}

	switch v := value.(type) {
	case bool:
		if v {
			fieldValue = "true"
		} else {
			fieldValue = "false"
		}
	case []interface{}:
		valueJson, _ := json.Marshal(v)
		fieldValue = string(valueJson)
	case map[string]interface{}:
		valueJson, _ := json.Marshal(v)
		fieldValue = string(valueJson)
	case float64:
		fieldValue = strconv.FormatFloat(v, 'f', -1, 64)
	case nil:
		fieldValue = ""
	case string:
		fieldValue = v
	default:
		panic(fmt.Sprintf("Unsupported JSON value in user export: %T, %v\n", v, v))
	}

	return fieldValue, nil
}
