package controller

import (
	"github.com/gin-gonic/gin"
	"rbac-opa-server-mariadb/app/service"
)

type ApiController interface {
	EvaluateRule(c *gin.Context)
}

type ApiControllerImpl struct {
	svc service.ApiService
}

func ApiControllerInit(apiService service.ApiService) *ApiControllerImpl {
	return &ApiControllerImpl{
		svc: apiService,
	}
}

func (a ApiControllerImpl) EvaluateRule(c *gin.Context) {
	a.svc.EvaluateRule(c)
}
