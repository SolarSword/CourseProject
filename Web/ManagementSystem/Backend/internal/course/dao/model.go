package dao

type Course struct {
	CourseId   string `gorm:"primaryKey"`
	CourseName string
	CollegeId  string
	Credit     int
	Brief      string
}

func (Course) TableName() string {
	return "course_tab"
}

type CourseModule struct {
	CourseModuleId   string `gorm:"primaryKey"`
	CourseId         string
	ProfessorId      string
	TaId             string
	Semester         string
	Classroom        string
	ClassPeriodStart string
	ClassPeriodEnd   string
	Duration         int
	CourseCapacity   int
	MinStuNum        int
	ScoreRatio       string
	Status           int
}

func (CourseModule) TableName() string {
	return "course_module_tab"
}

const (
	SELECTION_IN_PROGRESS = 1
	NORMAL_TEACHING       = 2
	COURSE_ENDED          = 3
	CANCELED              = 4
	REVIEWING             = 5
)

type CourseModuleStu struct {
	Id             int `gorm:"primaryKey"`
	CourseModuleId string
	CourseId       string
	StuId          string
	Scroes         string
	FinalScore     int
	Status         int
}

func (CourseModuleStu) TableName() string {
	return "course_module_stu_tab"
}

const (
	SELECTING        = 1
	ENROLLED         = 2
	SELECTION_FAILED = 3
	ENDED            = 4
	FAILED           = 5
)
