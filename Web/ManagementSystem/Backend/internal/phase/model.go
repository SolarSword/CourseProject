package phase

// error code
const (
	REQUEST_BODY_ERROR = 1
	PERMISSON_ERROR    = 2
)

// error message
const (
	REQUEST_BODY_ERROR_MSG = "request body error"
	PERMISSON_ERROR_MSG    = "permission error"
)

type StartCourseSelectionPhaseRequest struct {
	RoleID  string `json:"role_id"`
	EndTime int64  `json:"end_time"`
}

type StartCourseSelectionPhaseResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}
