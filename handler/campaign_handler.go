package handler

import (
	"bwu-startup/helper"
	"bwu-startup/model/response"
	"bwu-startup/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignSvc service.CampaignService
}

func NewCampaignHandler(campaignSvc service.CampaignService) *campaignHandler {
	return &campaignHandler{campaignSvc: campaignSvc}
}

func (ch *campaignHandler) GetCampaigns(ctx *gin.Context) {
	userId := ctx.Query("user_id")

	campaigns, err := ch.campaignSvc.GetCampaigns(userId)
	if err != nil {
		response := helper.JSONResponse("error to get campaigns", "error", http.StatusBadRequest, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.JSONResponse("list of campaigns", "success", http.StatusBadRequest, response.FormatCampaignsResponse(campaigns))
	ctx.JSON(http.StatusBadRequest, response)
	return
}
