package repository

import (
	"payment-gwf/entity"

	"gorm.io/gorm"
)

type RepositoryStatusEkspedisi interface {
	FindAll() ([]*entity.StatusEkspedisi, error)
	FindAllByUserID(ID int) ([]*entity.StatusEkspedisi, error)
	FindByID(ID int) (*entity.StatusEkspedisi, error)
	Save(statusEkspedisi *entity.StatusEkspedisi) (*entity.StatusEkspedisi, error)
	Update(statusEkspedisi *entity.StatusEkspedisi) (*entity.StatusEkspedisi, error)
	Delete(statusEkspedisi *entity.StatusEkspedisi) (*entity.StatusEkspedisi, error)
}

type repositoryStatusEkspedisi struct {
	db *gorm.DB
}

func NewRepositoryStatusEkspedisi(db *gorm.DB) *repositoryStatusEkspedisi {
	return &repositoryStatusEkspedisi{db}
}

func (r *repositoryStatusEkspedisi) FindByID(ID int) (*entity.StatusEkspedisi, error) {
	var status *entity.StatusEkspedisi

	err := r.db.Where("id = ?", ID).Find(&status).Error

	if err != nil {
		return status, err
	}
	return status, nil
}

func (r *repositoryStatusEkspedisi) FindAllByUserID(ID int) ([]*entity.StatusEkspedisi, error) {
	var status []*entity.StatusEkspedisi

	err := r.db.Preload("Payment").Preload("Ongkir").Preload("Order").Preload("Order.Items").Preload("Order.Items.Product").Preload("User").Where("user_id = ?", ID).Find(&status).Error
	if err != nil {
		return status, err
	}
	return status, nil
}

func (r *repositoryStatusEkspedisi) FindAll() ([]*entity.StatusEkspedisi, error) {
	var status []*entity.StatusEkspedisi

	err := r.db.Preload("Payment").Preload("Ongkir").Preload("Order").Preload("Order.Items").Preload("Order.Items.Product").Preload("User").Find(&status).Error

	if err != nil {
		return status, err
	}

	return status, nil
}

func (r *repositoryStatusEkspedisi) Save(status *entity.StatusEkspedisi) (*entity.StatusEkspedisi, error) {
	err := r.db.Create(&status).Error

	if err != nil {
		return status, err
	}
	return status, nil
}

func (r *repositoryStatusEkspedisi) FindById(ID int) (*entity.StatusEkspedisi, error) {
	var status *entity.StatusEkspedisi

	err := r.db.Preload("Ongkir").Where("id = ?", ID).Find(&status).Error

	if err != nil {
		return status, err
	}
	return status, nil
}

func (r *repositoryStatusEkspedisi) Update(status *entity.StatusEkspedisi) (*entity.StatusEkspedisi, error) {
	err := r.db.Save(&status).Error
	if err != nil {
		return status, err
	}

	return status, nil

}

func (r *repositoryStatusEkspedisi) Delete(status *entity.StatusEkspedisi) (*entity.StatusEkspedisi, error) {
	err := r.db.Delete(&status).Error
	if err != nil {
		return status, err
	}

	return status, nil
}
