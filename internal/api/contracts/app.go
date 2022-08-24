package contracts

import "github.com/kataras/iris/v12"

type App struct {
	Config     map[string]string
	Datasource *Datasources
	Iris       *iris.Application
	Service    *Services
}
