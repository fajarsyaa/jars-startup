package repository

import (
	"bwu-startup/model"

	"gorm.io/gorm"
)

type CampaignRepository interface {
	FindAll() ([]model.Campaign, error)
	FindByUserID(userId string) ([]model.Campaign, error)
	FindById(Id string) (*model.Campaign, error)
	Create(campaign model.Campaign) (*model.Campaign, error)
	Update(campaign model.Campaign) (*model.Campaign, error)
}

type campaignRepository struct {
	db *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) *campaignRepository {
	return &campaignRepository{db: db}
}

func (cr *campaignRepository) FindAll() ([]model.Campaign, error) {
	var campaigns []model.Campaign

	// preload("dari field di struct","kondisi dari field di db")
	err := cr.db.Preload("CampaignImages", "campaign_images.is_primary = true").Find(&campaigns).Error
	if err != nil {
		return nil, err
	}

	return campaigns, err
}

func (cr *campaignRepository) FindByUserID(userId string) ([]model.Campaign, error) {
	var campaigns []model.Campaign

	// preload("dari field di struct","kondisi dari field di db")
	err := cr.db.Where("user_id = ?", userId).Preload("CampaignImages", "campaign_images.is_primary = true").Find(&campaigns).Error
	if err != nil {
		return nil, err
	}

	return campaigns, nil
}

func (cr *campaignRepository) FindById(Id string) (*model.Campaign, error) {
	var campaign model.Campaign
	// preload("dari field di struct")
	err := cr.db.Where("id = ?", Id).Preload("User").Preload("CampaignImages").Find(&campaign).Error
	if err != nil {
		return nil, err
	}

	return &campaign, nil
}

func (cr *campaignRepository) Create(campaign model.Campaign) (*model.Campaign, error) {
	err := cr.db.Create(&campaign).Error
	if err != nil {
		return nil, err
	}

	return &campaign, nil
}

func (cr *campaignRepository) Update(campaign model.Campaign) (*model.Campaign, error) {
	err := cr.db.Save(&campaign).Error
	if err != nil {
		return nil, err
	}

	return &campaign, nil
}
