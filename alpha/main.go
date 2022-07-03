package main

import (
	"alpha/service"
	av1c "connect/gen/alpha/v1/alphav1connect"
	"context"
	"core/log"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bufbuild/connect-go"
	grpchealth "github.com/bufbuild/connect-grpchealth-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const memorySize = 1024

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	compress := connect.WithCompressMinBytes(memorySize)

	mux := http.NewServeMux()

	mux.Handle(service.NewAlphaService().Handler())

	mux.Handle(grpchealth.NewHandler(
		grpchealth.NewStaticChecker(av1c.AlphaServiceName),
		compress,
	))

	server := &http.Server{
		Addr:    ":" + port,
		Handler: h2c.NewHandler(log.Middleware(mux), &http2.Server{}),
	}

	signals := make(chan os.Signal, 1)

	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Log().Info("server run ...")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Log().Sugar().Fatal("HTTP listen and serve: %v", err)
		}
	}()

	<-signals

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Log().Sugar().Fatal("HTTP shutdown: %v", err)
	}
}
