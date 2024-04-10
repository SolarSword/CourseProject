package dao

import (
	"log"

	DB "course.project/authentication/internal/common/db"
)

func IsValidUser(userName string, passWord string) bool {
	user := &User{}
	DB.Db.GetDB().Where(&User{UserName: userName}).First(&user)
	return passWord == user.PassWord
}

// in actual project, it would be better that a DAO only works
// for one "component"
func AddProductToTable(product Product) error {
	res := DB.Db.GetDB().Create(&product)
	if res.Error != nil {
		log.Fatalf("fail to create product: %v, err: %v", product, res.Error)
		return res.Error
	}
	return nil
}

func UpdateProductByProductID(product Product) error {
	res := DB.Db.GetDB().Where("product_id = ?", product.ProductId).
		Updates(product)
	if res.Error != nil {
		log.Fatalf("fail to update product: %v, err: %v", product, res.Error)
		return res.Error
	}
	return nil
}

func BatchGetProductsWithPagination(limit int, offset int) ([]Product, error) {
	products := []Product{}
	res := DB.Db.GetDB().Find(&products).Limit(limit).Offset(offset)
	if res.Error != nil {
		log.Fatalf("fail to get products, err: %v", res.Error)
		return nil, res.Error
	}
	return products, nil
}

func DeleteProductWithProductID(productID int64) error {
	res := DB.Db.GetDB().Where("product_id = ?", productID).Delete(&Product{})
	if res.Error != nil {
		log.Fatalf("fail to delete product %v, err: %v", productID, res.Error)
		return res.Error
	}
	return nil
}
