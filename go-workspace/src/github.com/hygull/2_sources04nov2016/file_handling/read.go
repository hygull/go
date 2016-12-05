package main

import "io/ioutil"
import "fmt"
import "os"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You didn't provide the file name while running go run " + os.Args[0])
		return
	}
	fmt.Println(os.Args)
	filePath := os.Args[1] //It's not required...this is for understanding what
	//is os.Args[1]
	filesBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error while reading file, check the path please...")
		return
	}
	fmt.Println("Contents of " + os.Args[0] + " :\n\n")
	fmt.Println(string(filesBytes))
}
