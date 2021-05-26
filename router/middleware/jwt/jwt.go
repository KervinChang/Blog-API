package jwt

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/KervinChang/Blog-API/pkg/exception"
	"github.com/KervinChang/Blog-API/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = exception.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = exception.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = exception.INVALID_AUTH_TOKEN
			} else if time.Now().Unix() > claims.ExpiresAt{
				code = exception.AUTH_FAIL
			}
		}

		if code != exception.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg": exception.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}