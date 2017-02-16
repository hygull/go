package main

import _ "strconv"
import "fmt"
import "time"

func main() {
	t := time.Now()
	fmt.Printf("%T %v\n", t.UTC().Unix(), t.UTC().Unix())
	fmt.Printf("%T %v\n", t.UTC().UnixNano(), t.UTC().UnixNano())

}
