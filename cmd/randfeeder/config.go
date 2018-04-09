package main

type config struct {
	Port              int    `envconfig:"PORT" default:"6010"`
	DataMean		  float64 `split_words:"true" required:"true"`
	DataVar			  float64 `split_words:"true" required:"true"`
	NoiseMean		  float64 `split_words:"true" required:"true"`
	NoiseVar		  float64 `split_words:"true" required:"true"`
}
