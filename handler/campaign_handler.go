package handler

import (
	"bwu-startup/helper"
	"bwu-startup/model"
	"bwu-startup/model/request"
	"bwu-startup/model/response"
	"bwu-startup/service"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

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

func (ch *campaignHandler) UpdateCampaign(ctx *gin.Context) {
	requestId := request.GetCampaignDetailRequest{}
	err := ctx.ShouldBindUri(&requestId)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.JSONResponse("bad request", "error", http.StatusBadRequest, gin.H{"errors": errors})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	requestForm := request.CreteaCampaingRequest{}
	err = ctx.ShouldBindJSON(&requestForm)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.JSONResponse("bad request", "error", http.StatusBadRequest, gin.H{"errors": errors})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// currentUser from middleware authentication
	getCurrentUser := ctx.MustGet("currentUser").(*model.User)
	requestForm.User = *getCurrentUser

	updatedCampaign, err := ch.campaignSvc.UpdateCampaign(requestId, requestForm)
	if err != nil {
		response := helper.JSONResponse("failed update campaign", "error", http.StatusInternalServerError, gin.H{"errors": err.Error()})
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.JSONResponse("success update campaign", "success", http.StatusOK, response.FormatCampaignResponse(*updatedCampaign))
	ctx.JSON(http.StatusOK, response)
}

func (ch *campaignHandler) SaveCampaignImage(ctx *gin.Context) {
	request := request.CreateCampaignImageRequest{}
	err := ctx.ShouldBind(&request)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.JSONResponse("bad request", "error", http.StatusBadRequest, gin.H{"errors": errors})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		response := helper.JSONResponse("bad request", "error", http.StatusBadRequest, gin.H{"errors": err.Error()})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// currentUser from middleware authentication
	getCurrentUser := ctx.MustGet("currentUser").(*model.User)
	idUser := getCurrentUser.ID
	chunkedId := strings.Split(idUser, "-")
	pathImage := fmt.Sprintf("public/images/campaign/%s-%s.%s", chunkedId[0], file.Filename, filepath.Ext(file.Filename))
	if err != nil {
		response := helper.JSONResponse("failed to upload campaign image", "error", http.StatusInternalServerError, gin.H{"is_uploaded": false})
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	request.UserId = idUser
	_, err = ch.campaignSvc.SaveCampaignImage(request, pathImage)
	if err != nil {
		response := helper.JSONResponse("failed to upload campaign image", "error", http.StatusInternalServerError, gin.H{"is_uploaded": false})
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.JSONResponse("success upload campaign image", "success ", http.StatusOK, gin.H{"is_uploaded": true})
	ctx.JSON(http.StatusOK, response)
	return
}
