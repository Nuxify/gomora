package errors

const (
	// DatabaseError is the code for any database changes errors
	DatabaseError string = "DATABASE_ERROR"
	// DuplicateRecord is the code for duplicate records
	DuplicateRecord string = "DUPLICATE_RECORD"
	// ForbiddenAccess is the code for forbidden access
	ForbiddenAccess string = "FORBIDDEN_ACCESS"
	// HystrixTimeout is the code for hystrix timeouts
	HystrixTimeout string = "HYSTRIX_TIMEOUT"
	// InvalidRequestPayload is the code for binding errors
	InvalidRequestPayload string = "INVALID_REQUEST_PAYLOAD"
	// InvalidPayload is the code for payload not satisfying requirements
	InvalidPayload string = "INVALID_PAYLOAD"
	// MaximumLimitReached is the code when the max limit is reached
	MaximumLimitReached string = "MAX_LIMIT_REACHED"
	// MissingAPIEndpoint is the code for 404 API endpoints
	MissingAPIEndpoint string = "MISSING_API_ENDPOINT"
	// MissingConfiguration is the code for configurations not found error
	MissingConfiguration string = "MISSING_CONFIGURATION"
	// MissingRecord is the code for no record found
	MissingRecord string = "MISSING_RECORD"
	// ServerError is the code for server error
	ServerError string = "SERVER_ERROR"
	// ServerMaintenance is the code for server maintenance
	ServerMaintenance string = "SERVER_MAINTENANCE"
	// StorageUploadFailed is the code when storage upload (like to s3) failed
	StorageUploadFailed string = "STORAGE_UPLOAD_FAILED"
	// SystemScriptFailed is the code when scripts failed
	SystemScriptFailed string = "SYSTEM_SCRIPT_FAILED"
	// UnauthorizedAccess is the code for accessing restricted routes
	UnauthorizedAccess string = "UNAUTHORIZED_ACCESS"
)
