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

func (uh *userHandler) Login(ctx *gin.Context) {
	request := request.LoginRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.JSONResponse("bad request", "error", http.StatusBadRequest, gin.H{"errors": errors})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := uh.usrSvc.Login(request)
	if err != nil {
		response := helper.JSONResponse("login failed", "error", http.StatusInternalServerError, gin.H{"errors": err.Error()})
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	formatData := response.FormatUserResponse(*user, "initokenabalabal")
	response := helper.JSONResponse("login success", "success", http.StatusOK, formatData)
	ctx.JSON(http.StatusOK, response)
}

func (uh *userHandler) CheckAvailableEmail(ctx *gin.Context) {
	request := request.AvailableEmailRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.JSONResponse("bad request", "error", http.StatusBadRequest, gin.H{"errors": errors})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	isAvailableEmail, err := uh.usrSvc.CheckAvailableEmail(request)
	if err != nil {
		response := helper.JSONResponse("check email failed", "error", http.StatusInternalServerError, gin.H{"errors": err.Error()})
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	if !isAvailableEmail {
		response := helper.JSONResponse("email is already in use", "error", http.StatusBadRequest, gin.H{"is_available": isAvailableEmail})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.JSONResponse("email is available", "success", http.StatusOK, gin.H{"is_available": isAvailableEmail})
	ctx.JSON(http.StatusOK, response)
}
