package service

import (
	av1 "connect/gen/alpha/v1"
	av1c "connect/gen/alpha/v1/alphav1connect"
	"context"
	"core/log"
	"net/http"

	"github.com/bufbuild/connect-go"
)

type alphaServer struct{}

func newAlphaServer() *alphaServer {
	return &alphaServer{}
}

func (as *alphaServer) Info(
	ctx context.Context,
	req *connect.Request[av1.InfoRequest],
) (*connect.Response[av1.InfoResponse], error) {
	log.GetLogCtx(ctx).Info(req.Msg.Message)

	res := connect.NewResponse(&av1.InfoResponse{
		Message: req.Msg.Message,
	})

	res.Header().Set("Alpha-Version", "v1")

	return res, nil
}

type AlphaService struct {
	server *alphaServer
}

func NewAlphaService() *AlphaService {
	return &AlphaService{
		server: newAlphaServer(),
	}
}

func (as *AlphaService) Handler(opts ...connect.HandlerOption) (string, http.Handler) {
	return av1c.NewAlphaServiceHandler(as.server, opts...)
}
