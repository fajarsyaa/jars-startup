package handler

import (
	"bwu-startup/helper"
	"bwu-startup/model/request"
	"bwu-startup/model/response"
	"bwu-startup/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	usrSvc service.UserService
}

func NewUserHandler(usrSvc service.UserService) *userHandler {
	return &userHandler{usrSvc: usrSvc}
}

func (uh *userHandler) RegisterUser(ctx *gin.Context) {
	var request request.RegisterUserRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.JSONResponse("bad request", "error", http.StatusBadRequest, gin.H{"errors": errors})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	data, err := uh.usrSvc.UserRegister(request)
	if err != nil {
		response := helper.JSONResponse("register failed", "error", http.StatusInternalServerError, nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	formatData := response.FormatUserResponse(*data, "initokenabalabal")
	response := helper.JSONResponse("register success", "success", http.StatusOK, formatData)
	ctx.JSON(http.StatusOK, response)
}
