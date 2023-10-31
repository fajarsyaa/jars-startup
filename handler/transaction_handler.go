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

type transactionHandler struct {
	trxSvc service.TransactionService
}

func NewTransactionHandler(svc service.TransactionService) *transactionHandler {
	return &transactionHandler{trxSvc: svc}
}

func (th *transactionHandler) GetCampaignTransactions(ctx *gin.Context) {
	request := request.GetCampaignTransactionRequest{}
	err := ctx.ShouldBindUri(&request)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.JSONResponse("bad request", "error", http.StatusBadRequest, gin.H{"errors": errors})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// currentUser from middleware authentication
	getCurrentUser := ctx.MustGet("currentUser").(*model.User)
	request.User = *getCurrentUser

	transactions, err := th.trxSvc.GetTransactionByCampaignID(request)
	if err != nil {
		response := helper.JSONResponse("failed to get campaign's transaction", "error", http.StatusInternalServerError, nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.JSONResponse("success  to get campaign's transaction", "success", http.StatusOK, response.FormatTransactionResponses(transactions))
	ctx.JSON(http.StatusOK, response)
}
