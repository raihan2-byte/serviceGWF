package service

import (
	"fmt"
	"payment-gwf/entity"
	"payment-gwf/repository"

	"github.com/midtrans/midtrans-go/coreapi"
)

type ServicePaymentDetails interface {
	// GetAllPaymentDetails() ([]*entity.PaymentDetails, error)
	// CreatePaymentDetails(paymentID int, paymentDonationID int) (*entity.PaymentDetails, error)
	// GetPaymentDetailsByID(ID int) (*entity.PaymentDetails, error)
	// DeletePaymentDetails(ID int) (*entity.PaymentDetails, error)
	// GetAllPaymentDetailsByUserID(ID int) ([]*entity.PaymentDetails, error)
	// DoPaymentDetails(req input.SubmitPaymentRequest, makeDonationID int) (*entity.DoPayment, error)
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

// func (s *servicePaymentDetails) CreatePaymentDetails(paymentID int, paymentDonationID int) (*entity.PaymentDetails, error) {
// 	findPayment, err := s.repositoryPayment.FindById(paymentID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	findPaymntDonation, err := s.repositoryPaymentDonation.FindById(paymentDonationID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	getPayment := &entity.PaymentDetails{}
// 	getPayment.PaymentID = &findPayment.ID
// 	getPayment.PaymentDonationID = &findPaymntDonation.ID

// 	if err != nil {
// 		return nil, err
// 	}
// 	return getPayment, nil

// }

// func (s *servicePaymentDonation) DoPaymentDetails(req input.SubmitPaymentRequest, makeDonationID int) (*entity.DoPayment, error) {
// 	// getOrder, err := s.repositoryOrder.FindById(orderID)
// 	// if err != nil {
// 	// 	return nil, fmt.Errorf("error fetching order: %w", err)
// 	// }

// 	getDonationID, err := s.repositoryMakeDonation.FindById(makeDonationID)
// 	if err != nil {
// 		return nil, fmt.Errorf("error fetching order: %w", err)
// 	}

// 	valGetMakeDonationID := strconv.Itoa(getDonationID.ID)
// 	chosenBank, ok := listOfBank[req.BankTransfer]
// 	if !ok {
// 		return nil, fmt.Errorf("unsupported bank")
// 	}

// 	log.Println("Order ID:", valGetMakeDonationID)
// 	log.Println("Total Price:", getDonationID.Amount)
// 	log.Println("Chosen Bank:", chosenBank)

// 	resp, err := s.midtransGateway.client.ChargeTransaction(&coreapi.ChargeReq{
// 		PaymentType: coreapi.PaymentTypeBankTransfer,
// 		TransactionDetails: midtrans.TransactionDetails{
// 			OrderID:  valGetMakeDonationID,
// 			GrossAmt: int64(getDonationID.Amount),
// 		},
// 		BankTransfer: &coreapi.BankTransferDetails{
// 			Bank: chosenBank,
// 		},
// 	})

// 	pay := &entity.PaymentDetails{}

// 	pay.StatusPayment = "pending"
// 	pay.MakeDonationID = getDonationID.ID
// 	pay.TransactionID = resp.TransactionID
// 	pay.UserID = getDonationID.UserID
// 	_, err = s.repositoryPaymentDonation.Update(pay)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// if err != nil {
// 	// 	return nil, fmt.Errorf("error during charge transaction: %w", err)
// 	// }

// 	return mapChargeToResponseDonation(resp)
// }

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

// func (s *servicePaymentDonation) GetAllPaymentDetails() ([]*entity.PaymentDetails, error) {
// 	address, err := s.repositoryPaymentDonation.FindAll()
// 	if err != nil {
// 		return address, err
// 	}
// 	return address, nil
// }

// func (s *servicePaymentDonation) GetPaymentDetailsByID(ID int) (*entity.PaymentDetails, error) {
// 	address, err := s.repositoryPaymentDonation.FindById(ID)
// 	if err != nil {
// 		return address, err
// 	}
// 	return address, nil
// }

// func (s *servicePaymentDonation) GetAllPaymentDetailsByUserID(ID int) ([]*entity.PaymentDetails, error) {
// 	getUserID, err := s.repositoryUser.FindById(ID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	get, err := s.repositoryPaymentDonation.FindAllByUserID(getUserID.ID)
// 	if err != nil {
// 		return get, err
// 	}
// 	return get, nil
// }

// func (s *servicePaymentDonation) DeletePaymentDetails(ID int) (*entity.PaymentDetails, error) {
// 	address, err := s.repositoryPaymentDonation.FindById(ID)
// 	if err != nil {
// 		return address, err
// 	}
// 	addressDel, err := s.repositoryPaymentDonation.Delete(address)
// 	if err != nil {
// 		return addressDel, err
// 	}
// 	return addressDel, nil
// }
