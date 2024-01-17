package phase

const (
	PHASE_ERROR = 1001
)

const (
	PHASE_ERROR_MSG = "the current phase is not course selection phase"
)

type StartCourseSelectionPhaseRequest struct {
	RoleID  string `json:"role_id"`
	EndTime int64  `json:"end_time"`
}

type StartCourseSelectionPhaseResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

type EndCourseSelectionRequest struct {
	RoleID string `json:"role_id"`
}

type EndCourseSelectionResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

type GetPhaseResponse struct {
	phaseType int   `json:"type"`
	endTime   int64 `json:"end_time"`
}
