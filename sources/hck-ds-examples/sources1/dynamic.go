package main

import "fmt"

type Dynamic struct {
	detailsMap map[string]interface{}
	isActive   bool
}

func main() {
	mp := make(map[string]interface{}{"Name": "Rob Pike", "Age": 24})
	me := Dynamic{mp, true}
	fmt.Println(me)
}
