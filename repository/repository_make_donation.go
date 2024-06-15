package repository

import (
	"payment-gwf/entity"

	"gorm.io/gorm"
)

type RepositoryMakeDonation interface {
	FindAll() ([]*entity.MakeDonation, error)
	Save(donations *entity.MakeDonation) (*entity.MakeDonation, error)
	FindById(ID int) (*entity.MakeDonation, error)
	FindDonationByUserID(ID int) (*entity.MakeDonation, error)
	Update(donations *entity.MakeDonation) (*entity.MakeDonation, error)
	Delete(product *entity.MakeDonation) (*entity.MakeDonation, error)
}

type repositoryMakeDonation struct {
	db *gorm.DB
}

func NewRepositoryMakeDonation(db *gorm.DB) *repositoryMakeDonation {
	return &repositoryMakeDonation{db}
}

func (r *repositoryMakeDonation) FindDonationByUserID(ID int) (*entity.MakeDonation, error) {
	var donations *entity.MakeDonation

	err := r.db.Preload("User").Where("user_id = ?", ID).Find(&donations).Error

	if err != nil {
		return donations, err
	}
	return donations, nil
}

func (r *repositoryMakeDonation) FindAll() ([]*entity.MakeDonation, error) {
	var donations []*entity.MakeDonation

	err := r.db.Find(&donations).Error

	if err != nil {
		return donations, err
	}

	return donations, nil
}

func (r *repositoryMakeDonation) Save(donations *entity.MakeDonation) (*entity.MakeDonation, error) {
	err := r.db.Create(&donations).Error

	if err != nil {
		return donations, err
	}
	return donations, nil
}

func (r *repositoryMakeDonation) FindById(ID int) (*entity.MakeDonation, error) {
	var donations *entity.MakeDonation

	err := r.db.Where("id = ?", ID).Find(&donations).Error

	if err != nil {
		return donations, err
	}
	return donations, nil
}

func (r *repositoryMakeDonation) Update(donations *entity.MakeDonation) (*entity.MakeDonation, error) {
	err := r.db.Save(&donations).Error
	if err != nil {
		return donations, err
	}

	return donations, nil

}

func (r *repositoryMakeDonation) Delete(donations *entity.MakeDonation) (*entity.MakeDonation, error) {
	err := r.db.Delete(&donations).Error
	if err != nil {
		return donations, err
	}

	return donations, nil
}
