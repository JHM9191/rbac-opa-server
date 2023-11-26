package pkg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rbac-opa-server-mariadb/app/constants"
	"strings"
)

func PanicHandler(c *gin.Context) {
	if err := recover(); err != nil {
		str := fmt.Sprint(err)
		strArr := strings.Split(str, ":")
		key := strArr[0]

		switch key {
		case constants.DataNotFound.GetResponseStatus():
			c.JSON(http.StatusBadRequest, BuildResponse(constants.DataNotFound, Null()))
			c.Abort()
		case constants.Unauthorized.GetResponseStatus():
			c.JSON(http.StatusUnauthorized, BuildResponse(constants.Unauthorized, Null()))
			c.Abort()
		case constants.InvalidRequest.GetResponseStatus():
			c.JSON(http.StatusBadRequest, BuildResponse(constants.InvalidRequest, Null()))
			c.Abort()
		default:
			c.JSON(http.StatusInternalServerError, BuildResponse(constants.InternalError, Null()))
			c.Abort()
		}
	}
}
