package service

import (
	"bwu-startup/model"
	"bwu-startup/model/request"
	"bwu-startup/repository"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

type CampaignService interface {
	GetCampaigns(userID string) ([]model.Campaign, error)
	GetCampaignDetailById(request request.GetCampaignDetailRequest) (*model.Campaign, error)
	CreateCampaign(request request.CreteaCampaingRequest) (*model.Campaign, error)
	UpdateCampaign(ID request.GetCampaignDetailRequest, request request.CreteaCampaingRequest) (*model.Campaign, error)
	SaveCampaignImage(request request.CreateCampaignImageRequest, pathFile string) (*model.CampaignImage, error)
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

func (cs *campaignService) CreateCampaign(request request.CreteaCampaingRequest) (*model.Campaign, error) {
	campaign := model.Campaign{}
	campaign.Name = request.Name
	campaign.ShortDescription = request.ShortDescription
	campaign.Description = request.Description
	campaign.Perks = request.Perks
	campaign.GoalAmount = request.GoalAmount
	campaign.UserId = request.User.ID
	campaign.ID = uuid.New().String()
	campaign.CreatedAt = time.Now()

	// slug
	idUserChunked := strings.Split(request.User.ID, "-")
	slugCandidate := fmt.Sprintf("%s %s", request.Name, idUserChunked[0])
	campaign.Slug = slug.Make(slugCandidate)

	newCampaign, err := cs.campaignRepo.Create(campaign)
	if err != nil {
		return nil, err
	}

	return newCampaign, nil
}

func (cs *campaignService) UpdateCampaign(ID request.GetCampaignDetailRequest, request request.CreteaCampaingRequest) (*model.Campaign, error) {
	campaign, err := cs.campaignRepo.FindById(ID.ID)
	if err != nil {
		return nil, err
	}

	if campaign.UserId != request.User.ID {
		return nil, errors.New("you not own this campaign")
	}

	campaign.Name = request.Name
	campaign.ShortDescription = request.ShortDescription
	campaign.Description = request.Description
	campaign.Perks = request.Perks
	campaign.GoalAmount = request.GoalAmount
	campaign.UpdatedAt = time.Now()

	updatedCampaign, err := cs.campaignRepo.Update(*campaign)
	if err != nil {
		return nil, err
	}

	return updatedCampaign, nil
}

func (cs *campaignService) SaveCampaignImage(request request.CreateCampaignImageRequest, pathFile string) (*model.CampaignImage, error) {
	campaign, err := cs.campaignRepo.FindById(request.CampaignId)
	if err != nil {
		return nil, err
	}

	if campaign.UserId != request.UserId {
		return nil, errors.New("you not own this campaign")
	}

	if request.IsPrimary {
		_, err := cs.campaignRepo.ChangesMarkAllImageToNonPrimary(request.CampaignId)
		if err != nil {
			return nil, err
		}
	}

	campaignImage := model.CampaignImage{
		ID:         uuid.New().String(),
		CampaignID: request.CampaignId,
		FileName:   pathFile,
		IsPrimary:  request.IsPrimary,
		CreatedAt:  time.Now(),
	}

	newCampaignImage, err := cs.campaignRepo.CreateCampaignImage(campaignImage)
	if err != nil {
		return nil, err
	}

	return newCampaignImage, nil
}
