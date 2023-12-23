package handler

import (
	"net/http"

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

func (h *campaignHandler) GetAll(c *gin.Context) {
	campaigns, err := h.campaignService.GetAll()
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed to get campaigns", http.StatusInternalServerError, "error", errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.APIResponse("Successfully to get campaign", http.StatusOK, "success", campaigns)
	c.JSON(http.StatusOK, response)
}
