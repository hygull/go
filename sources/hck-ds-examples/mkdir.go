/*
	{
		"cretaed_after" : "Sat Dec 10 12:01:49 IST 2016"
		"aim_of_program" : "To create a new directory in current working directory(On MAC/UNIX)"
		"coded_by" : "Rishikesh Agrawani"
	}
*/
package main

import "os"
import "fmt"

func main() {
	//This is for MAC Or UNIX machines that uses / as path separator.
	//For platform independent path separators check my next gist(on github)/post
	if _, err := os.Stat("./UploadFile"); os.IsNotExist(err) {
		os.Mkdir("uploadFile", 0777)
		fmt.Println("Directory created.")
	} else {
		fmt.Println("Directory exists.")
	}
}

/*FIRST RUN:-
				Directory created.

  SECOND RUN:-
  				Directory exists.
*/
