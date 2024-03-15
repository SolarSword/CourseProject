package product_business

import (
	cm "course.project/authentication/internal/common/common"
)

type Product struct {
	ProductId   int64  `json:"product_id"`
	ProductName string `json:"product_name"`
	ShopId      int64  `json:"shop_id"`
	Price       int64  `json:"price"`
	Stock       int    `json:"stock"`
	Sold        int    `json:"sold"`
	Status      int    `json:"status"`
	Description string `json:"description"`
}

type AddProductRequest struct {
	Product Product `json:"product"`
}

type AddProductResponse struct {
	Error cm.Error `json:"error"`
}

type GetProductsResponse struct {
	Products []Product `json:"products"`
	Error    cm.Error  `json:"error"`
}

type EditProductRequest struct {
	Product Product `json:"product"`
}

type EditProductResponse struct {
	Error cm.Error `json:"error"`
}

type DeleteProductRequest struct {
	ProductId int64 `json:"product_id"`
}

type DeleteProductResponse struct {
	Error cm.Error `json:"error"`
}
