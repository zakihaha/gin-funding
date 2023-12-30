package transaction

import "github.com/zakihaha/gin-funding/user"

type GetTransactionsByCampaignIDInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}

type CreateTransactionInput struct {
	CampaignID int `json:"campaign_id" binding:"required"`
	Amount     int `json:"amount" binding:"required"`
	User       user.User
}
