/*6 jan 2017
https://www.hackerearth.com/practice/basic-programming/implementation/basics-of-implementation/practice-problems/algorithm/acronym-2/
My fast sumission at : https://www.hackerearth.com/submission/6693800/
*/

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
	var lengths []int
	//var sortedWords []string
	//numToStringMap := make(map[int]string)

	fmt.Scanf("%d", &noOfDislikedWords)

	dislikedWordsList := make([]string, noOfDislikedWords)

	for i = 0; i < noOfDislikedWords; i++ {
		var word string
		fmt.Scanf("%s", &word)

		dislikedWordsList[i] = word

		lengths = append(lengths, len(word))

	}
	fmt.Println(lengths)
	fmt.Println(dislikedWordsList)
	for k := len(lengths) - 2; k >= 0; k-- {
		swapped := false
		for j := 0; j <= k; j++ {
			if lengths[j+1] > lengths[j] {
				//[j+1], dislikedWordsList[j] = dislikedWordsList[j], dislikedWordsList[j+1]
				lengths[j+1], lengths[j] = lengths[j], lengths[j+1]
				dislikedWordsList[j+1], dislikedWordsList[j] = dislikedWordsList[j], dislikedWordsList[j+1]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	fmt.Println(dislikedWordsList)
	fmt.Println(lengths)
	fmt.Scanf("%d", &noOfWordsInSentence)

	reader := bufio.NewReader(os.Stdin)
	sentence, _ = reader.ReadString('\n')

	for i, _ := range lengths {
		// fmt.Println("replacing-->", numToStringMap[index])
		sentence = strings.Replace(sentence, dislikedWordsList[i], "", -1)
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
