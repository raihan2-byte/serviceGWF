package repository

import (
	"payment-gwf/entity"

	"gorm.io/gorm"
)

type RepositoryAddress interface {
	FindAll() ([]*entity.BuyerAddress, error)
	Save(address *entity.BuyerAddress) (*entity.BuyerAddress, error)
	FindById(ID int) (*entity.BuyerAddress, error)
	Update(address *entity.BuyerAddress) (*entity.BuyerAddress, error)
	Delete(address *entity.BuyerAddress) (*entity.BuyerAddress, error)
	FindAllAddressByCategory(ID int) ([]*entity.BuyerAddress, error)
}

type repositoryAddress struct {
	db *gorm.DB
}

func NewRepositoryAddress(db *gorm.DB) *repositoryAddress {
	return &repositoryAddress{db}
}

func (r *repositoryAddress) FindAllAddressByCategory(ID int) ([]*entity.BuyerAddress, error) {
	var address []*entity.BuyerAddress

	err := r.db.Preload("Category").Where("category_id = ? ", ID).Find(&address).Error

	if err != nil {
		return address, err
	}

	return address, nil
}

func (r *repositoryAddress) FindAll() ([]*entity.BuyerAddress, error) {
	var address []*entity.BuyerAddress

	err := r.db.Find(&address).Error

	if err != nil {
		return address, err
	}

	return address, nil
}

func (r *repositoryAddress) Save(address *entity.BuyerAddress) (*entity.BuyerAddress, error) {
	err := r.db.Create(&address).Error

	if err != nil {
		return address, err
	}
	return address, nil
}

func (r *repositoryAddress) FindById(ID int) (*entity.BuyerAddress, error) {
	var address *entity.BuyerAddress

	err := r.db.Where("id = ?", ID).Find(&address).Error

	if err != nil {
		return address, err
	}
	return address, nil
}

func (r *repositoryAddress) Update(address *entity.BuyerAddress) (*entity.BuyerAddress, error) {
	err := r.db.Save(&address).Error
	if err != nil {
		return address, err
	}

	return address, nil

}

func (r *repositoryAddress) Delete(address *entity.BuyerAddress) (*entity.BuyerAddress, error) {
	err := r.db.Delete(&address).Error
	if err != nil {
		return address, err
	}

	return address, nil
}
