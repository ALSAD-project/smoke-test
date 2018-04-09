package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/kelseyhightower/envconfig"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat/distmv"
)

func main() {
	rfConfig := config{}
	if err := envconfig.Process("rf", &rfConfig); err != nil {
		log.Fatalf("Error on processing configuration: %s", err.Error())
		return
	}

	dist, ok := distmv.NewNormal(
		[]float64{rfConfig.DataMean, rfConfig.NoiseMean},
		mat.NewSymDense(2, []float64{rfConfig.DataVar, 0, 0, rfConfig.NoiseVar}),
		rand.New(rand.NewSource(1234)),
	)
	if !ok {
		log.Fatalf("Error on creating data source")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		respHandler(resp, req, dist)
	})

	log.Printf("Server is listening on port %d", rfConfig.Port)
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%d", rfConfig.Port),
		mux,
	))
}

func respHandler(w http.ResponseWriter, r *http.Request, dist *distmv.Normal) {
	log.Println("Accepted new connection.")

	v := make([]float64, 2)
	v = dist.Rand(v)

	data := strconv.FormatFloat(v[0] + v[1], 'f', 6, 64)
	noise := strconv.FormatFloat(v[1], 'f',6,64)
	response := data + "," + noise

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
