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
	DeleteProduct(ID int) (*entity.Products, error)
	UpdateProduct(ID int, input input.ProductInput) (*entity.Products, error)
}

type serviceProduct struct {
	repositoryProduct repository.RepositoryProduct
}

func NewServiceProduct(repositoryProduct repository.RepositoryProduct) *serviceProduct {
	return &serviceProduct{repositoryProduct}
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

	find.Name = input.Title
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

	product.Name = input.Title
	product.Price = input.Price
	product.Stock = input.Stock

	newProduct, err := s.repositoryProduct.Save(product)
	if err != nil {
		return newProduct, err
	}
	return newProduct, nil
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

func (s *serviceProduct) DeleteProduct(ID int) (*entity.Products, error) {

	product, err := s.repositoryProduct.FindById(ID)
	if err != nil {
		return product, err
	}
	productDel, err := s.repositoryProduct.Delete(product)

	if err != nil {
		return productDel, err
	}
	return productDel, nil

}
