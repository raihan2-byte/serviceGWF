package repository

import (
	"payment-gwf/entity"

	"gorm.io/gorm"
)

type RepositoryPaymentDetails interface {
	FindAll() ([]*entity.PaymentDetails, error)
	Save(payment *entity.PaymentDetails) (*entity.PaymentDetails, error)
	FindById(ID int) (*entity.PaymentDetails, error)
	Update(payment *entity.PaymentDetails) (*entity.PaymentDetails, error)
	Delete(payment *entity.PaymentDetails) (*entity.PaymentDetails, error)
	FindAllByUserID(ID int) ([]*entity.PaymentDetails, error)
	FindByTransactionID(ID string) (*entity.PaymentDetails, error)
}

type repositoryPaymentDetails struct {
	db *gorm.DB
}

func NewRepositoryPaymentDetails(db *gorm.DB) *repositoryPaymentDetails {
	return &repositoryPaymentDetails{db}
}

func (r *repositoryPaymentDetails) FindByTransactionID(ID string) (*entity.PaymentDetails, error) {
	var payment *entity.PaymentDetails

	err := r.db.Where("transaction_id = ?", ID).Find(&payment).Error

	if err != nil {
		return payment, err
	}
	return payment, nil
}

func (r *repositoryPaymentDetails) FindAll() ([]*entity.PaymentDetails, error) {
	var payment []*entity.PaymentDetails

	err := r.db.Preload("Order").Preload("User").Preload("Order.Items").Preload("Order.Items.Product").Preload("Order.Ongkir").Find(&payment).Error

	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (r *repositoryPaymentDetails) Save(payment *entity.PaymentDetails) (*entity.PaymentDetails, error) {
	err := r.db.Create(&payment).Error

	if err != nil {
		return payment, err
	}
	return payment, nil
}

func (r *repositoryPaymentDetails) FindById(ID int) (*entity.PaymentDetails, error) {
	var payment *entity.PaymentDetails

	err := r.db.Preload("Order").Preload("User").Preload("Order.Items").Preload("Order.Items.Product").Preload("Order.Ongkir").Where("id = ?", ID).Find(&payment).Error

	if err != nil {
		return payment, err
	}
	return payment, nil
}

func (r *repositoryPaymentDetails) FindAllByUserID(ID int) ([]*entity.PaymentDetails, error) {
	var find []*entity.PaymentDetails

	err := r.db.Preload("MakeDonation").Preload("User").Where("user_id = ?", ID).Find(&find).Error
	if err != nil {
		return find, err
	}
	return find, nil
}

func (r *repositoryPaymentDetails) Update(payment *entity.PaymentDetails) (*entity.PaymentDetails, error) {
	err := r.db.Save(&payment).Error
	if err != nil {
		return payment, err
	}

	return payment, nil

}

func (r *repositoryPaymentDetails) Delete(payment *entity.PaymentDetails) (*entity.PaymentDetails, error) {
	err := r.db.Delete(&payment).Error
	if err != nil {
		return payment, err
	}

	return payment, nil
}
