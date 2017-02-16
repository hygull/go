/*
	{
		"date_of_creation" => "16 Dec 2016, Thurs",
		"aim_of_program"   => "To generate MD5 hashes(checksums)",
		"coded_by"         => "Rishikesh Agrawani",
		"generate_online"  => "{ http://onlinemd5.com/ },
							   { http://www.md5.cz/ }",
		"Go_version"	   => "1.7",
	}
*/
package main

import "fmt"
import "crypto/sha1"

import "io"

func main() {
	lang := "Golang"
	//1st way to create SHA1 hash
	sha1Hasher := sha1.New()
	sha1Hasher.Write([]byte(lang))
	sha1Hash := sha1Hasher.Sum(nil)
	fmt.Printf("%s %x\n\n", "SHA1 Hash for \""+lang+"\" is ", sha1Hash) /* %x format verb is used to convert hash result to a hex string*/

	//2nd way to generate SHA1 hash
	sha1Hash2 := sha1.Sum([]byte(lang))
	fmt.Println(sha1Hash2)
	/* fmt.Println(string(sha1Hash2))  -->  cannot convert sha1Hash2  (type [20]byte) to type string */
	fmt.Printf("%s %x\n\n", "SHA1 Hash for \""+lang+"\" is ", sha1Hash2) /* %x format verb is used to convert hash result to a hex string*/

	//3rd way to generate SHA1 hash
	sha1Hasher3 := sha1.New()
	io.WriteString(sha1Hasher3, lang)
	fmt.Printf("%s %x\n", "SHA1 Hash for \""+lang+"\" is ", sha1Hasher3.Sum(nil)) /* %x format verb is used to convert hash result to a hex string*/

}

/*
SHA1 Hash for "Golang" is  92f4645da18d229c1e99797e4ee0203092e4c3ed

[146 244 100 93 161 141 34 156 30 153 121 126 78 224 32 48 146 228 195 237]
SHA1 Hash for "Golang" is  92f4645da18d229c1e99797e4ee0203092e4c3ed

SHA1 Hash for "Golang" is  92f4645da18d229c1e99797e4ee0203092e4c3ed
*/
