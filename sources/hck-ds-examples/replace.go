/*
	{
		"created_before" 		=>	"Thu Dec 15 09:45:14 IST 2016",
		"aim_of_program" 	=>	"To use string's Replace() function for replacing text(s) from a string",
		"coded_by" 			=>	"Rishikesh Agrawani",
	}
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	/* List of 5 examples for Replace() function */
	/* Number of occurences of character 'J' is 7*/
	fmt.Println(strings.Replace("JNDJA_JS_JN_MY_JNSPJRATJON_LJST", "J", "I", 1))
	fmt.Println(strings.Replace("JNDJA_JS_JN_MY_JNSPJRATJON_LJST", "J", "I", 2))
	fmt.Println(strings.Replace("JNDJA_JS_JN_MY_JNSPJRATJON_LJST", "J", "I", 7))
	fmt.Println(strings.Replace("JNDJA_JS_JN_MY_JNSPJRATJON_LJST", "J", "I", 8))
	fmt.Println(strings.Replace("JNDJA_JS_JN_MY_JNSPJRATJON_LJST", "J", "I", -1))
	fmt.Println()

	/* 2nd example with multi-word replacement */
	s := "Golang is nice and C is also nice. It's a nice thing to learn both at nice places"
	fmt.Println(strings.Replace(s, "nice", "great", 1))
	fmt.Println(strings.Replace(s, "nice", "great", 2))
	fmt.Println(strings.Replace(s, "nice", "great", -1))
	fmt.Println()

	/* 3rd example using a slice of strings */
	mainStrings := []string{"JNDJA", "CommittWW", "Gish Gor My Gather", "Cute@Cat@On@My@House"} //strings in predefined package

	replacementStrings := []string{"J", "W", "G", "@"}

	replacingStrings := []string{"I", "e", "F", " "}

	for i, str := range mainStrings {
		fmt.Println(strings.Replace(str, replacementStrings[i], replacingStrings[i], -1))
	}
}

/*

	INDJA_JS_JN_MY_JNSPJRATJON_LJST
	INDIA_JS_JN_MY_JNSPJRATJON_LJST
	INDIA_IS_IN_MY_INSPIRATION_LJST
	INDIA_IS_IN_MY_INSPIRATION_LIST
	INDIA_IS_IN_MY_INSPIRATION_LIST

	Golang is great and C is also nice. It's a nice thing to learn both at nice places
	Golang is great and C is also great. It's a nice thing to learn both at nice places
	Golang is great and C is also great. It's a great thing to learn both at great places

	INDIA
	Committee
	Fish For My Father
	Cute Cat On My House

*/
