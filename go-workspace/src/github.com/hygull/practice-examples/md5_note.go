package main 
import ("fmt";"crypto/md5";"strconv";"time";"encoding/hex")
func getMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}

func isAppSignInKeyCorrect(appSignInKeyMd5 string,authenticationToken string,email string) bool{
	reversed_email:=[]byte(email); //string is immutable in Go
	//email reversing for loop
	for front,last:=0,len(reversed_email)-1;front<last;front,last=front+1,last-1{
		reversed_email[front],reversed_email[last]=reversed_email[last],reversed_email[front];
	}
	reversed_email_and_auth_token:=string(reversed_email)+authenticationToken;

	fmt.Printf("%v (%T)",reversed_email_and_auth_token,reversed_email_and_auth_token)
	fmt.Println()
	fmt.Println(appSignInKeyMd5)

	/*md5_of_reversed_email_and_auth_token:=md5.Sum([]byte(reversed_email_and_auth_token))*/
	md5_of_reversed_email_and_auth_token:=getMD5Hash(reversed_email_and_auth_token)
	fmt.Println(md5_of_reversed_email_and_auth_token)

	/*fmt.Println()
	fmt.Printf("%T : %T",md5_of_reversed_email_and_auth_token,reversed_email)
	fmt.Println()
	fmt.Printf("%v %T",md5_of_reversed_email_and_auth_token,md5_of_reversed_email_and_auth_token)
	fmt.Println()
	fmt.Printf("%v %T",md5_of_reversed_email_and_auth_token[:],md5_of_reversed_email_and_auth_token[:])
	fmt.Println(string(md5_of_reversed_email_and_auth_token[:]))
	*/
	return md5_of_reversed_email_and_auth_token==appSignInKeyMd5
}

func getStoken(userId int, authenticationToken string) string{
	unixTime:=strconv.FormatInt(time.Now().UTC().Unix(), 10)
	tempStoken:= strconv.Itoa(userId)+authenticationToken+unixTime //<userId><authenticationToken><unixTime>
	stoken:=getMD5Hash(tempStoken)			 //md5("<userId><authenticationToken><unixTime>")
	return stoken
}

func main(){
	AuthenticationKey:="eZ$21#@54>4074W8Ndkf**WE32a"//we2376THWEKm"

	var email , app_sign_in_key string
	var userId int

	fmt.Printf("%-30s","Enter your user id : ")
	fmt.Scan(&userId)
	fmt.Printf("%-30s","Enter email        : ");
	fmt.Scan(&email)
	fmt.Printf("%-30s","App signin key     : ");
	fmt.Scan(&app_sign_in_key)

	fmt.Println("\n*********************** Token Test **********************\n")
	if !isAppSignInKeyCorrect(app_sign_in_key,AuthenticationKey,email){ //Call to a Middleware
		fmt.Println("The token is invalid")
		return
	}else{
		fmt.Println("\n\nToken is correct");
		fmt.Println("stoken : " ,getStoken(userId,AuthenticationKey))
		fmt.Println()
	}
}