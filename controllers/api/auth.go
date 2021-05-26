package api

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/KervinChang/Blog-API/models"
	"github.com/KervinChang/Blog-API/pkg/exception"
	"github.com/KervinChang/Blog-API/pkg/logging"
	"github.com/KervinChang/Blog-API/pkg/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := exception.INVALID_PARAMS
	if ok {
		IsExist := models.CheckAuth(username, password)
		if IsExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = exception.INVALID_AUTH_TOKEN
			} else {
				data["token"] = token
				code = exception.SUCCESS
			}
		} else {
			code = exception.AUTH_FAIL
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : exception.GetMsg(code),
		"data" : data,
	})
}