package semester

const (
	INVALID_TIME_DURATION = 2001
	INVALID_SEMESTER      = 2002
	DUPLICATED_SEMESTER   = 2003
)

const (
	INVALID_TIME_DURATION_MSG = "invalid time duration"
	INVALID_SEMESTER_MSG      = "invalid semester"
	DUPLICATED_SEMESTER_MSG   = "duplicated semester"
)

type CreateSemesterRequest struct {
	RoleID    string `json:"role_id"`
	Semester  string `json:"semester"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	Type      int    `json:"type"`
}

type CreateSemesterResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

const (
	DAY_HOURS                = 24
	WEEK_DAYS                = 7
	SEMESTER_MIN             = 16
	SEMESTER_MAX             = 20
	SHORT_SEMESTER_MIN       = 2
	SHORT_SEMESTER_MAX       = 6
	SEMESTER_MIN_HOURS       = DAY_HOURS * WEEK_DAYS * SEMESTER_MIN
	SEMESTER_MAX_HOURS       = DAY_HOURS * WEEK_DAYS * SEMESTER_MAX
	SHORT_SEMESTER_MIN_HOURS = DAY_HOURS * WEEK_DAYS * SHORT_SEMESTER_MIN
	SHORT_SEMESTER_MAX_HOURS = DAY_HOURS * WEEK_DAYS * SHORT_SEMESTER_MAX
)
