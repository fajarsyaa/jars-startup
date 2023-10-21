package response

import "bwu-startup/model"

type campaignFormatter struct {
	ID               string `json:"id"`
	UserId           string `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	Slug             string `json:"slug"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
}

func formatCampaignResponse(campaign model.Campaign) campaignFormatter {
	responseFormatter := campaignFormatter{}
	responseFormatter.ID = campaign.ID
	responseFormatter.UserId = campaign.UserId
	responseFormatter.Name = campaign.Name
	responseFormatter.ShortDescription = campaign.ShortDescription
	responseFormatter.GoalAmount = campaign.GoalAmount
	responseFormatter.CurrentAmount = campaign.CurrentAmount
	responseFormatter.Slug = campaign.Slug
	responseFormatter.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		responseFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return responseFormatter
}

func FormatCampaignsResponse(campaigns []model.Campaign) []campaignFormatter {
	formattedCampaigns := []campaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := formatCampaignResponse(campaign)
		formattedCampaigns = append(formattedCampaigns, campaignFormatter)
	}

	return formattedCampaigns
}
