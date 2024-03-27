package main

import (
	"apple/higher/v2/utils"
	"apple/higher/v2/v2_endpoint"
	"apple/higher/v2/v2_service"
	"apple/higher/v2/v2_transport"
	"net/http"
)

func main() {
	utils.NewLoggerServer()
	server := v2_service.NewService(utils.GetLogger())
	endpoints := v2_endpoint.NewEndPointServer(server, utils.GetLogger())
	httpHandler := v2_transport.NewHttpHandler(endpoints, utils.GetLogger())
	utils.GetLogger().Info("server run 0.0.0.0:8888")
	_ = http.ListenAndServe("0.0.0.0:8888", httpHandler)
}
