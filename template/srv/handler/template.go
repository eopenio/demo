package handler

import (
	"context"
	"go.uber.org/zap"

	"github.com/eopenio/util/logutil"

	template "github.com/eopenio/demo/template/srv/proto/template"
)

type Template struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Template) Call(ctx context.Context, req *template.Request, rsp *template.Response) error {
	rsp.Msg = "Hello " + req.Name
	logutil.BgLogger().Info("Received Template.Call request", zap.String("rsp.Msg", rsp.Msg))
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Template) Stream(ctx context.Context, req *template.StreamingRequest, stream template.Template_StreamStream) error {
	logutil.BgLogger().Info("Received Template.Stream request with count: %d", zap.Int64("req.Count", req.Count))

	for i := 0; i < int(req.Count); i++ {
		logutil.BgLogger().Info("Responding: %d", zap.Int("i", i))
		if err := stream.Send(&template.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Template) PingPong(ctx context.Context, stream template.Template_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		logutil.BgLogger().Info("Got ping %v", zap.Int64("req.Stroke", req.Stroke))
		if err := stream.Send(&template.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
