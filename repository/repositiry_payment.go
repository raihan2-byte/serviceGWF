package repository

import (
	"payment-gwf/entity"

	"gorm.io/gorm"
)

type RepositoryPayment interface {
	FindAll() ([]*entity.Payment, error)
	Save(payment *entity.Payment) (*entity.Payment, error)
	FindById(ID int) (*entity.Payment, error)
	Update(payment *entity.Payment) (*entity.Payment, error)
	Delete(payment *entity.Payment) (*entity.Payment, error)
	FindAllByUserID(ID int) ([]*entity.Payment, error)
}

type repositoryPayment struct {
	db *gorm.DB
}

func NewRepositoryPayment(db *gorm.DB) *repositoryPayment {
	return &repositoryPayment{db}
}

func (r *repositoryPayment) FindAll() ([]*entity.Payment, error) {
	var payment []*entity.Payment

	err := r.db.Find(&payment).Error

	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (r *repositoryPayment) Save(payment *entity.Payment) (*entity.Payment, error) {
	err := r.db.Create(&payment).Error

	if err != nil {
		return payment, err
	}
	return payment, nil
}

func (r *repositoryPayment) FindById(ID int) (*entity.Payment, error) {
	var payment *entity.Payment

	err := r.db.Where("id = ?", ID).Find(&payment).Error

	if err != nil {
		return payment, err
	}
	return payment, nil
}

func (r *repositoryPayment) FindAllByUserID(ID int) ([]*entity.Payment, error) {
	var find []*entity.Payment

	err := r.db.Preload("User").Preload("Order").Where("user_id = ?", ID).Find(&find).Error
	if err != nil {
		return find, err
	}
	return find, nil
}

func (r *repositoryPayment) Update(payment *entity.Payment) (*entity.Payment, error) {
	err := r.db.Save(&payment).Error
	if err != nil {
		return payment, err
	}

	return payment, nil

}

func (r *repositoryPayment) Delete(payment *entity.Payment) (*entity.Payment, error) {
	err := r.db.Delete(&payment).Error
	if err != nil {
		return payment, err
	}

	return payment, nil
}
