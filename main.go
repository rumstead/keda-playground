package main

import (
	"google.golang.org/grpc"
	pb "keda-cnp-scaler/pkg/scalers/protos"
	"keda-cnp-scaler/pkg/scalers/static"
	"log"
	"net"
	"net/http"
)

func main() {
	port := ":6000"
	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	staticScaler := static.NewStaticScaler()
	err = setupHTTPServer(staticScaler)
	if err != nil {
		log.Fatal(err)
	}

	pb.RegisterExternalScalerServer(grpcServer, staticScaler)

	log.Printf("listening on %s\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

func setupHTTPServer(scaler *static.Scaler) error {
	http.HandleFunc("/switch", func(writer http.ResponseWriter, request *http.Request) {
		scaler.Swap()
	})
	return http.ListenAndServe(":8080", nil)
}
