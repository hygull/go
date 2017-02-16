package main

import "fmt"

type Obj struct {
	Record   map[string]interface{}
	isActive bool
}

func storeDetails(data []interface{}, isActice bool) Obj {
	tempMap := make(map[string]interface{})
	obj := Obj{}
	dataName := "data"
	for index, item := range data {
		switch t := item.(type) {
		case int:
			tempMap[dataName+fmt.Sprint(index)] = item.(int)
		case int8:
			tempMap[dataName+fmt.Sprint(index)] = item.(int8)
		case int16:
			tempMap[dataName+fmt.Sprint(index)] = item.(int16)
		case int32:
			tempMap[dataName+fmt.Sprint(index)] = item.(int32)
		case int64:
			tempMap[dataName+fmt.Sprint(index)] = item.(int64)
		case uint:
			tempMap[dataName+fmt.Sprint(index)] = item.(uint)
		case uint8:
			tempMap[dataName+fmt.Sprint(index)] = item.(uint8)
		case uint16:
			tempMap[dataName+fmt.Sprint(index)] = item.(uint16)
		case uint32:
			tempMap[dataName+fmt.Sprint(index)] = item.(uint32)
		case uint64:
			tempMap[dataName+fmt.Sprint(index)] = item.(uint64)
		// case rune:    /*  ./test.go:35: duplicate case rune in type switch previous case at ./test.go:21 */
		// 	tempMap[dataName+fmt.Sprint(index)] = item.(rune)
		// case byte:    /*   ./test.go:37: duplicate case byte in type switch previous case at ./test.go:27*/
		// 	tempMap[dataName+fmt.Sprint(index)] = item.(byte)
		case bool:
			tempMap[dataName+fmt.Sprint(index)] = item.(bool)
		case uintptr:
			tempMap[dataName+fmt.Sprint(index)] = item.(uintptr)
		case complex64:
			tempMap[dataName+fmt.Sprint(index)] = item.(complex64)
		case complex128:
			tempMap[dataName+fmt.Sprint(index)] = item.(complex128)
		case float32:
			tempMap[dataName+fmt.Sprint(index)] = item.(float32)
		case float64:
			tempMap[dataName+fmt.Sprint(index)] = item.(float64)
		case string:
			tempMap[dataName+fmt.Sprint(index)] = item.(string)
		default:
			fmt.Println("Got ", t)
			panic("Illegal item ")

		}
	}
	obj.Record = tempMap
	obj.isActive = isActice
	return obj
}
func main() {
	// fmt.Printf("%T\n", 45.5)
	// fmt.Printf("%T\n", 56)

	fmt.Println(storeDetails([]interface{}{"Rishikesh", 24, 45.56}, true))
}
