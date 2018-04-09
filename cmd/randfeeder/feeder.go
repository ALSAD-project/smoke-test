package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ALSAD-project/smoke-test/pkg/randfeeder/handler"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	rfConfig := config{}
	if err := envconfig.Process("rf", &rfConfig); err != nil {
		log.Fatalf("Error on processing configuration: %s", err.Error())
		return
	}

	hdr, err := handler.NewRandHandler(
		rfConfig.DataMean,
		rfConfig.DataVar,
		rfConfig.NoiseMean,
		rfConfig.NoiseVar,
	)
	if err != nil {
		log.Fatalf("Failed to create handler: %s", err.Error())
		return
	}

	mux := http.NewServeMux()
	mux.Handle("/", hdr)

	log.Printf("Server is listening on port %d", rfConfig.Port)
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%d", rfConfig.Port),
		mux,
	))
}