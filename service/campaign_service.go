package service

import (
	"bwu-startup/model"
	"bwu-startup/model/request"
	"bwu-startup/repository"
)

type CampaignService interface {
	GetCampaigns(userID string) ([]model.Campaign, error)
	GetCampaignDetailById(request request.GetCampaignDetailRequest) (*model.Campaign, error)
}

type campaignService struct {
	campaignRepo repository.CampaignRepository
}

func NewCampaignService(cmpgn repository.CampaignRepository) *campaignService {
	return &campaignService{campaignRepo: cmpgn}
}

func (cs *campaignService) GetCampaigns(userID string) ([]model.Campaign, error) {
	if userID == "" {
		campaigns, err := cs.campaignRepo.FindAll()
		if err != nil {
			return nil, err
		}

		return campaigns, nil
	}

	campaigns, err := cs.campaignRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	return campaigns, nil
}

func (cs *campaignService) GetCampaignDetailById(request request.GetCampaignDetailRequest) (*model.Campaign, error) {
	campaign, err := cs.campaignRepo.FindById(request.ID)
	if err != nil {
		return nil, err
	}

	return campaign, nil
}
