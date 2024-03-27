package main

import (
	"apple/higher/endpoint"
	"apple/higher/service"
	"apple/higher/transport"
	"apple/higher/utils"
	"net/http"
)

func main() {
	utils.NewLoggerServer()
	server := service.NewService(utils.GetLogger())
	endpoints := endpoint.NewEndPointServer(server, utils.GetLogger())
	httpHandler := transport.NewHttpHandler(endpoints, utils.GetLogger())
	utils.GetLogger().Info("server run 0.0.0.0:8888")
	_ = http.ListenAndServe("0.0.0.0:8888", httpHandler)
}
