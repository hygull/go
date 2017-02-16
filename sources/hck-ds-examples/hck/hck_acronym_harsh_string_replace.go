package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var noOfDislikedWords, noOfWordsInSentence, i uint8
	var sentence string
	var acronym []byte

	fmt.Scanf("%d", &noOfDislikedWords)

	dislikedWordsList := make([]string, noOfDislikedWords)

	for i = 0; i < noOfDislikedWords; i++ {
		var word string
		fmt.Scanf("%s", &word)
		dislikedWordsList[i] = word
	}
	fmt.Println("ok...")
	fmt.Scanf("%d", &noOfWordsInSentence)
	//fmt.Scanln("%s", &sentence) /* i/p string*/
	reader := bufio.NewReader(os.Stdin)
	sentence, _ = reader.ReadString('\n')

	fmt.Println(sentence)
	for _, w := range dislikedWordsList {
		sentence = strings.Replace(sentence, w, "", -1)
	}
	// fmt.Println(sentence)
	// fmt.Println(sentence)
	//s3 := strings.Split(sentence, " ")
	// s4 := strings.Fields(sentence)
	// fmt.Printf("%T %v %v", s4, s4, len(s4))

	validWordsList := strings.Fields(sentence)
	fmt.Println(validWordsList)
	l := len(validWordsList)
	for i, w := range validWordsList {
		if i == l-1 {
			acronym = append(acronym, w[0]-32)
			break
		}
		acronym = append(acronym, w[0]-32, 46)
	}
	fmt.Println(string(acronym))
}

/*Steps of compilation with error fixing

admins-MacBook-Pro-3:GoFiles admin$ go run ./hck/hck_acronym_harsh_string_replace.go
3
app
boy
car
ok...
6
this is an app joker boy and car
this is an app joker boy and car

t.i.a.j.a
admins-MacBook-Pro-3:GoFiles admin$ go run ./hck/hck_acronym_harsh_string_replace.go
3
app
boy
car
ok...
this is an app joker boy and car
his is an app joker boy and car

?.?.?.?.?
admins-MacBook-Pro-3:GoFiles admin$ go run ./hck/hck_acronym_harsh_string_replace.go
3
app
boy
car
ok...
6
this is an app joker boy and car
this is an app joker boy and car

T.I.A.J.A
admins-MacBook-Pro-3:GoFiles admin$
*/

/*
admins-MacBook-Pro-3:GoFiles admin$ go run ./hck/hck_acronym_harsh_string_replace.go
5
hey
girls
i
am
single
ok...
11
hey all boys and girls welcome to hackerearth easy september challenge
hey all boys and girls welcome to hackerearth easy september challenge

A.B.A.W.T.H.E.S.C
*/
