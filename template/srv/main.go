package main

import (
	"github.com/eopenio/examples/template/srv/handler"
	_ "github.com/eopenio/examples/template/srv/initial"
	template "github.com/eopenio/examples/template/srv/proto/template"
	"github.com/eopenio/util/logutil"
	"github.com/micro/go-micro/v2"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("eopenio.micro.srv.template"),
		micro.Version("v1.0"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	template.RegisterTemplateHandler(service.Server(), new(handler.Template))

	// Run service
	if err := service.Run(); err != nil {
		logutil.BgLogger().Error(err.Error())
	}
}
