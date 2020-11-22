package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"taxCalculator/proto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 0, "the rest server port")
	endPoint := flag.String("endpoint", ":8080", "gRPC rest endpoint")
	flag.Parse()

	muxServer := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logrus.Info("registering taxCalculator service handler endpoint for rest")
	if err := proto.RegisterTaxCalculatorServiceHandlerFromEndpoint(
		ctx, muxServer, *endPoint, []grpc.DialOption{grpc.WithInsecure()},
	); err != nil {
		logrus.Errorf("could not register tax calculator service endpoint : %v", err)
		return
	}

	logrus.Info("listening and serving taxCalculator service handler endpoint for rest")
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port),
		muxServer); err != nil {
		logrus.Errorf("could not serve : %v", err)
		return
	}
}
