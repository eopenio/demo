package main

import (
	"github.com/eopenio/demo/template/srv/handler"
	_ "github.com/eopenio/demo/template/srv/initial"
	"github.com/eopenio/demo/template/srv/model"
	template "github.com/eopenio/demo/template/srv/proto/template"
	"github.com/eopenio/util/logutil"
	"github.com/eopenio/util/terror"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("eopenio.micro.srv.template"),
		micro.Version("v1.0"),
	)

	// Initialise service
	service.Init(

		// Register Handler
		micro.Action(func(c *cli.Context) error {
			logutil.BgLogger().Info("template-srv start...")
			model.Init()
			terror.MustNil(template.RegisterTemplateHandler(service.Server(), new(handler.Template)))
			return nil
		}),
		micro.AfterStop(func() error {
			logutil.BgLogger().Info("template-srv stop...")
			return nil
		}),
		micro.AfterStart(func() error {
			return nil
		}),
	)

	// Run service
	if err := service.Run(); err != nil {
		logutil.BgLogger().Error(err.Error())
	}
}
