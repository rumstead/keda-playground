package main

import (
	"fmt"
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
	setupHTTPServer(staticScaler)

	pb.RegisterExternalScalerServer(grpcServer, staticScaler)

	log.Printf("listening on %s\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

func setupHTTPServer(scaler *static.Scaler) {
	http.HandleFunc("/switch", func(writer http.ResponseWriter, request *http.Request) {
		old := scaler.Swap()
		response := fmt.Sprintf("From %t to %t\n", old, scaler.Down())
		_, _ = writer.Write([]byte(response))
	})
	go func() error {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}()
}
