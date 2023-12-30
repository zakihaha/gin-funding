package transaction

import (
	"time"

	"github.com/zakihaha/gin-funding/campaign"
	"github.com/zakihaha/gin-funding/user"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	PaymentURL string
	CreatedAt  time.Time
	UpdatedAt  time.Time

	User     user.User
	Campaign campaign.Campaign
}
