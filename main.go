package main

import (
	"flag"
	"fmt"
	"log"
	"math"
)

var (
	availBandwidths       = []float64{10.4, 15.6, 20.8, 31.25, 41.7, 62.5, 125.0, 250.0, 500.0}
	availSpreadingFactors = []float64{6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0}
	availCodeRates        = []float64{1.0, 2.0, 3.0, 4.0}
)

func main() {
	sf := flag.Float64("sf", 9.0, "Spreading factor")
	cr := flag.Float64("cr", 2.0, "Code rate")
	bw := flag.Float64("bandwidth", 125.0, "Bandwidth in kHz")
	flag.Parse()

	if err := validateInputs(*bw, *sf, *cr); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data rate = %.6fkbps for Bandwidth = %f, Code Rate = %f, Spreading Factor = %f\n",
		calcLoraDataRate(*bw, *sf, *cr), *bw, *cr, *sf)

}

func calcLoraDataRate(bandwidth float64, spreadingFactor float64, codeRate float64) float64 {
	return spreadingFactor * (bandwidth / math.Pow(2.0, spreadingFactor)) * (4.0 / (codeRate + 4.0))
}

func validateInputs(bandwidth float64, spreadingFactor float64, codeRate float64) error {
	if !inArray(bandwidth, availBandwidths) {
		return fmt.Errorf("Bandwidth: %f is not an available bandwidth. Available Bandwidths: %v", bandwidth, availBandwidths)
	}

	if !inArray(spreadingFactor, availSpreadingFactors) {
		return fmt.Errorf("Spreading Factor: %f is not an available spreading factor. Available Spreading Factors: %v", spreadingFactor, availSpreadingFactors)
	}

	if !inArray(codeRate, availCodeRates) {
		return fmt.Errorf("Code rate: %f is not an available code rate. Available Code Rates: %v", codeRate, availCodeRates)
	}

	return nil
}

func inArray(in float64, arr []float64) bool {
	for _, item := range arr {
		if in == item {
			return true
		}
	}

	return false
}
