package dao

type Semester struct {
	Id        int `gorm:"primaryKey"`
	Type      int
	Semester  string
	StartTime int64
	EndTime   int64
}

func (Semester) TableName() string {
	return "semester_tab"
}

const (
	SEMESTER       = 1
	SHORT_SEMESTER = 2
)
