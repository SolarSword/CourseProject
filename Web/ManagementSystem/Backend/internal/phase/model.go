package phase

type StartCourseSelectionPhaseRequest struct {
	RoleID  string `json:"role_id"`
	EndTime int64  `json:"end_time"`
}

type StartCourseSelectionPhaseResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}
