// go:build wireinject
//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"
	"rbac-opa-server-mariadb/app/controller"
	"rbac-opa-server-mariadb/app/repository"
	"rbac-opa-server-mariadb/app/service"
)

var db = wire.NewSet(ConnectToDB)

var apiCtrlSet = wire.NewSet(
	controller.ApiControllerInit,
	wire.Bind(
		new(controller.ApiController),
		new(*controller.ApiControllerImpl),
	))

var apiSvcSet = wire.NewSet(
	service.ApiServiceInit,
	wire.Bind(
		new(service.ApiService),
		new(*service.ApiServiceImpl),
	))

var dataRepoSet = wire.NewSet(
	repository.DataRepositoryInit,
	wire.Bind(
		new(repository.DataRepository),
		new(*repository.DataRepositoryImpl),
	))

func Init() *Initialization {
	wire.Build(
		NewInitialization,
		db,
		apiCtrlSet,
		apiSvcSet,
		dataRepoSet,
	)

	return nil
}
