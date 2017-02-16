package main

import "fmt"

type Float64 float64
type FloatSlice64 []float64
type Intf interface {
	Sort()
}

func (f Float64) Sort() {
	fmt.Println(f)
}

func (f FloatSlice64) Sort() {
	fmt.Println(f[1])
}
func Print(i interface{}) {

	switch t := i.(type) {

	case []int:
		fmt.Println("[]int...")

	case []int8:
		fmt.Println("[]int8...")
	case []int16:
		fmt.Println("[]int16...")
	case []int32:
		fmt.Println("[]int32...")
	case []int64:
		fmt.Println("[]int64...")
	case []float64:
		fmt.Println("[]float64...")
	case []float32:
		fmt.Println("[]float32...")

	default:
		fmt.Println("I don't want this value of type\n", t)
	}
}

func main() {
	Print(1)
	Print(int8(2))
	Print(float64(3))
	Print(7.7)
	Print(-4)
	Print(0)
	Print([]int{0, -5, 53, 51, 47})
	Print([]float32{0, -5.3, 53.56, 51, 47.5})
	// var i Intf
	v := Float64(5)
	fs := FloatSlice64([]float64{45.56, 32.45, 8757.445, 85885.5454, 455454.656})
	v.Sort()
	fs.Sort()
}
