package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

var golangSession = sessions.NewCookieStore([]byte("TheVerySecretInformation"))

func home(rw http.ResponseWriter, req *http.Request) {
	session, err := golangSession.Get(req, "golangers")
	fmt.Printf("Typeof(golangSession) %T", golangSession)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["name"] = "golangers"
	session.Values["Content"] = "Dobian hill"
	session.Values["age"] = 24
	session.Save(req, rw)
	http.Redirect(rw, req, "/v1/goal_http", 301)
}
func about(rw http.ResponseWriter, req *http.Request) {
	session, err := golangSession.Get(req, "golangers")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("\nExtracted cookie value : ", session.Values["name"], session.Values["Content"])
	sessionBytes, _ := json.Marshal(*session)
	fmt.Printf("Typeof (*session) : %T", *session)
	fmt.Println(*session, err)
	fmt.Println("Converted data", string(sessionBytes))
	rw.Write([]byte("Maps"))
}
func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/v1/goal_http", about)
	http.ListenAndServe(":8000", nil)
}
