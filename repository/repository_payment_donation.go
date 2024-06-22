package repository

import (
	"payment-gwf/entity"

	"gorm.io/gorm"
)

type RepositoryPaymentDonation interface {
	FindAll() ([]*entity.PaymentDonation, error)
	Save(payment *entity.PaymentDonation) (*entity.PaymentDonation, error)
	FindById(ID string) (*entity.PaymentDonation, error)
	Update(payment *entity.PaymentDonation) (*entity.PaymentDonation, error)
	Delete(payment *entity.PaymentDonation) (*entity.PaymentDonation, error)
	FindAllByUserID(ID int) ([]*entity.PaymentDonation, error)
	FindByTransactionID(ID string) (*entity.PaymentDonation, error)
	// UpdateStatusByID(paymentID int, statusPayment string) (*entity.Payment, error)
}

type repositoryPaymentDonation struct {
	db *gorm.DB
}

func NewRepositoryPaymentDonation(db *gorm.DB) *repositoryPaymentDonation {
	return &repositoryPaymentDonation{db}
}

// func (r *repositoryPaymentDonation) UpdateStatusByID(paymentID int, statusPayment string) (*entity.Payment, error) {
// 	payment := &entity.Payment{}
// 	err := r.db.Model(payment).Where("id = ?", paymentID).Update("status_payment", statusPayment).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Retrieve the updated payment record to return
// 	err = r.db.Where("id = ?", paymentID).First(payment).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return payment, nil
// }

func (r *repositoryPaymentDonation) FindByTransactionID(ID string) (*entity.PaymentDonation, error) {
	var payment *entity.PaymentDonation

	err := r.db.Where("transaction_id = ?", ID).Find(&payment).Error

	if err != nil {
		return payment, err
	}
	return payment, nil
}

func (r *repositoryPaymentDonation) FindAll() ([]*entity.PaymentDonation, error) {
	var payment []*entity.PaymentDonation

	err := r.db.Preload("Order").Preload("User").Preload("Order.Items").Preload("Order.Items.Product").Preload("Order.Ongkir").Find(&payment).Error

	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (r *repositoryPaymentDonation) Save(payment *entity.PaymentDonation) (*entity.PaymentDonation, error) {
	err := r.db.Create(&payment).Error

	if err != nil {
		return payment, err
	}
	return payment, nil
}

func (r *repositoryPaymentDonation) FindById(ID string) (*entity.PaymentDonation, error) {
	var payment *entity.PaymentDonation

	err := r.db.Preload("Order").Preload("User").Preload("Order.Items").Preload("Order.Items.Product").Preload("Order.Ongkir").Where("id = ?", ID).Find(&payment).Error

	if err != nil {
		return payment, err
	}
	return payment, nil
}

func (r *repositoryPaymentDonation) FindAllByUserID(ID int) ([]*entity.PaymentDonation, error) {
	var find []*entity.PaymentDonation

	err := r.db.Preload("MakeDonation").Preload("User").Where("user_id = ?", ID).Find(&find).Error
	if err != nil {
		return find, err
	}
	return find, nil
}

func (r *repositoryPaymentDonation) Update(payment *entity.PaymentDonation) (*entity.PaymentDonation, error) {
	err := r.db.Save(&payment).Error
	if err != nil {
		return payment, err
	}

	return payment, nil

}

func (r *repositoryPaymentDonation) Delete(payment *entity.PaymentDonation) (*entity.PaymentDonation, error) {
	err := r.db.Delete(&payment).Error
	if err != nil {
		return payment, err
	}

	return payment, nil
}
