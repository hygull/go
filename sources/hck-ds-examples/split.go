package main

import "fmt"
import "strings" //Multiple imports are not required for importing multiple packages

func main() {

	/* Split() function's usage */
	fmt.Printf("***** %s *****\n", "Using Split() function...")
	//First example
	sentence1 := "This is Golang developed at google" //Only 1 space is there (between 2 consecutive words)
	wordsList := strings.Split(sentence1, " ")
	fmt.Println("Sentence    : ", sentence1)
	fmt.Println("Words list  : ", wordsList)
	fmt.Println("No. of words:", len(wordsList)) //Prints 6

	//Second example
	sentence2 := "This    is    Golang  developed at google" //4 spaces -> 4 spaces -> 2 spaces -> 1 space for remaining
	wordsList = strings.Split(sentence2, " ")
	fmt.Println("Sentence    : ", sentence2)
	fmt.Println("Words list  : ", wordsList)
	fmt.Println("No. of words: ", len(wordsList)) //Prints 13 [n spaces will add (n-1) black strings to the list]

	//Third example
	sentence3 := "https://www.hygull.com/images/nature/samsung_moblie.jpg"
	wordsList = strings.Split(sentence3, ".")
	fmt.Println("Sentence    : ", sentence3)
	fmt.Println("Words list  : ", wordsList)
	fmt.Println("No. of words:", len(wordsList)) //Prints 4

	if wordsList[len(wordsList)-1] != "jpg" {
		fmt.Println("\n", sentence3, " is not any URL of jpg image")
	} else {
		fmt.Println("\n", sentence3, " looks like an URL of jpg image")
	}

	/* Fields() function's usage */
	fmt.Printf("\n***** %s *****\n", "Using Fields() function...")

	//First example
	sentence1 = "This is Golang developed at google" //Only 1 space is there (between 2 consecutive words)
	wordsList = strings.Fields(sentence1)            //Fields don't think about multiple spaces it considers mutiple spaces as a single one
	fmt.Println("Sentence    : ", sentence1)
	fmt.Println("Words list  : ", wordsList)
	fmt.Println("No. of words:", len(wordsList)) //Prints 6

	//Second example
	sentence2 = "This    is    Golang  developed at google" //4 spaces -> 4 spaces -> 2 spaces -> 1 space for remaining
	wordsList = strings.Fields(sentence2)
	fmt.Println("Sentence    : ", sentence2)
	fmt.Println("Words list  : ", wordsList)
	fmt.Println("No. of words: ", len(wordsList)) //Still it prints 6

	//Third example
	sentence3 = "https://www.hygull.com/images/nature/samsung_moblie.jpg" //URL string
	wordsList = strings.Fields(sentence3)
	fmt.Println("Sentence    : ", sentence3)
	fmt.Println("Words list  : ", wordsList)
	fmt.Println("No. of words:", len(wordsList)) //Prints 1...as there's no any space in the URL

}

/*admins-MacBook-Pro-3:GoFiles admin$ go run split.go

***** Using Split() function... *****
Sentence    :  This is Golang developed at google
Words list  :  [This is Golang developed at google]
No. of words: 6
Sentence    :  This    is    Golang  developed at google
Words list  :  [This    is    Golang  developed at google]
No. of words:  13
Sentence    :  https://www.hygull.com/images/nature/samsung_moblie.jpg
Words list  :  [https://www hygull com/images/nature/samsung_moblie jpg]
No. of words: 4

 https://www.hygull.com/images/nature/samsung_moblie.jpg  looks like an URL of jpg image

***** Using Fields() function... *****
Sentence    :  This is Golang developed at google
Words list  :  [This is Golang developed at google]
No. of words: 6
Sentence    :  This    is    Golang  developed at google
Words list  :  [This is Golang developed at google]
No. of words:  6
Sentence    :  https://www.hygull.com/images/nature/samsung_moblie.jpg
Words list  :  [https://www.hygull.com/images/nature/samsung_moblie.jpg]
No. of words: 1

*/
