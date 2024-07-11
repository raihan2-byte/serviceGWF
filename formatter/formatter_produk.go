package formatter

import (
	"payment-gwf/entity"
	"time"
)

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
	Description string    `json:"description"`
	FileName    []string  `json:"file_names"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FormatterProduct(produk *entity.Products) Product {
	fileNames := make([]string, len(produk.FileName))
	for i, image := range produk.FileName {
		fileNames[i] = image.FileName
	}

	return Product{
		ID:          produk.ID,
		Name:        produk.Name,
		Price:       produk.Price,
		Stock:       produk.Stock,
		Description: produk.Description,
		FileName:    fileNames,
		CreatedAt:   produk.CreatedAt,
		UpdatedAt:   produk.UpdatedAt,
	}
}

func FormatterGetProducts(produk []*entity.Products) []Product {
	produkGetFormatter := []Product{}

	for _, product := range produk {
		produkFormatter := FormatterProduct(product)
		produkGetFormatter = append(produkGetFormatter, produkFormatter)
	}

	return produkGetFormatter
}
