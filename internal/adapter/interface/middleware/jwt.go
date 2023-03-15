package middleware

import (
	"cuckoo/internal/utils"
	"cuckoo/pkg/app"
	"cuckoo/pkg/cuckoo_error"
	"github.com/gin-gonic/gin"
	"time"
)

// JWT 中间件
func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cuckoo := app.NewApplication(ctx)

		token := ctx.Request.Header.Get("Authorization")
		if token == "" {
			cuckoo.ErrorResponse(cuckoo_error.UserAuthTokenError)
			ctx.Abort()
			return
		}

		claims, err := utils.ParseToken(token)
		if err != nil {
			cuckoo.ErrorResponse(cuckoo_error.UserAuthTokenError)
			ctx.Abort()
			return
		}

		if time.Now().Unix() > claims.ExpiresAt {
			cuckoo.ErrorResponse(cuckoo_error.UserAuthTokenExpired)
			ctx.Abort()
			return
		}

		cuckoo.SetUserInfo(claims.UserInfo)

		ctx.Next()
	}
}
