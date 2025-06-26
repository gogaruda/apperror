package apperror

// General codes
const (
	CodeInternalError   = "INTERNAL_ERROR"
	CodeInvalidInput    = "INVALID_INPUT"
	CodeUnauthorized    = "UNAUTHORIZED"
	CodeForbidden       = "FORBIDDEN"
	CodeNotImplemented  = "NOT_IMPLEMENTED"
	CodeTimeout         = "TIMEOUT"
	CodeDependencyError = "DEPENDENCY_ERROR"
	CodeBadRequest      = "BAD_REQUEST"
	CodeValidationError = "VALIDATION_ERROR"
	CodeEncodingError   = "ENCODING_ERROR"
	CodeDecodingError   = "DECODING_ERROR"
	CodeParseError      = "PARSE_ERROR"
	CodeBindError       = "BIND_ERROR"
	CodePrepareError    = "PREPARE_ERROR"
	CodeMarshalError    = "MARSHAL_ERROR"
	CodeUnmarshalError  = "UNMARSHAL_ERROR"
)

// Resource-specific codes
const (
	CodeUserNotFound      = "USER_NOT_FOUND"
	CodeUserConflict      = "USER_CONFLICT"
	CodeUsernameConflict  = "USERNAME_CONFLICT"
	CodeEmailConflict     = "EMAIL_CONFLICT"
	CodeResourceNotFound  = "RESOURCE_NOT_FOUND"
	CodeResourceConflict  = "RESOURCE_CONFLICT"
	CodeRoleNotFound      = "ROLE_NOT_FOUND"
	CodeAuthNotFound      = "AUTH_NOT_FOUND"
	CodeTokenInvalid      = "TOKEN_INVALID"
	CodeTokenExpired      = "TOKEN_EXPIRED"
	CodePermissionDenied  = "PERMISSION_DENIED"
	CodeInvalidCredential = "INVALID_CREDENTIAL"
)

// DB-related
const (
	CodeDBError        = "DB_ERROR"
	CodeDBNoRows       = "DB_NO_ROWS"
	CodeDBConnFailed   = "DB_CONN_FAILED"
	CodeDBTxFailed     = "DB_TX_FAILED"
	CodeDBConstraint   = "DB_CONSTRAINT"
	CodeDBPrepareError = "DB_PREPARE_ERROR"
)
