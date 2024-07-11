package repository

import (
	"payment-gwf/entity"

	"gorm.io/gorm"
)

type RepositoryPayment interface {
	FindAll() ([]*entity.Payment, error)
	Save(payment *entity.Payment) (*entity.Payment, error)
	FindById(ID int) (*entity.Payment, error)
	FindByTransactionID(ID string) (*entity.Payment, error)
	Update(payment *entity.Payment) (*entity.Payment, error)
	Delete(payment *entity.Payment) (*entity.Payment, error)
	FindAllByUserID(ID int) ([]*entity.Payment, error)
	SaveMidtrans(payment *entity.DoPayment) (*entity.DoPayment, error)
	FindByOrderId(ID string) (*entity.Payment, error)
}

type repositoryPayment struct {
	db *gorm.DB
}

func NewRepositoryPayment(db *gorm.DB) *repositoryPayment {
	return &repositoryPayment{db}
}

func (r *repositoryPayment) SaveMidtrans(payment *entity.DoPayment) (*entity.DoPayment, error) {
	err := r.db.Create(&payment).Error

	if err != nil {
		return payment, err
	}
	return payment, nil
}

func (r *repositoryPayment) FindAll() ([]*entity.Payment, error) {
	var payment []*entity.Payment

	err := r.db.Preload("Order").Preload("User").Preload("Order.Items").Preload("Order.Items.Product").Preload("Order.Ongkir").Find(&payment).Error

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

func (r *repositoryPayment) FindByTransactionID(ID string) (*entity.Payment, error) {
	var payment *entity.Payment

	err := r.db.Where("transaction_id = ?", ID).Find(&payment).Error

	if err != nil {
		return payment, err
	}
	return payment, nil
}

func (r *repositoryPayment) FindById(ID int) (*entity.Payment, error) {
	var payment *entity.Payment

	err := r.db.Preload("Order").Preload("User").Preload("Order.Items").Preload("Order.Items.Product").Preload("Order.Ongkir").Where("id = ?", ID).Find(&payment).Error

	if err != nil {
		return payment, err
	}
	return payment, nil
}

func (r *repositoryPayment) FindAllByUserID(ID int) ([]*entity.Payment, error) {
	var find []*entity.Payment

	err := r.db.Preload("Order").Preload("User").Preload("Order.Items").Preload("Order.Items.Product").Preload("Order.Ongkir").Where("user_id = ?", ID).Find(&find).Error
	if err != nil {
		return find, err
	}
	return find, nil
}

// func (r *repositoryPayment) Update(payment *entity.Payment) (*entity.Payment, error) {
// 	err := r.db.Model(&entity.Payment{}).Where("order_id = ? ",
// 		payment.OrderID).Update("transaction_id", payment.TransactionID).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return payment, nil

// }

func (r *repositoryPayment) FindByOrderId(ID string) (*entity.Payment, error) {
	var payment *entity.Payment

	err := r.db.Preload("User").Preload("Order").Where("order_id = ?", ID).Find(&payment).Error

	if err != nil {
		return payment, err
	}
	return payment, nil
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
