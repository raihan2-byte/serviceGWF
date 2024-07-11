package service

import (
	"fmt"
	"payment-gwf/entity"
	"payment-gwf/input"
	"payment-gwf/repository"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type ServicePaymentDonation interface {
	GetAllPaymentDonation() ([]*entity.PaymentDonation, error)
	GetPaymentDonationByID(ID string) (*entity.PaymentDonation, error)
	DeletePaymentDonation(ID string) (*entity.PaymentDonation, error)
	GetAllDonationByUserID(ID int) ([]*entity.PaymentDonation, error)
	DoPaymentDonation(req input.SubmitPaymentRequest, makeDonationID string, userID int) (*entity.DoPayment, error)
	HandleNotificationPaymentDonation(req *entity.MidtransNotificationRequest) error
	FindStatus(orderID string) (*entity.PaymentDonation, error)
}

type servicePaymentDonation struct {
	repositoryPaymentDetails  repository.RepositoryPaymentDetails
	repositoryPaymentDonation repository.RepositoryPaymentDonation
	repositoryUser            repository.RepositoryUser
	repositoryOrder           repository.RepositoryOrder
	repositoryMakeDonation    repository.RepositoryMakeDonation
	midtransGateway           *midtransGateway
}

func NewServicePaymentDonation(repositoryPaymentDetails repository.RepositoryPaymentDetails, repositoryPaymentDonation repository.RepositoryPaymentDonation, repositoryPayment repository.RepositoryPayment, repositoryUser repository.RepositoryUser, repositoryOrder repository.RepositoryOrder, repositoryMakeDonation repository.RepositoryMakeDonation, midtransGateway *midtransGateway) *servicePaymentDonation {
	return &servicePaymentDonation{
		repositoryPaymentDetails:  repositoryPaymentDetails,
		repositoryPaymentDonation: repositoryPaymentDonation,
		repositoryUser:            repositoryUser,
		repositoryOrder:           repositoryOrder,
		repositoryMakeDonation:    repositoryMakeDonation,
		midtransGateway:           midtransGateway,
	}
}

func (s *servicePaymentDonation) FindStatus(orderID string) (*entity.PaymentDonation, error) {
	get, err := s.repositoryPaymentDonation.FindByOrderId(orderID)
	if err != nil {
		return get, err
	}

	return get, nil
}

func (s *servicePaymentDonation) DoPaymentDonation(req input.SubmitPaymentRequest, makeDonationID string, userID int) (*entity.DoPayment, error) {
	getDonationID, err := s.repositoryMakeDonation.FindById(makeDonationID)
	if err != nil {
		return nil, fmt.Errorf("error fetching donation: %w", err)
	}

	if getDonationID.UserID != userID {
		return nil, fmt.Errorf("unauthorized user")
	}

	chosenBank, ok := listOfBank[req.BankTransfer]
	if !ok {
		return nil, fmt.Errorf("unsupported bank")
	}

	resp, err := s.midtransGateway.client.ChargeTransaction(&coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  getDonationID.ID,
			GrossAmt: int64(getDonationID.Amount),
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: chosenBank,
		},
	})

	// if err != nil {
	// 	return nil, fmt.Errorf("error during charge transaction: %w", err)
	// }

	pay := &entity.PaymentDonation{
		StatusPayment:  "pending",
		MakeDonationID: getDonationID.ID,
		UserID:         getDonationID.UserID,
		TransactionID:  resp.TransactionID,
	}

	// mappedResp, err := mapChargeToResponse(resp)
	// if err != nil {
	// 	return nil, fmt.Errorf("error mapping charge response: %w", err)
	// }

	// _, err = s.repositoryPaymentDonation.SaveMidtrans(mappedResp)
	// if err != nil {
	// 	return nil, fmt.Errorf("error saving payment to Midtrans: %w", err)
	// }

	_, err = s.repositoryPaymentDonation.Save(pay)
	if err != nil {
		return nil, fmt.Errorf("error saving payment donation: %w", err)
	}

	return mapChargeToResponseDonation(resp)
}
func mapChargeToResponseDonation(resp *coreapi.ChargeResponse) (*entity.DoPayment, error) {
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
		TransactionID:   resp.TransactionID,
		OrderID:         resp.OrderID,
		GrossAmount:     resp.GrossAmount,
		VaNumbers:       vaNumbers,
		TransactionTime: resp.TransactionTime,
		MerchantID:      "G317569034",
	}, nil
}

func (s *servicePaymentDonation) HandleNotificationPaymentDonation(req *entity.MidtransNotificationRequest) error {
	findPayment, err := s.repositoryPaymentDonation.FindByTransactionID(req.TransactionID)
	if err != nil {
		return err
	}

	switch req.TransactionStatus {
	case "settlement":
		// Ubah status pembayaran menjadi 'settled'
		findPayment.StatusPayment = "settled"
	default:
		// Ubah status pembayaran sesuai dengan status yang diterima
		findPayment.StatusPayment = req.TransactionStatus
	}

	_, err = s.repositoryPaymentDonation.Update(findPayment)
	if err != nil {
		return err
	}
	return nil
}

func (s *servicePaymentDonation) GetAllPaymentDonation() ([]*entity.PaymentDonation, error) {
	address, err := s.repositoryPaymentDonation.FindAll()
	if err != nil {
		return address, err
	}
	return address, nil
}

func (s *servicePaymentDonation) GetPaymentDonationByID(ID string) (*entity.PaymentDonation, error) {
	address, err := s.repositoryPaymentDonation.FindById(ID)
	if err != nil {
		return address, err
	}
	return address, nil
}

func (s *servicePaymentDonation) GetAllDonationByUserID(ID int) ([]*entity.PaymentDonation, error) {
	getUserID, err := s.repositoryUser.FindById(ID)
	if err != nil {
		return nil, err
	}
	get, err := s.repositoryPaymentDonation.FindAllByUserID(getUserID.ID)
	if err != nil {
		return get, err
	}
	return get, nil
}

func (s *servicePaymentDonation) DeletePaymentDonation(ID string) (*entity.PaymentDonation, error) {
	address, err := s.repositoryPaymentDonation.FindById(ID)
	if err != nil {
		return address, err
	}
	addressDel, err := s.repositoryPaymentDonation.Delete(address)
	if err != nil {
		return addressDel, err
	}
	return addressDel, nil
}
