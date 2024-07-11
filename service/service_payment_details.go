package service

import (
	"fmt"
	"payment-gwf/entity"
	"payment-gwf/repository"

	"github.com/midtrans/midtrans-go/coreapi"
)

type ServicePaymentDetails interface {
	HandleNotificationPaymentDetails(req *entity.MidtransNotificationRequest) error
	updatePayment(req *entity.MidtransNotificationRequest) error
	updatePaymentDonation(req *entity.MidtransNotificationRequest) error
}

type servicePaymentDetails struct {
	repositoryPaymentDetails  repository.RepositoryPaymentDetails
	repositoryPayment         repository.RepositoryPayment
	repositoryPaymentDonation repository.RepositoryPaymentDonation
	repositoryUser            repository.RepositoryUser
	repositoryOrder           repository.RepositoryOrder
	repositoryMakeDonation    repository.RepositoryMakeDonation
	midtransGateway           *midtransGateway
}

func NewServicePaymentDetails(repositoryPaymentDetails repository.RepositoryPaymentDetails, repositoryPayment repository.RepositoryPayment, repositoryPaymentDonation repository.RepositoryPaymentDonation, repositoryUser repository.RepositoryUser, repositoryOrder repository.RepositoryOrder, repositoryMakeDonation repository.RepositoryMakeDonation, midtransGateway *midtransGateway) *servicePaymentDetails {
	return &servicePaymentDetails{
		repositoryPaymentDetails:  repositoryPaymentDetails,
		repositoryPayment:         repositoryPayment,
		repositoryPaymentDonation: repositoryPaymentDonation,
		repositoryUser:            repositoryUser,
		repositoryOrder:           repositoryOrder,
		repositoryMakeDonation:    repositoryMakeDonation,
		midtransGateway:           midtransGateway,
	}
}

func mapChargeToResponsePaymentDetails(resp *coreapi.ChargeResponse) (*entity.DoPayment, error) {
	if resp == nil {
		return nil, fmt.Errorf("nil response received")
	}

	var vaNumbers []entity.VaNumber
	for _, va := range resp.VaNumbers {
		vaNumbers = append(vaNumbers, entity.VaNumber{
			Bank:     va.Bank,
			VaNumber: va.VANumber,
		})
	}

	return &entity.DoPayment{
		TransactionID: resp.TransactionID,
		OrderID:       resp.OrderID,
		GrossAmount:   resp.GrossAmount,
		VaNumbers:     vaNumbers,
		MerchantID:    "G317569034",
	}, nil
}

func (s *servicePaymentDetails) updatePayment(req *entity.MidtransNotificationRequest) error {
	findPayment, err := s.repositoryPayment.FindByTransactionID(req.TransactionID)
	if err != nil {
		return err
	}

	findPayment.StatusPayment = req.TransactionStatus
	_, err = s.repositoryPayment.Update(findPayment)
	if err != nil {
		return err
	}

	return nil
}

func (s *servicePaymentDetails) updatePaymentDonation(req *entity.MidtransNotificationRequest) error {
	findPaymentDonation, err := s.repositoryPaymentDonation.FindByTransactionID(req.TransactionID)
	if err != nil {
		return err
	}

	findPaymentDonation.StatusPayment = req.TransactionStatus
	_, err = s.repositoryPaymentDonation.Update(findPaymentDonation)
	if err != nil {
		return err
	}

	return nil
}

func (s *servicePaymentDetails) HandleNotificationPaymentDetails(req *entity.MidtransNotificationRequest) error {
	var updateErr error

	updateErr = s.updatePayment(req)
	if updateErr == nil {
		return nil
	}

	updateErr = s.updatePaymentDonation(req)
	if updateErr == nil {
		return nil
	}

	return fmt.Errorf("transaction ID %s not found in either Payment or PaymentDonation", req.TransactionID)
}
