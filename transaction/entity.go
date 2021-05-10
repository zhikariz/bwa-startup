package transaction

import (
	"bwa-startup/campaign"
	"bwa-startup/user"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	PaymentUrl string
	User       user.User
	Campaign   campaign.Campaign
}
