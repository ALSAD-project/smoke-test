package handler

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"time"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat/distmv"
)

type randHandler struct {
	dist *distmv.Normal
	noiseProbability float64
	channel chan []float64
}

func NewRandHandler(
	variance float64,
	noiseMean float64,
	noiseVariance float64,
	noiseProbability float64,
) (http.Handler, error) {
	dist, ok := distmv.NewNormal(
		[]float64{0, 0, noiseMean},
		mat.NewSymDense(
			3,
			[]float64{
				variance, 0, 0,
				0, variance, 0,
				0, 0, noiseVariance,
			},
	),
		rand.New(rand.NewSource(uint64(time.Now().Unix()))),
	)
	if !ok {
		return nil, errors.New("failed to create normal distribution generator")
	}

	h := randHandler{
		dist:             dist,
		noiseProbability: noiseProbability,
		channel:          make(chan []float64, 10),
	}

	go h.startGeneration()

	return &h, nil
}

func (h *randHandler) startGeneration() {
	for {
		x := h.dist.Rand(nil)
		xVal := x[0]
		yVal := x[1]
		noise := x[2]

		if rand.Float64() < h.noiseProbability {
			scale := math.Sqrt(noise * noise / (xVal * xVal + yVal * yVal))
			xVal = xVal * scale
			yVal = yVal * scale
		}

		h.channel <- []float64{xVal, yVal}
	}
}

func (h *randHandler) getRand()(float64, float64) {
	r := <-h.channel
	return r[0], r[1]
}

func (h *randHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	x, y := h.getRand()

	w.Header().Set("Content-Type", "text/csv")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%.3f, %.3f", x, y)))
}
