package middleware

import (
	"bwu-startup/helper"
	"bwu-startup/helper/jwt_token"

	"bwu-startup/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authorization(auth jwt_token.JwtToken, us service.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		tokenUser := ""

		if !strings.Contains(header, "Bearer") {
			response := helper.JSONResponse("token not valid", "unauthorization", http.StatusUnauthorized, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		arrToken := strings.Split(header, " ")
		if len(arrToken) == 2 {
			tokenUser = arrToken[1]
		}

		token, err := auth.ValidateToken(tokenUser)
		if err != nil {
			response := helper.JSONResponse("token not valid", "unauthorization", http.StatusUnauthorized, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.JSONResponse("token not valid", "unauthorization", http.StatusUnauthorized, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := claim["user_id"].(string)

		user, err := us.GetUserById(userId)
		if err != nil {
			response := helper.JSONResponse("token not valid", "unauthorization", http.StatusUnauthorized, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		ctx.Set("currentUser", user)
	}
}
