package main

import "fmt"
import "errors"
import "ds"
import "github.com/fatih/color"

type Float32Slice []float32
type Float64Slice []float64
type IntSlice []int
type Int8Slice []int8
type Int16Slice []int16
type Int32Slice []int32
type Int64Slice []int64

func (f32Slice Float32Slice) Sort() {
	for i := len(f32Slice) - 2; i >= 0; i-- {
		swapped := false
		for j := 0; j <= i; j++ {
			if f32Slice[j] > f32Slice[j+1] {
				f32Slice[j], f32Slice[j+1] = f32Slice[j+1], f32Slice[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func (f64Slice Float64Slice) Sort() {
	for i := len(f64Slice) - 2; i >= 0; i-- {
		swapped := false
		for j := 0; j <= i; j++ {
			if f64Slice[j] > f64Slice[j+1] {
				f64Slice[j], f64Slice[j+1] = f64Slice[j+1], f64Slice[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func (iSlice IntSlice) Sort() {
	for i := len(iSlice) - 2; i >= 0; i-- {
		swapped := false
		for j := 0; j <= i; j++ {
			if iSlice[j] > iSlice[j+1] {
				iSlice[j], iSlice[j+1] = iSlice[j+1], iSlice[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func (i8Slice Int8Slice) Sort() {
	for i := len(i8Slice) - 2; i >= 0; i-- {
		swapped := false
		for j := 0; j <= i; j++ {
			if i8Slice[j] > i8Slice[j+1] {
				i8Slice[j], i8Slice[j+1] = i8Slice[j+1], i8Slice[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func (i16Slice Int16Slice) Sort() {
	for i := len(i16Slice) - 2; i >= 0; i-- {
		swapped := false
		for j := 0; j <= i; j++ {
			if i16Slice[j] > i16Slice[j+1] {
				i16Slice[j], i16Slice[j+1] = i16Slice[j+1], i16Slice[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func (i32Slice Int32Slice) Sort() {
	for i := len(i32Slice) - 2; i >= 0; i-- {
		swapped := false
		for j := 0; j <= i; j++ {
			if i32Slice[j] > i32Slice[j+1] {
				i32Slice[j], i32Slice[j+1] = i32Slice[j+1], i32Slice[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func (i64Slice Int64Slice) Sort() {
	for i := len(i64Slice) - 2; i >= 0; i-- {
		swapped := false
		for j := 0; j <= i; j++ {
			if i64Slice[j] > i64Slice[j+1] {
				i64Slice[j], i64Slice[j+1] = i64Slice[j+1], i64Slice[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func BubbleSort(listOfNums interface{}) error {
	var err error

	switch t := listOfNums.(type) {
	case []int:
		IntSlice(listOfNums.([]int)).Sort()
	case []int8:
		Int8Slice(listOfNums.([]int8)).Sort()
	case []int16:
		Int16Slice(listOfNums.([]int16)).Sort()
	case []int32:
		Int32Slice(listOfNums.([]int32)).Sort()
	case []int64:
		Int64Slice(listOfNums.([]int64)).Sort()
	case []float32:
		Float32Slice(listOfNums.([]float32)).Sort()
	case []float64:
		Float64Slice(listOfNums.([]float64)).Sort()
	default:
		color.Green("[hygull:ds] Got =>>>>> %v", t)
		color.Red("[hygull:ds] You have to .... pass any one among these => []int, []int8, []int16, []int32, []int64, []float32, []float64")

		err = errors.New("Didn't get any proper slice of numbers")
		//panic(err)
	}
	return err
}

func main() {
	a := []int{0, -5, 53, 51, 56, 47}
	fmt.Println("a:", a)
	BubbleSort(a)
	fmt.Println("a:", a)

	f := []float32{}
	fmt.Println("f:", f)
	ds.BubbleSort(f)
	fmt.Println("f:", f)

	s := []string{"Apple", "Ball"}
	fmt.Println("s:", s)
	ds.BubbleSort(s)
	fmt.Println("f:", s)
	color.Green("Rishikesh")
}
