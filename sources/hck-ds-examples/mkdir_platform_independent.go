/*
	{
		"cretaed_after" : "Sat Dec 10 12:01:49 IST 2016"
		"aim_of_program" : "To create a new directory in current working directory(On MAC/UNIX/Windows)"
		"coded_by" : "Rishikesh Agrawani"
	}
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	//In windows  -->    .\GolangDir
	//In Unix/MAC -->    ./GolangDir
	//Lets do platform independent implementation for making directory in current working directory
	if _, err := os.Stat("." + string(filepath.Separator) + "GolangDir"); os.IsNotExist(err) {
		os.Mkdir("."+string(filepath.Separator)+"GolangDir", 0777)
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
