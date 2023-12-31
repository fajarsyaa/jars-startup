package request

import "bwu-startup/model"

type GetCampaignDetailRequest struct {
	ID string `uri:"id" binding:"required"`
}

type CreteaCampaingRequest struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required"`
	User             model.User
}

type CreateCampaignImageRequest struct {
	CampaignId string `form:"campaign_id" binding:"required"`
	IsPrimary  bool   `form:"is_primary"`
	UserId     string
}
