- [User Export](#user-export)
  * [About User Export](#about-user-export)
  * [Create an export](#create-an-export)
    + [The usage limit](#the-usage-limit)
    + [The request body of Create an export](#the-request-body-of-create-an-export)
      - [Default CSV fields](#default-csv-fields)
      - [The field name](#the-field-name)
    + [The response body of Create an export](#the-response-body-of-create-an-export)
    + [The error response of Create an export](#the-error-response-of-create-an-export)
  * [Get the status of an export](#get-the-status-of-an-export)
    + [The response body of Get the status of an export](#the-response-body-of-get-the-status-of-an-export)
    + [The error response of Get the status of an export](#the-error-response-of-get-the-status-of-an-export)
  * [The response body](#the-response-body)
  * [The export file](#the-export-file)
    + [The name of the export file](#the-name-of-the-export-file)
    + [The metadata of the export file](#the-metadata-of-the-export-file)
    + [The storage of the export file](#the-storage-of-the-export-file)
    + [The content of the export file](#the-content-of-the-export-file)
    + [The record format](#the-record-format)
  * [Caveats](#caveats)

# User Export

User Export allows the developer to bulk export all users from Authgear to a file.

## About User Export

- It is not a synchronous operation. The export is created and runs in the background. The developer can query the status of it.

## Create an export

The endpoint is `POST /_api/admin/users/export`.

- The endpoint requires Admin API JWT token to access.
- This endpoint is added to Admin API server.
- It is not part of the GraphQL API.
- At most 1 running export at any given moment for a project.
- There is a usage limit of user export. See [the usage limit](#the-usage-limit)
- A pending export lasts for 24h before it expires.

### The usage limit

The usage limit is specified in `authgear.features.yaml`.
The following example shows the default usage limit.

```yaml
admin_api:
  user_export_usage:
    enabled: true
    period: day
    quota: 24
```

### The request body of Create an export

```
{
  "format": "ndjson",
  "csv": {
    "fields": [
      {
        "pointer": "/sub",
        "field_name": "user_id"
      }
    ]
  }
}
```

- `format`: Required. It must be `ndjson` or `csv`.
  - `ndjson`: The output is a ndjson file. See https://github.com/ndjson/ndjson-spec
  - `csv`: The output is a CSV file. See https://datatracker.ietf.org/doc/html/rfc4180
- `csv.fields`: Optional. See [Default CSV fields] for the list of default fields. If this is specified, then it must be an non-empty list.
  - `csv.fields.pointer`: Required. Select which field in the record to output. It must be a JSON pointer of at least one reference token. Each reference token must be non-empty. See https://datatracker.ietf.org/doc/html/rfc6901 and [The record format](#the-record-format)
  - `csv.fields.field_name`: See [The field name](#the-field-name). See [The content of the export file](#the-content-of-the-export-file) for how values are written.

#### Default CSV fields

If `csv.fields` is unspecified, the default is derived with the following rules:

- Let list be the following list.
```
[
  {"pointer": "/sub"},
  {"pointer": "/preferred_username"},
  {"pointer": "/email"},
  {"pointer": "/phone_number"},
  {"pointer": "/email_verified"},
  {"pointer": "/phone_number_verified"},
  {"pointer": "/name"},
  {"pointer": "/given_name"},
  {"pointer": "/middle_name"},
  {"pointer": "/nickname"},
  {"pointer": "/profile"},
  {"pointer": "/picture"},
  {"pointer": "/website"},
  {"pointer": "/gender"},
  {"pointer": "/birthdate"},
  {"pointer": "/zoneinfo"},
  {"pointer": "/locale"},
  {"pointer": "/address/formatted"},
  {"pointer": "/address/street_address"},
  {"pointer": "/address/locality"},
  {"pointer": "/address/region"},
  {"pointer": "/address/postal_code"},
  {"pointer": "/address/country"},
  {"pointer": "/roles"},
  {"pointer": "/groups"},
  {"pointer": "/disabled"},
  {"pointer": "/identities"},
  {"pointer": "/mfa/emails"},
  {"pointer": "/mfa/phone_numbers"},
  {"pointer": "/mfa/totps"},
  {"pointer": "/biometric_count"},
  {"pointer": "/passkey_count"}
]
```
- For each custom attribute of the project, add the following to the list.
```
{"pointer": "/custom_attributes/CUSTOM_ATTRIBUTE_NAME"}
```

For example, if the project has two custom attributes `member_id` and `loyalty_system_user_id`, the default is

```
[
  {"pointer": "/sub"},
  {"pointer": "/preferred_username"},
  {"pointer": "/email"},
  {"pointer": "/phone_number"},
  {"pointer": "/email_verified"},
  {"pointer": "/phone_number_verified"},
  {"pointer": "/name"},
  {"pointer": "/given_name"},
  {"pointer": "/middle_name"},
  {"pointer": "/nickname"},
  {"pointer": "/profile"},
  {"pointer": "/picture"},
  {"pointer": "/website"},
  {"pointer": "/gender"},
  {"pointer": "/birthdate"},
  {"pointer": "/zoneinfo"},
  {"pointer": "/locale"},
  {"pointer": "/address/formatted"},
  {"pointer": "/address/street_address"},
  {"pointer": "/address/locality"},
  {"pointer": "/address/region"},
  {"pointer": "/address/postal_code"},
  {"pointer": "/address/country"},
  {"pointer": "/roles"},
  {"pointer": "/groups"},
  {"pointer": "/disabled"},
  {"pointer": "/identities"},
  {"pointer": "/mfa/emails"},
  {"pointer": "/mfa/phone_numbers"},
  {"pointer": "/mfa/totps"},
  {"pointer": "/biometric_count"},
  {"pointer": "/passkey_count"},
  {"pointer": "/custom_attributes/member_id"},
  {"pointer": "/custom_attributes/loyalty_system_user_id"}
]
```

#### The field name

- `field_name` is optional.
- If `field_name` is given, then it is used as is.
- If `field_name` is not given, then it is derived from `pointer` with the following rules.
  - Let `parts` be the list of the reference tokens in `pointer`.
  - Join `parts` with the character `.`.

For example, given `pointer` is `/address/formatted`,

- Then `parts` is `["address", "formatted"]`.
- The join result is `address.formatted`.
- The field name is the join result.

For example, given `pointer` is `/roles/0`,

- Then `parts` is `["roles", "0"]`.
- The join result is `roles.0`.
- The field name is the join result.

Regardless of whether the field names are given or derived, they must be unique.
If the field names are not unique, it is an error when the export is created.
An error is immediately returned in this case, the export is not created.
See [The error response of Create an export](#the-error-response-of-create-an-export)

### The response body of Create an export

See [The response body](#the-response-body).

### The error response of Create an export

> we have a global middleware of handling Admin API authentication,
> that middleware returns 403 without a JSON body when authentication failed.

|Description|Name|Reason|Info|
|---|---|---|---|
|When user export is disabled|`InternalError`|`UserExportDisabled`|-|
|When the rate limit exceeded|`TooManyRequest`|`RateLimited`|{"bucket_name": "UserExport"}|
|When there is a running export|`TooManyRequest`|`MaximumConcurrentJobLimitExceeded`|-|
|When the input fails the validation|`Invalid`|`ValidationFailed`|The info should contain the JSON schema validation output|
|When the field names are not unique|`Invalid`|`UserExportNonUniqueFieldNames`|Output the full list of "field_names". Like `{"field_names": ["sub", "a", "b", "a"]}`|

## Get the status of an export

The endpoint is `GET /_api/admin/users/export/{ID}`

- The endpoint requires Admin API JWT token to access.
- This endpoint is added to Admin API server.
- It is not part of the GraphQL API.
- The result of a completed export lasts for 24h before it expires.
- When the export is completed, the response body includes a freshly signed URL to the export file. The signed URL is valid for 60s.

### The response body of Get the status of an export

See [The response body](#the-response-body).

### The error response of Get the status of an export

> we have a global middleware of handling Admin API authentication,
> that middleware returns 403 without a JSON body when authentication failed.

|Description|Name|Reason|Info|
|---|---|---|---|
|When user export is disabled|`InternalError`|`UserExportDisabled`|-|
|When the given ID does not refer to an export|`NotFound`|`TaskNotFound`|-|

## The response body

The response body of a just created export looks like

```
{
  "id": "some_opaque_string",
  "status": "pending",
  "created_at": "2024-01-01T00:00:00.000Z",
  "request": {
    "format": "ndjson"
  }
}
```

The response body of a completed export looks like

```
{
  "id": "some_opaque_string",
  "status": "completed",
  "created_at": "2024-01-01T00:00:00.000Z",
  "completed_at": "2024-01-01T00:01:00.000Z",
  "request": {
    "format": "ndjson"
  },
  "download_url": "https://some-signed-url?with-a=signature"
}
```

The response body of a failed export looks like

```
{
  "id": "some_opaque_string",
  "status": "completed",
  "created_at": "2024-01-01T00:00:00.000Z",
  "failed_at": "2024-01-01T00:01:00.000Z",
  "request": {
    "format": "ndjson"
  },
  "error": {
    "message": "blahblahblah",
    "reason": "SomeReason"
  }
}
```

- `error`: The API error object we have been using in all other API.

## The export file

### The name of the export file

The name of the export file is `{{ .AppID }}-{{ .TaskID }}-{{ .TaskCompletedAtInISO9601BasicFormat}}.{ndjson|csv}`.

For example, given

- The project is `myapp`.
- The id of the task is `userexport_deadbeef`.
- The completion time of the task is `2024-09-09T10:46:51.275Z`.
- The requested `format` is `ndjson`.

The name of the export file is `myapp-userexport_deadbeef-20240909104651Z.ndjson`.

### The metadata of the export file

- `Content-Disposition`: `attachment; filename=FILENAME`
- `Content-Type`: `application/x-ndjson` for `ndjson`, `text/csv` for `csv`.

Since the export file is accessed with a signed URL of a short validity, setting `Cache-Control` is not really helpful.

### The storage of the export file

- GCS: https://cloud.google.com/storage/docs/lifecycle
- S3: https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lifecycle-mgmt.html
- Azure: https://learn.microsoft.com/en-us/azure/storage/blobs/lifecycle-management-overview?tabs=azure-portal

The above cloud storage can be configured to delete objects of a certain age.
So there is no need to housekeep manually.

The following environment variables are added to configure the cloud storage to store export files.
Every project in the deployment shares the same cloud storage.

```
USEREXPORT_OBJECT_STORE_TYPE=AWS_S3
USEREXPORT_OBJECT_STORE_AWS_S3_BUCKET_NAME=
USEREXPORT_OBJECT_STORE_AWS_S3_REGION=
USEREXPORT_OBJECT_STORE_AWS_S3_ACCESS_KEY_ID=
USEREXPORT_OBJECT_STORE_AWS_S3_SECRET_ACCESS_KEY=

USEREXPORT_OBJECT_STORE_TYPE=GCP_GCS
USEREXPORT_OBJECT_STORE_GCP_GCS_BUCKET_NAME=
USEREXPORT_OBJECT_STORE_GCP_GCS_SERVICE_ACCOUNT=
USEREXPORT_OBJECT_STORE_GCP_GCS_CREDENTIALS_JSON_PATH=

USEREXPORT_OBJECT_STORE_TYPE=AZURE_BLOB_STORAGE
USEREXPORT_OBJECT_STORE_AZURE_BLOB_STORAGE_STORAGE_ACCOUNT=
USEREXPORT_OBJECT_STORE_AZURE_BLOB_STORAGE_CONTAINER=
USEREXPORT_OBJECT_STORE_AZURE_BLOB_STORAGE_SERVICE_URL=
USEREXPORT_OBJECT_STORE_AZURE_BLOB_STORAGE_ACCESS_KEY=
```

If `USEREXPORT_OBJECT_STORE_TYPE` is not set, then the user export feature is disabled.
See [The error response of Create an export](#the-error-response-of-create-an-export) and
[The error response of Get the status of an export](#the-error-responseof-get-the-status-of-an-export).

### The content of the export file

- If `format` is `ndjson`, then the file contains a record (See [The record format](#the-record-format)) per line.
  - Each line is terminated by a `\n` (The newline character).
  - The number of lines correspond to the number of exported records.
  - Exporting in a project without any user will in a file of zero length.

> Implementation note: Please make sure each line is ended by a `\n`!

- If `format` is `csv`, then the file starts with a header, followed by records.
  - The header correspond to the `csv.fields`.
  - The order of the field in the header correspond to the order in `csv.fields`.
  - The name of the field is taken from `csv.fields.field_name` if it is specified, or derived from `csv.fields.pointer` if `field_name` is absent.
  - The CSV follows the format documented in RFC4180. Internally, it is handled with https://pkg.go.dev/encoding/csv.
  - If `pointer` resolves to
    - JSON string, then the string is written directly.
    - JSON number, then the number is written directly.
    - JSON boolean, then `true` or `false` is written directly.
    - JSON null, then an empty string is written.
    - non-existing value, then an empty string is written.
    - JSON array, then the compact JSON serialization is written.
    - JSON object, then the compact JSON serialization is written.

For example, given the request

```
{
  "format": "csv",
  "csv": {
    "fields": [
      {
        "pointer": "/sub"
      },
      {
        "pointer": "/roles"
      },
      {
        "pointer": "/address"
      },
      {
        "pointer": "/address/formatted",
        "field_name": "address_formatted"
      }
    ]
  }
}
```

, and a record

```
{
  "sub": "opaque_user_id",

  "address": {
    "formatted": "1 Unnamed Road, Central, Hong Kong Island, HK",
    "street_address": "1 Unnamed Road",
    "locality": "Central",
    "region": "Hong Kong",
    "postal_code": "N/A",
    "country": "HK"
  },

  "roles": ["role_a", "role_b"]
}
```

The content of the file is

```
sub,roles,address,address_formatted
opaque_user_id,"[""role_a"",""role_b""]","{""formatted"":""1 Unnamed Road, Central, Hong Kong Island, HK"",""street_address"":""1 Unnamed Road"",""locality"":""Central"",""region"":""Hong Kong"",""postal_code"":""N/A"",""country"":""HK""}","1 Unnamed Road, Central, Hong Kong Island, HK"
```

### The record format

Here is an example of the record

```
{
  "sub": "opaque_user_id",

  "preferred_username": "louischan",
  "email": "louischan@oursky.com",
  "phone_number": "+85298765432",

  "email_verified": true,
  "phone_number_verified": true,

  "name": "Louis Chan",
  "given_name": "Louis",
  "family_name": "Chan",
  "middle_name": "",
  "nickname": "Lou",
  "profile": "https://example.com",
  "picture": "https://example.com",
  "website": "https://example.com",
  "gender": "male",
  "birthdate": "1990-01-01",
  "zoneinfo": "Asia/Hong_Kong",
  "locale": "zh-Hant-HK",
  "address": {
    "formatted": "1 Unnamed Road, Central, Hong Kong Island, HK",
    "street_address": "1 Unnamed Road",
    "locality": "Central",
    "region": "Hong Kong",
    "postal_code": "N/A",
    "country": "HK"
  },

  "custom_attributes": {
    "member_id": "123456789"
  },

  "roles": ["role_a", "role_b"],
  "groups": ["group_a"],

  "disabled": false,
  "delete_at": "1985-04-12T23:20:50.52Z",

  "identities": [
    {
      "type": "login_id",
      "login_id": {
        "type": "username",
        "key": "username",
        "value": "louischan",
        "original_value": "LOUISCHAN"
      },
      "claims": {
        "preferred_username": "louischan"
      }
    },
    {
      "type": "login_id",
      "login_id": {
        "type": "email",
        "key": "email",
        "value": "louischan@oursky.com",
        "original_value": "LOUISCHAN@oursky.com"
      },
      "claims": {
        "email": "louischan@oursky.com"
      }
    },
    {
      "type": "login_id",
      "login_id": {
        "type": "phone",
        "key": "phone",
        "value": "+85298765432",
        "original_value": "+85298765432"
      },
      "claims": {
        "phone_number": "+85298765432"
      }
    },
    {
      "type": "oauth",
      "oauth": {
        "provider_alias": "google",
        "provider_type": "google",
        "provider_subject_id": "blahblahblah"
        "user_profile": {
          "email": "louischan@oursky.com"
        }
      },
      "claims": {
        "email": "louischan@oursky.com"
      }
    },
    {
      "type": "ldap",
      "ldap": {
        "server_name": "myldap",
        "last_login_username": "louischan",
        "user_id_attribute_name": "uid",
        "user_id_attribute_value": "blahblahblah",
        "attributes": {
          "dn": "the DN"
        }
      },
      "claims": {
        "preferred_username": "louischan"
      }
    }
  ],

  "mfa": {
    "emails": ["louischan@oursky.com"],
    "phone_numbers": ["+85298765432"]
    "totps" [
      {
        "secret": "the-secret",
        "uri": "otpauth://totp...."
      }
    ]
  },

  "biometric_count": 0,
  "passkey_count": 0,
}
```

- `sub`: The Authgear user ID.
- `preferred_username`: The primary username of the user.
- `email`: The primary email of the user.
  - `email_verified`: Whether the email is verified.
- `phone_number`: The primary phone number of the user.
  - `phone_number_verified`: Whether the phone number is verified.
- `name`, `given_name`, `family_name`, `middle_name`, `nickname`: The OIDC standard attributes about names. See https://openid.net/specs/openid-connect-core-1_0.html#StandardClaims
- `profile`, `picture`, `website`, `gender`, `birthdate`, `zoneinfo`, `locale`, `address`: Other OIDC standard attributes. See https://openid.net/specs/openid-connect-core-1_0.html#StandardClaims
- `custom_attributes`: All custom_attributes of the user.
- `roles`: The role keys of all roles directly assigned to the user. This DOES NOT include roles implied by the groups.
- `groups`: The group keys of all groups assigned to the user.
- `disabled`: Whether the user is disabled.
- `delete_at`: The timestamp in RFC3339 format indicating when the scheduled deletion of the user will happen. If it is absent, the user is not scheduled for deletion.
- `identities`: The array of identities that may contain the following types. It is an empty array if the user no identities of the following types.
  - `login_id`
  - `oauth`
  - `ldap`
- `mfa.emails`: The list of MFA emails the user has. It is an empty array if the user has none.
- `mfa.phone_numbers`: The list of MFA phone numbers the user has. It is an empty array if the user has none.
- `mfa.totps`: The list of MFA TOTP authenticator the user has. It is an empty array if the user has none.
- `biometric_count`: The number of biometric login the user has.
- `passkey_count`: The number of passkey the user has.

> Future work: Support exporting the password hash.

## Caveats

Since we do not export the password hash,
The exported JSON record cannot be imported into another Authgear project directly.
