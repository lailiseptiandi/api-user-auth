package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lailiseptiandi/api-user-auth/app/services"
	"github.com/lailiseptiandi/api-user-auth/app/utils"
)

func AuthMiddleware(userService services.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var accessToken string

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			accessToken = fields[1]
		} else {
			accessToken = ""
		}

		if accessToken == "" {
			resp := utils.ResponseError(nil, "You are not logged in")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}

		sub, err := utils.ValidateToken(accessToken)
		if err != nil {
			resp := utils.ResponseError(nil, "Invalid Token")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}
		jsonData, _ := json.Marshal(sub)
		var data map[string]interface{}
		_ = json.Unmarshal([]byte(jsonData), &data)

		user, err := userService.FindUserById(uint(int(data["id"].(float64))))
		if err != nil {
			resp := utils.ResponseError(nil, err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}

		ctx.Set("currentUser", user)
		ctx.Next()
	}
}
