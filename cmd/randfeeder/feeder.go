package main

import (
	"fmt"
	"log"
	"math"
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
		rfConfig.DataVariance,
		math.Abs(rfConfig.NoiseMagnitude),
		rfConfig.NoiseVariance,
		math.Min(1, math.Abs(rfConfig.NoiseProbability)),
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