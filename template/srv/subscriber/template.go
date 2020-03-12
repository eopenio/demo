package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	template "github.com/eopenio/examples/template/srv/proto/template"
)

type Template struct{}

func (e *Template) Handle(ctx context.Context, msg *template.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *template.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
