package dao

type User struct {
	UserId   int `gorm:"primaryKey"`
	Password string
}

func (User) TableName() string {
	return "user_tab"
}

type Role struct {
	RoleId         string `gorm:"primaryKey"`
	UserId         int
	CollegeId      string
	Name           string
	Gender         int
	Type           int
	Email          string
	Grade          int
	EnrollmentYear int
	Status         int
}

func (Role) TableName() string {
	return "role_tab"
}
