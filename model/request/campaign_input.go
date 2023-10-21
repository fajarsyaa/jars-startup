package request

type GetCampaignDetailRequest struct {
	ID string `uri:"id" binding:"required"`
}
