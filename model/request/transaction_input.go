package request

import "bwu-startup/model"

type GetCampaignTransactionRequest struct {
	ID   string `uri:"id" binding:"required"`
	User model.User
}
