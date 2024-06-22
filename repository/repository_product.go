package repository

import (
	"payment-gwf/entity"

	"gorm.io/gorm"
)

type RepositoryProduct interface {
	FindAll() ([]*entity.Products, error)
	Save(product *entity.Products) (*entity.Products, error)
	FindById(ID int) (*entity.Products, error)
	CreateImage(product entity.ProductImage) error
	UpdatedImage(product entity.ProductImage) error
	DeleteImages(productID int) error
	Update(product *entity.Products) (*entity.Products, error)
	Delete(product *entity.Products) (*entity.Products, error)
	FindAllProductByCategory(ID int) ([]*entity.Products, error)
}

type repositoryProduct struct {
	db *gorm.DB
}

func NewRepositoryProduct(db *gorm.DB) *repositoryProduct {
	return &repositoryProduct{db}
}

func (r *repositoryProduct) UpdatedImage(product entity.ProductImage) error {
	return r.db.Save(&product).Error
}

func (r *repositoryProduct) FindAllProductByCategory(ID int) ([]*entity.Products, error) {
	var product []*entity.Products

	err := r.db.Preload("Category").Where("category_id = ? ", ID).Find(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repositoryProduct) CreateImage(product entity.ProductImage) error {
	err := r.db.Preload("FileName").Create(&product).Error
	return err

}

func (r *repositoryProduct) DeleteImages(productID int) error {
	err := r.db.Where("product_id = ?", productID).Delete(&entity.ProductImage{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repositoryProduct) FindAll() ([]*entity.Products, error) {
	var product []*entity.Products

	err := r.db.Preload("FileName").Find(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repositoryProduct) Save(product *entity.Products) (*entity.Products, error) {
	err := r.db.Preload("FileName").Create(&product).Error

	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repositoryProduct) FindById(ID int) (*entity.Products, error) {
	var product *entity.Products

	err := r.db.Preload("FileName").Where("id = ?", ID).Find(&product).Error

	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repositoryProduct) Update(product *entity.Products) (*entity.Products, error) {
	err := r.db.Save(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil

}

func (r *repositoryProduct) Delete(product *entity.Products) (*entity.Products, error) {
	err := r.db.Preload("FileName").Delete(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}
