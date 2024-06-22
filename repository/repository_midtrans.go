package repository

import (
	"payment-gwf/entity"

	"gorm.io/gorm"
)

type RepositoryMidtrans interface {
	FindAll() ([]*entity.DoPayment, error)
	Save(product *entity.DoPayment) (*entity.DoPayment, error)
	FindById(ID int) (*entity.DoPayment, error)
	Update(product *entity.DoPayment) (*entity.DoPayment, error)
	Delete(product *entity.DoPayment) (*entity.DoPayment, error)
}

type repositoryMidtrans struct {
	db *gorm.DB
}

func NewRepositoryMidtrans(db *gorm.DB) *repositoryMidtrans {
	return &repositoryMidtrans{db}
}

func (r *repositoryMidtrans) FindAll() ([]*entity.DoPayment, error) {
	var product []*entity.DoPayment

	err := r.db.Preload("FileName").Find(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repositoryMidtrans) Save(product *entity.DoPayment) (*entity.DoPayment, error) {
	err := r.db.Create(&product).Error

	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repositoryMidtrans) FindById(ID int) (*entity.DoPayment, error) {
	var product *entity.DoPayment

	err := r.db.Where("id = ?", ID).Find(&product).Error

	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repositoryMidtrans) Update(product *entity.DoPayment) (*entity.DoPayment, error) {
	err := r.db.Save(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil

}

func (r *repositoryMidtrans) Delete(product *entity.DoPayment) (*entity.DoPayment, error) {
	err := r.db.Delete(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}
