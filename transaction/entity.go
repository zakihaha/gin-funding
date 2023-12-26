package transaction

import (
	"time"

	"github.com/zakihaha/gin-funding/user"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	CreatedAt  time.Time
	UpdatedAt  time.Time

	// Campaign Campaign
	User user.User
}
