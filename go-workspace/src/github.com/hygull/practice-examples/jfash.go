/*
 # Start date     : 21/09/2016, Wednesday
 # Aim of program : To create a REST API for an app of Android devices
 # Coded by       : Rishikesh Agrawani
 # Guided by      : Rathnakara Sir
 # Note           : Lists of access urls are in a file named access_urls_new.txt
*/

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"regexp"
)

/*********************Storing new post details to database **************/
func sendEmailAndMessage(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")

	if req.Method != "POST" { //Checking whether the data is being sent using POST method or not(If already implemented in Android then not required)
		//showSuccessOrErrorAsJSON(rw, "PostMethodNotFoundError", "The data is not being sent using POST method", 105, 5)
		fmt.Println("The data is not being sent using POST method")
		return
	}

	req.ParseForm()
	//fmt.Printf("len(req.Form) : %d %T\n",len(req.Form),req.Form);
	if len(req.Form) == 0 { //If there's no post data then inform the user
		//showSuccessOrErrorAsJSON(rw, "NoPostDataError", "You haven't sent the POST data", 104, 4)
		fmt.Println("You haven't send the POST data")
		return
	}

	email := req.Form["email"][0]
	name := req.Form["name"][0]
	phone := req.Form["phone"][0]
	message := req.Form["message"][0]

	if !areAllFieldsAreInValidForm(email, message, phone, name) {
		htmlMsg := "<h1 style='color:red'>All the fields are mandatory</h1>"
		//fmt.Fprintf(rw, htmlMsg)
		rw.Write([]byte(htmlMsg))
		return
	}

	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !Re.MatchString(email) {
		fmt.Println("This is not a valid email, enter a proper email id.")
		htmlMsg := "<h1 style='color:red'>Email is not valid</h1>"
		fmt.Fprintf(rw, htmlMsg)
		return
	}
	//If all the data are not blank
	fmt.Println("Entered details for sending details to JEEVI: ", email, name, phone, message)
	fmt.Fprintf(rw, "<h1 style='color:green'>Hello</h1>")
} //End of function storeNewPostToDb

/******************Checking for empty values (Using variadic functions)**/
func areAllFieldsAreInValidForm(postData ...string) bool {
	isAnyBlank := false
	for _, data := range postData {
		if data == "" {
			isAnyBlank = true
			break
		}
	}
	fmt.Println(!isAnyBlank)
	return !isAnyBlank //An efficient way to check for empty fields
}

/******************Routes & Handlers call********************************/
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	//HOME PAGE(Temporary)
	myRouter.HandleFunc("/v1/send/email", sendEmailAndMessage)

	log.Fatal(http.ListenAndServe(":9000", myRouter))
}

/***************DEFINITION OF STARTER OF THIS REST-API PROGRAM************/
func main() {
	fmt.Println("\n**************** WELCOME TO ****************") //Console message to show the
	fmt.Println("******************  FASHION *****************")  //To show the status of working API
	fmt.Print("\nJEEVI's SERVER is running on  127.0.0.1:9000\n\n")
	handleRequests()
}

/*************************************END*********************************/
