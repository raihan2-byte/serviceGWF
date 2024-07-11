package service

import (
	"payment-gwf/entity"
	"payment-gwf/input"
	"payment-gwf/repository"
)

type ServiceProduct interface {
	CreateProduct(input input.ProductInput) (*entity.Products, error)
	GetProducts() ([]*entity.Products, error)
	GetProduct(ID int) (*entity.Products, error)
	CreateProductImage(productID int, FileName string) error
	DeleteProduct(ID int) error
	UpdatedImage(fileName string, productID int) error
	UpdateProduct(ID int, input input.ProductInput) (*entity.Products, error)
	DeleteProductImages(productID int) error
}

type serviceProduct struct {
	repositoryProduct repository.RepositoryProduct
}

func NewServiceProduct(repositoryProduct repository.RepositoryProduct) *serviceProduct {
	return &serviceProduct{repositoryProduct}
}
func (s *serviceProduct) UpdatedImage(fileName string, productID int) error {
	productImage := entity.ProductImage{
		ProductID: productID,
		FileName:  fileName,
	}

	return s.repositoryProduct.UpdatedImage(productImage)
}

func (s *serviceProduct) CreateProductImage(productID int, FileName string) error {
	createProcut := entity.ProductImage{}

	createProcut.FileName = FileName
	createProcut.ProductID = productID

	err := s.repositoryProduct.CreateImage(createProcut)
	if err != nil {
		return err
	}
	return nil

}

func (s *serviceProduct) GetProductByCategory(ID int) ([]*entity.Products, error) {
	product, err := s.repositoryProduct.FindAllProductByCategory(ID)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (s *serviceProduct) GetProducts() ([]*entity.Products, error) {

	product, err := s.repositoryProduct.FindAll()
	if err != nil {
		return product, err
	}
	return product, nil
}

func (s *serviceProduct) UpdateProduct(ID int, input input.ProductInput) (*entity.Products, error) {
	find, err := s.repositoryProduct.FindById(ID)
	if err != nil {
		return find, err
	}

	find.Name = input.Name
	find.Price = input.Price
	find.Stock = input.Stock

	newProduct, err := s.repositoryProduct.Update(find)
	if err != nil {
		return newProduct, err
	}
	return newProduct, nil
}

func (s *serviceProduct) CreateProduct(input input.ProductInput) (*entity.Products, error) {
	product := &entity.Products{}

	product.Name = input.Name
	product.Price = input.Price
	product.Stock = input.Stock
	product.Description = input.Description

	newProduct, err := s.repositoryProduct.Save(product)
	if err != nil {
		return newProduct, err
	}
	return newProduct, nil
}

func (s *serviceProduct) DeleteProductImages(productID int) error {
	return s.repositoryProduct.DeleteImages(productID)
}

func (s *serviceProduct) GetProduct(ID int) (*entity.Products, error) {

	product, err := s.repositoryProduct.FindById(ID)

	if err != nil {
		return nil, err
	}

	if product.ID == 0 {
		return nil, err
	}

	return product, nil
}

func (s *serviceProduct) DeleteProduct(ID int) error {
	// Temukan produk berdasarkan ID
	product, err := s.repositoryProduct.FindById(ID)
	if err != nil {
		return err // Produk tidak ditemukan atau terjadi kesalahan lainnya
	}

	// Hapus gambar yang terkait dengan produk
	err = s.repositoryProduct.DeleteImages(ID)
	if err != nil {
		return err // Tangani kesalahan jika penghapusan gambar gagal
	}

	// Hapus produk dari basis data
	_, err = s.repositoryProduct.Delete(product)
	if err != nil {
		return err // Tangani kesalahan jika penghapusan produk gagal
	}

	return nil
}
