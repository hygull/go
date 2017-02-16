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

	fmt.Scanf("%d", &noOfWordsInSentence)

	reader := bufio.NewReader(os.Stdin)
	sentence, _ = reader.ReadString('\n')

	for _, w := range dislikedWordsList {
		sentence = strings.Replace(sentence, w, "", -1)
	}
	validWordsList := strings.Fields(sentence)

	fmt.Println("Ok", validWordsList)
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

/* INPUT:-

5
hey
girls
i
am
single
11
hey all boys and girls welcome to hackerearth easy september challenge
*/

/* OUTPUT:-

A.B.A.W.T.H.E.S.C
*/
