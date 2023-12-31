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

func (h *transactionHandler) GetTransactionsByUserID(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	transactions, err := h.transactionService.GetTransactionsByUserID(int(userID))
	if err != nil {
		data := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed to get transactions by user id", http.StatusInternalServerError, "error", data)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formatter := transaction.FormatUserTransactions(transactions)

	response := helper.APIResponse("Successfully to get transactions by user id", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var input transaction.CreateTransactionInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		data := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to get transactions by user id", http.StatusInternalServerError, "error", data)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	createdTransaction, err := h.transactionService.CreateTransaction(input)
	if err != nil {
		data := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed to create transaction", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := transaction.FormatTransaction(createdTransaction)

	response := helper.APIResponse("Successfully to create transaction", http.StatusCreated, "success", formatter)

	c.JSON(http.StatusCreated, response)
}

func (h *transactionHandler) GetNotification(c *gin.Context) {
	var input transaction.TransactionNotificationInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		data := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to get notification", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = h.transactionService.Webhook(input)
	if err != nil {
		data := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed to process payment", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, input)
}
