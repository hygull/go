package main

import "fmt"
import "errors"

type Float32Slice []float32
type Float64Slice []float64
type IntSlice []int
type Int8Slice []int8
type Int16Slice []int16
type Int32Slice []int32

func (f32Slice Float32Slice) Sort() {

}

func (f64Slice Float64Slice) Sort() {

}

func (iSlice IntSlice) Sort() {
	for i := len(iSlice) - 2; i >= 0; i-- {
		swapped := false
		for j := 0; j <= i; j++ {
			if iSlice[j] > iSlice[j+1] {
				iSlice[j] = iSlice[j] + iSlice[j+1]
				iSlice[j+1] = iSlice[j] - iSlice[j+1]
				iSlice[j] = iSlice[j] - iSlice[j+1]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func (i8Slice Int8Slice) Sort() {

}

func (i16Slice Int16Slice) Sort() {

}

func (i32Slice Int32Slice) Sort() {

}

func BubbleSort(listOfNums interface{}) error {

	switch t := listOfNums.(type) {

	case []int:
		IntSlice(listOfNums.([]int)).Sort()
	// case []int8:

	// case []int16:

	// case []int32:

	// case []int64:

	// case []float64:

	// case []float32:

	default:
		fmt.Println("Got ", t)
		return errors.New("Didn't get any proper slice of numbers")
	}
	return nil
}

func ShowError(errMsg string) {

}

func main() {
	a := []int{0, -5, 53, 51, 47}
	fmt.Println("a:", a)
	BubbleSort(a)
	fmt.Println("a:", a)

	BubbleSort([]float32{0, -5.3, 53.56, 51, 47.5})

}
