package subscriber

import (
	"context"
	"github.com/eopenio/util/logutil"
	"go.uber.org/zap"

	template "github.com/eopenio/demo/template/srv/proto/template"
)

type Template struct{}

func (e *Template) Handle(ctx context.Context, msg *template.Message) error {
	logutil.BgLogger().Info("Handler Received message: ", zap.String("msg.Say", msg.Say))
	return nil
}

func Handler(ctx context.Context, msg *template.Message) error {
	logutil.BgLogger().Info("Function Received message: ", zap.String("msg.Say", msg.Say))
	return nil
}
