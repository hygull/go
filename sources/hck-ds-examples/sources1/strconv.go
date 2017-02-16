package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*----------- Without panic() -----------------------------------------*/
	a, err := strconv.Atoi("1234")
	fmt.Println(a, err)

	b, err := strconv.Atoi("abcd")
	fmt.Println(b, err)

	fmt.Println("Nice", b)

	/*
		1234 <nil>
		0 strconv.ParseInt: parsing "abcd": invalid syntax
		Nice 0
	*/

	/*-------------- Recovering panic() -----------------------------------*/
	a2, err2 := strconv.Atoi("1234")
	fmt.Println(a2, err2)

	b2, err2 := strconv.Atoi("abcd")
	//panic(err2)
	recoverMsg := recover()

	fmt.Println("Nice", b2)
	fmt.Println("Recover message:- ", recoverMsg)

	/*----------- With panic() ---------------------------------------------*/
	a1, err1 := strconv.Atoi("1234")
	fmt.Println(a1, err1)

	b1, err1 := strconv.Atoi("abcd")
	panic(err1)
	fmt.Println("Nice", b1) //It will not be printed

	/*
		1234 <nil>
		panic: strconv.ParseInt: parsing "abcd": invalid syntax

		goroutine 1 [running]:
		panic(0x8f8c0, 0xc42007e060)
			/usr/local/go/src/runtime/panic.go:500 +0x1a1
		main.main()
			/Users/admin/projects/Go/GoFiles/sources1/strconv.go:29 +0x45a
		exit status 2
	*/

}
