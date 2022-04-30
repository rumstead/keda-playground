package main

import (
	"google.golang.org/grpc"
	pb "keda-cnp-scaler/pkg/scalers/protos"
	"keda-cnp-scaler/pkg/scalers/static"
	"log"
	"net"
)

func main() {
	port := ":6000"
	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	pb.RegisterExternalScalerServer(grpcServer, &static.Scaler{})

	log.Printf("listening on %s\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
