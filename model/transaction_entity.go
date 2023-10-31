package model

import "time"

type Transaction struct {
	ID         string
	CampaignId string
	UserID     string
	Status     string
	Code       string
	Amount     int
	CreatedAt  time.Time
	UpdattedAt time.Time
	User       User
}
