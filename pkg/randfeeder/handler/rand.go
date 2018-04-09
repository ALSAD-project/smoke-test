package handler

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat/distmv"
)

type randHandler struct {
	dist *distmv.Normal
}

func NewRandHandler(mean, variance, noiseMean, noiseVariance float64) (http.Handler, error) {
	dist, ok := distmv.NewNormal(
		[]float64{mean, noiseMean},
		mat.NewSymDense(2, []float64{variance, 0, 0, noiseVariance}),
		rand.New(rand.NewSource(uint64(time.Now().Unix()))),
	)
	if !ok {
		return nil, errors.New("failed to create normal distribution generator")
	}

	return &randHandler{
		dist: dist,
	}, nil
}

func (h *randHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v := make([]float64, 2)
	v = h.dist.Rand(v)

	data := strconv.FormatFloat(v[0] + v[1], 'f', 6, 64)
	noise := strconv.FormatFloat(v[1], 'f',6,64)
	response := data + "," + noise

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
