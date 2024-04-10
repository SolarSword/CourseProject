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
	Id          int64 `gorm:"primaryKey"`
	ProductId   int64
	ProductName string
	ShopId      int64
	Price       int32
	Stock       int32
	Sold        int32
	Status      int32
	Description string
}

func (Product) TableName() string {
	return "product_tab"
}
