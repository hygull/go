/*
	{
		"date_of_creation" => "16 Dec 2016, Thurs",
		"aim_of_program"   => "To generate MD5 hashes(checksums)",
		"coded_by"         => "Rishikesh Agrawani",
		"generate_online"  => "{ http://onlinemd5.com/ },
							        { http://www.md5.cz/ }",
		"Go_version"	    => "1.7",
	}
*/
package main

import "fmt"
import "crypto/md5"
import "encoding/hex"
import "io"

func main() {
	socialMsg := "Visit my 'Go Of Google' group on facebook"

	//1st way to generate md5 hash
	fmt.Println("\n1st way to generate md5 hash :-")
	md5Hasher := md5.New()
	md5Hasher.Write([]byte(socialMsg))
	md5Hash1 := hex.EncodeToString(md5Hasher.Sum(nil))
	fmt.Println("MD5 Hash for \""+socialMsg+"\" is ", md5Hash1)

	//2nd and direct way to generate md5 hash
	fmt.Println("\n2nd way to generate md5 hash :-")
	md5Hash2 := md5.Sum([]byte(socialMsg)) //Returns [16]byte
	/* fmt.Println(string(md5Hash2))  -->  cannot convert md5Hash2  (type [16]byte) to type string */
	fmt.Println(md5Hash2)                                                 //Prints a byte array of length 16
	fmt.Printf("%s %x\n", "MD5 Hash for \""+socialMsg+"\" is ", md5Hash2) /* %x format verb is used to convert hash result to a hex string*/

	//3rd way to generate md5 hash
	fmt.Println("\n3rd way to generate md5 hash :-")
	md5Hasher3 := md5.New()
	io.WriteString(md5Hasher3, socialMsg)
	fmt.Printf("%s %x\n", "MD5 Hash for \""+socialMsg+"\" is ", md5Hasher.Sum(nil)) /* %x format verb is used to convert hash result to a hex string*/
}

/*


1st way to generate md5 hash :-
MD5 Hash for "Visit my 'Go Of Google' group on facebook" is  4ea34dd54b698ea698531109d1c10fdd

2nd way to generate md5 hash :-
[78 163 77 213 75 105 142 166 152 83 17 9 209 193 15 221]
MD5 Hash for "Visit my 'Go Of Google' group on facebook" is  4ea34dd54b698ea698531109d1c10fdd

3rd way to generate md5 hash :-
MD5 Hash for "Visit my 'Go Of Google' group on facebook" is  4ea34dd54b698ea698531109d1c10fdd

*/
