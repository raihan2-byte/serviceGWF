package service

import (
	"errors"
	"payment-gwf/entity"
	"payment-gwf/input"
	"payment-gwf/repository"

	"github.com/google/uuid"
)

type ServiceMakeDonation interface {
	CreateDonation(userID int, input input.MakeDonationInput) (*entity.MakeDonation, error)
	GetDonations() ([]*entity.MakeDonation, error)
	GetDonation(userID int) (*entity.MakeDonation, error)
	DeleteDonation(ID string) (*entity.MakeDonation, error)
}

type serviceMakeDonation struct {
	repositoryMakeDonation    repository.RepositoryMakeDonation
	repositoryUser            repository.RepositoryUser
	repositoryPaymentDonation repository.RepositoryPaymentDonation
}

func NewServiceMakeDonation(repositoryMakeDonation repository.RepositoryMakeDonation, repositoryUser repository.RepositoryUser, repositoryPaymentDonation repository.RepositoryPaymentDonation) *serviceMakeDonation {
	return &serviceMakeDonation{repositoryMakeDonation, repositoryUser, repositoryPaymentDonation}
}

func (s *serviceMakeDonation) GetDonations() ([]*entity.MakeDonation, error) {

	donations, err := s.repositoryMakeDonation.FindAll()
	if err != nil {
		return donations, err
	}
	return donations, nil
}

func (s *serviceMakeDonation) CreateDonation(userID int, input input.MakeDonationInput) (*entity.MakeDonation, error) {
	donations := &entity.MakeDonation{}

	getUser, err := s.repositoryUser.FindById(userID)
	if err != nil {
		return nil, err
	}

	donations.Name = input.Name
	donations.Amount = input.Amount
	if input.Amount < 5000 {
		return nil, errors.New("input must up to 5000")
	}
	donations.Message = input.Message
	donations.UserID = getUser.ID
	donations.ID = uuid.New().String()

	newDonations, err := s.repositoryMakeDonation.Save(donations)
	if err != nil {
		return newDonations, err
	}

	// savePaymentDonation := &entity.PaymentDonation{}
	// savePaymentDonation.StatusPayment = "pending"
	// savePaymentDonation.MakeDonationID = newDonations.ID
	// savePaymentDonation.UserID = newDonations.UserID

	// _, err = s.repositoryPaymentDonation.Save(savePaymentDonation)
	// if err != nil {
	// 	return nil, err
	// }
	return newDonations, nil
}

func (s *serviceMakeDonation) GetDonation(userID int) (*entity.MakeDonation, error) {

	findUser, err := s.repositoryUser.FindById(userID)
	if err != nil {
		return nil, err
	}

	donations, err := s.repositoryMakeDonation.FindDonationByUserID(findUser.ID)
	if err != nil {
		return donations, err
	}
	return donations, nil
}

func (s *serviceMakeDonation) DeleteDonation(ID string) (*entity.MakeDonation, error) {

	donations, err := s.repositoryMakeDonation.FindById(ID)
	if err != nil {
		return donations, err
	}
	donationsDel, err := s.repositoryMakeDonation.Delete(donations)

	if err != nil {
		return donationsDel, err
	}
	return donationsDel, nil

}
