package main

import (
	"flag"
	"fmt"
	"net"
	"taxCalculator/proto"
	"taxCalculator/service"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 8080, "the grpc server port")
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logrus.Errorf("could not listen on tcp server at %d: %v", *port, err)
		return
	}

	grpcServer := grpc.NewServer()
	proto.RegisterTaxCalculatorServiceServer(grpcServer, service.NewTaxCalculatorService())

	logrus.Infof("gRPC server is trying to be up and listen on port %d", *port)
	if serveErr := grpcServer.Serve(listener); serveErr != nil {
		logrus.Errorf("could not serve tax calculator service : %v", serveErr)
		return
	}
	defer grpcServer.GracefulStop()
}
