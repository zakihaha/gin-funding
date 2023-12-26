package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zakihaha/gin-funding/helper"
	"github.com/zakihaha/gin-funding/transaction"
	"github.com/zakihaha/gin-funding/user"
)

type transactionHandler struct {
	transactionService transaction.Service
}

func NewTransactionHandler(
	transactionService transaction.Service,
) *transactionHandler {
	return &transactionHandler{
		transactionService,
	}
}

func (h *transactionHandler) GetTransactionsByCampaignID(c *gin.Context) {
	var input transaction.GetTransactionsByCampaignIDInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		data := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed to get transactions by campaign ID", http.StatusInternalServerError, "error", data)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err := h.transactionService.GetTransactionsByCampaignID(input)
	if err != nil {
		data := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed to get transactions by campaign ID", http.StatusInternalServerError, "error", data)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formatter := transaction.FormatCampaignTransactions(transactions)

	response := helper.APIResponse("Successfully to get transactions by campaign ID", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
