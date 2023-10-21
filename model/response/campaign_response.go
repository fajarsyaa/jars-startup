package response

import (
	"bwu-startup/model"
	"strings"
)

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

type campaignDetailFormatter struct {
	ID               string                   `json:"id"`
	UserId           string                   `json:"user_id"`
	Name             string                   `json:"name"`
	ShortDescription string                   `json:"short_description"`
	Description      string                   `json:"description"`
	ImageURL         string                   `json:"image_url"`
	Slug             string                   `json:"slug"`
	GoalAmount       int                      `json:"goal_amount"`
	CurrentAmount    int                      `json:"current_amount"`
	Perks            []string                 `json:"perks"`
	User             campaignUserFormatter    `json:"user"`
	Images           []campaignImageFormatter `json:"images"`
}

type campaignUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type campaignImageFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignsDetailResponse(campaign model.Campaign) campaignDetailFormatter {
	responseFormatter := campaignDetailFormatter{}
	responseFormatter.ID = campaign.ID
	responseFormatter.UserId = campaign.UserId
	responseFormatter.Name = campaign.Name
	responseFormatter.ShortDescription = campaign.ShortDescription
	responseFormatter.Description = campaign.Description
	responseFormatter.Slug = campaign.Slug
	responseFormatter.GoalAmount = campaign.GoalAmount
	responseFormatter.CurrentAmount = campaign.CurrentAmount
	responseFormatter.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		responseFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	perks := []string{}
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	responseFormatter.Perks = perks

	userCampaign := campaignUserFormatter{}
	userCampaign.Name = campaign.User.Name
	userCampaign.ImageURL = campaign.User.AvatarFileName
	responseFormatter.User = userCampaign

	imagesCampaign := []campaignImageFormatter{}
	for _, image := range campaign.CampaignImages {
		img := campaignImageFormatter{
			ImageURL:  image.FileName,
			IsPrimary: image.IsPrimary,
		}

		imagesCampaign = append(imagesCampaign, img)
	}

	responseFormatter.Images = imagesCampaign

	return responseFormatter
}
