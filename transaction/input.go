package transaction

import "github.com/zakihaha/gin-funding/user"

type GetTransactionsByCampaignIDInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
