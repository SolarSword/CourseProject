package dao

type User struct {
	UserId   int `gorm:"primaryKey"`
	UserName string
	PassWord string
}

func (User) TableName() string {
	return "user_tab"
}

type Product struct {
	Id          int `gorm:"primaryKey"`
	ProductId   int
	ProductName string
	ShopId      int
	Price       int
	Stock       int
	Sold        int
	Status      int
	Description int
}

func (Product) TableName() string {
	return "product_tab"
}
