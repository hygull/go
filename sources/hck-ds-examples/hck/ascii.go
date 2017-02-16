package main

import "fmt"

func main() {
	for i := 65; i < 91; i++ {
		fmt.Println(string(i), i)
	}
	for i := 97; i < 123; i++ {
		fmt.Println(string(i), i)
	}

}
