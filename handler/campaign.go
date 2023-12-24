package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zakihaha/gin-funding/campaign"
	"github.com/zakihaha/gin-funding/helper"
)

type campaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(
	campaignService campaign.Service,
) *campaignHandler {
	return &campaignHandler{
		campaignService,
	}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.campaignService.GetCampaigns(userID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed to get campaigns", http.StatusInternalServerError, "error", errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formatter := campaign.FormatCampaigns(campaigns)

	response := helper.APIResponse("Successfully to get campaign", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed to get detail campaign", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.campaignService.GetCampaignByID(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed to get detail campaign", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := campaign.FormatCampaignDetail(campaignDetail)

	response := helper.APIResponse("Successfully to get detail campaign", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
