package response

import (
	"bwu-startup/model"
	"time"
)

type TransactionFormatter struct {
	ID        string    `json="id"`
	Name      string    `json="name"`
	Amount    int       `json="amount"`
	CreatedAt time.Time `json="created_at"`
}

func FormatTransactionResponse(trx model.Transaction) TransactionFormatter {
	return TransactionFormatter{
		ID:        trx.ID,
		Name:      trx.User.Name,
		Amount:    trx.Amount,
		CreatedAt: trx.CreatedAt,
	}
}

func FormatTransactionResponses(trxs []model.Transaction) []TransactionFormatter {
	if len(trxs) == 0 {
		return []TransactionFormatter{}
	}

	var trxFormatters []TransactionFormatter

	for _, trx := range trxs {
		formatter := FormatTransactionResponse(trx)
		trxFormatters = append(trxFormatters, formatter)
	}

	return trxFormatters
}
