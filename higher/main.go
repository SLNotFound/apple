package main

import (
	"apple/higher/endpoint"
	"apple/higher/service"
	"apple/higher/transport"
	"apple/higher/utils"
	"go.uber.org/ratelimit"
	"golang.org/x/time/rate"
	"net/http"
)

func main() {
	utils.NewLoggerServer()
	golangLimit := rate.NewLimiter(10, 1)
	uberLimit := ratelimit.New(1)
	server := service.NewService(utils.GetLogger())
	endpoints := endpoint.NewEndPointServer(server, utils.GetLogger(), golangLimit, uberLimit)
	httpHandler := transport.NewHttpHandler(endpoints, utils.GetLogger())
	utils.GetLogger().Info("server run 0.0.0.0:8888")
	_ = http.ListenAndServe("0.0.0.0:8888", httpHandler)
}
