package config

import (
	"rbac-opa-server-mariadb/app/controller"
	"rbac-opa-server-mariadb/app/repository"
	"rbac-opa-server-mariadb/app/service"
)

type Initialization struct {
	ApiCtrl  controller.ApiController
	apiSvc   service.ApiService
	dataRepo repository.DataRepository
}

func NewInitialization(
	apiCtrl controller.ApiController,
	apiSvc service.ApiService,
	dataRepo repository.DataRepository,
) *Initialization {
	return &Initialization{
		ApiCtrl:  apiCtrl,
		apiSvc:   apiSvc,
		dataRepo: dataRepo,
	}
}
