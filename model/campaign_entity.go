package model

import "time"

type Campaign struct {
	ID               string
	UserId           string
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	Slug             string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImages   []CampaignImage
}

type CampaignImage struct {
	ID         string
	CampaignID string
	FileName   string
	IsPrimary  bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
