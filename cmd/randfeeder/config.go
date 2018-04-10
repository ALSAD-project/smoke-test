package main

type config struct {
	Port             int     `envconfig:"PORT" default:"6010"`
	DataVariance     float64 `split_words:"true" required:"true"`
	NoiseMagnitude   float64 `split_words:"true" required:"true"`
	NoiseVariance    float64 `split_words:"true" required:"true"`
	NoiseProbability float64 `split_words:"true" default:"0.3"`
}
