package main

import (
	"fmt"
	"math"
)

func Round(input float64) float64 {
	if input < 0 {
		return math.Ceil(input - 0.5)
	}
	return math.Floor(input + 0.5)
}

func RoundUp(input float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * input
	round = math.Ceil(digit)
	newVal = round / pow
	return
}

func RoundDown(input float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * input
	round = math.Floor(digit)
	newVal = round / pow
	return
}

func main() {
	var f float64 = 123.123456

	fmt.Printf("%0.2f \n", f)

	fmt.Printf("%0.2f \n", Round(f)) // round half

	fmt.Printf("%0.2f \n", RoundUp(f, 2))

	fmt.Printf("%0.2f \n", RoundDown(f, 2))

}
