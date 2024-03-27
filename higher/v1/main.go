package main

import (
	"apple/higher/v1/v1_endpoint"
	"apple/higher/v1/v1_service"
	"apple/higher/v1/v1_transport"
	"fmt"
	"net/http"
)

func main() {
	server := v1_service.NewService()
	endpoints := v1_endpoint.NewEndPointServer(server)
	httpHandler := v1_transport.NewHttpHandler(endpoints)
	fmt.Println("server run 0.0.0.0:8888")
	_ = http.ListenAndServe("0.0.0.0:8888", httpHandler)
}
