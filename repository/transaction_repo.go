package repository

import (
	"bwu-startup/model"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetByCampaignID(campaignID string) ([]model.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db: db}
}

func (tr *transactionRepository) GetByCampaignID(campaignID string) ([]model.Transaction, error) {
	transactions := []model.Transaction{}

	err := tr.db.Preload("User").Where("campaign_id = ?", campaignID).Order("created_at desc").Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
