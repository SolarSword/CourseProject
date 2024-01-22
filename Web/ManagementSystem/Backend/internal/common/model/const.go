package model

// error code
const (
	REQUEST_BODY_ERROR       = 1
	PERMISSON_ERROR          = 2
	COMPULSORY_FIELD_MISSING = 3
	DB_ERROR                 = 4
	INVALID_LIMIT            = 5
	INVALID_OFFSET           = 6
)

// error message
const (
	REQUEST_BODY_ERROR_MSG       = "request body error"
	PERMISSON_ERROR_MSG          = "permission error"
	COMPULSORY_FIELD_MISSING_MSG = "compulsory field missing"
	DB_ERROR_MSG                 = "database error"
	INVALID_LIMIT_MSG            = "invalid limit"
	INVALID_OFFSET_MSG           = "invalid offset"
)

// default query parameters
const (
	DEFAULT_LIMIT  = "10"
	DEFAULT_OFFSET = "0"
)
