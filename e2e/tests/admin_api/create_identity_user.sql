{{ $user_id := (uuidv4) }}
{{ $anon_identity_id := (uuidv4) }}

INSERT INTO _auth_user
  ("id", "app_id", "created_at", "updated_at", "last_login_at", "login_at", "is_disabled", "disable_reason", "standard_attributes", "custom_attributes", "is_deactivated", "delete_at", "is_anonymized", "anonymize_at", "anonymized_at", "last_indexed_at", "require_reindex_after")
VALUES
  ('{{ $user_id }}', '{{ .AppID }}', '2024-08-27 07:51:42.040683', '2024-08-27 07:51:42.056654', NULL, NULL, 'f', NULL, '{}', '{}', 'f', NULL, 'f', NULL, NULL, '2024-08-27 07:51:42.079532', '2024-08-27 07:51:42.059056');

INSERT INTO _auth_identity (
  "id",
  "app_id",
  "type",
  "user_id",
  "created_at",
  "updated_at"
) VALUES (
  '{{ $anon_identity_id }}',
  '{{ .AppID }}',
  'anonymous',
  '{{ $user_id }}',
  '2024-08-27 07:51:42.040683',
  '2024-08-27 07:51:42.040683'
);

INSERT INTO _auth_identity_anonymous (
  "id",
  "app_id",
  "key_id",
  "key"
) VALUES (
  '{{ $anon_identity_id }}',
  '{{ .AppID }}',
  'key_id',
  '{}'
);
