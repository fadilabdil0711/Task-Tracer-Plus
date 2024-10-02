package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				if ctx.GetHeader("Content-Type") == "application/json" {
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
					return
				}
				ctx.Redirect(http.StatusSeeOther, "/")
				return
			}
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		token, err := jwt.ParseWithClaims(cookie, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})
		if err != nil || !token.Valid {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		claims, ok := token.Claims.(*model.Claims)
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("email", claims.Email)

		ctx.Next()
	})
}
