package repository

import (
	"payment-gwf/entity"

	"gorm.io/gorm"
)

type RepositoryRajaOngkir interface {
	// FindAll() ([]*entity.ApplyShippingResponse, error)
	Save(address *entity.ApplyShippingResponse) (*entity.ApplyShippingResponse, error)
	FindById(ID int) (*entity.ApplyShippingResponse, error)
	// Update(address *entity.ApplyShippingResponse) (*entity.ApplyShippingResponse, error)
	// Delete(address *entity.ApplyShippingResponse) (*entity.ApplyShippingResponse, error)
	// FindAllAddressByCategory(ID int) ([]*entity.ApplyShippingResponse, error)
}

type repositoryRajaOngkir struct {
	db *gorm.DB
}

func NewRepositoryRajaOngkir(db *gorm.DB) *repositoryRajaOngkir {
	return &repositoryRajaOngkir{db}
}

func (r *repositoryRajaOngkir) Save(address *entity.ApplyShippingResponse) (*entity.ApplyShippingResponse, error) {
	err := r.db.Create(&address).Error

	if err != nil {
		return address, err
	}
	return address, nil
}

func (r *repositoryRajaOngkir) FindById(ID int) (*entity.ApplyShippingResponse, error) {
	var ongkir *entity.ApplyShippingResponse

	err := r.db.Where("id = ?", ID).Find(&ongkir).Error

	if err != nil {
		return ongkir, err
	}
	return ongkir, nil
}
