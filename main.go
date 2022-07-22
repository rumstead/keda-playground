package main

import (
	"log"
	"net/http"

	"keda-cnp-scaler/pkg/scalers/static"
)

func main() {
	staticScaler := static.NewStaticScaler()
	setupHTTPServer(staticScaler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func setupHTTPServer(scaler *static.Scaler) {
	http.HandleFunc("/switch", scaler.HandleSwap)
	http.HandleFunc("/scale", scaler.HandleScale)
}
