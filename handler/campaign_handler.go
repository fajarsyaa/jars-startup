package handler

import (
	"bwu-startup/helper"
	"bwu-startup/model"
	"bwu-startup/model/request"
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

	response := helper.JSONResponse("list of campaigns", "success", http.StatusOK, response.FormatCampaignsResponse(campaigns))
	ctx.JSON(http.StatusOK, response)
	return
}

func (ch *campaignHandler) GetCampaignDetail(ctx *gin.Context) {
	request := request.GetCampaignDetailRequest{}

	err := ctx.ShouldBindUri(&request)
	if err != nil {
		response := helper.JSONResponse("error to get campaign detail", "error", http.StatusBadRequest, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	campaign, err := ch.campaignSvc.GetCampaignDetailById(request)
	if err != nil {
		response := helper.JSONResponse("error to get campaign detail", "error", http.StatusBadRequest, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.JSONResponse("campaign detail", "success", http.StatusOK, response.FormatCampaignsDetailResponse(*campaign))
	ctx.JSON(http.StatusOK, response)
	return
}

func (ch *campaignHandler) CreateCampaign(ctx *gin.Context) {
	request := request.CreteaCampaingRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.JSONResponse("bad request", "error", http.StatusBadRequest, gin.H{"errors": errors})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// currentUser from middleware authentication
	getCurrentUser := ctx.MustGet("currentUser").(*model.User)
	request.User = *getCurrentUser

	newCampaign, err := ch.campaignSvc.CreateCampaign(request)
	if err != nil {
		response := helper.JSONResponse("failed create campaign", "error", http.StatusInternalServerError, gin.H{"errors": err.Error()})
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.JSONResponse("success create campaign", "success", http.StatusOK, response.FormatCampaignResponse(*newCampaign))
	ctx.JSON(http.StatusOK, response)
}
