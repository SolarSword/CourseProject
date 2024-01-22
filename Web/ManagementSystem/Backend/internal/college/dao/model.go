package dao

type College struct {
	CollegeId   string `gorm:"primaryKey"`
	CollegeName string
}

func (College) TableName() string {
	return "college_tab"
}
