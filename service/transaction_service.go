package service

import (
	"bwu-startup/model"
	"bwu-startup/model/request"
	"bwu-startup/repository"
	"errors"
)

type TransactionService interface {
	GetTransactionByCampaignID(request request.GetCampaignTransactionRequest) ([]model.Transaction, error)
}

type transactionService struct {
	trxRepo  repository.TransactionRepository
	campRepo repository.CampaignRepository
}

func NewTransactionService(tr repository.TransactionRepository, camp repository.CampaignRepository) *transactionService {
	return &transactionService{trxRepo: tr, campRepo: camp}
}

func (ts *transactionService) GetTransactionByCampaignID(request request.GetCampaignTransactionRequest) ([]model.Transaction, error) {

	campaign, err := ts.campRepo.FindById(request.ID)
	if err != nil {
		return []model.Transaction{}, err
	}

	if campaign.UserId != request.User.ID {
		return []model.Transaction{}, errors.New("you not own this campaign")
	}

	transactions, err := ts.trxRepo.GetByCampaignID(request.ID)
	if err != nil {
		return []model.Transaction{}, err
	}

	return transactions, nil
}
