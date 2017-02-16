package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

//var golangSession *sessions.CookieStore
var golangSession = sessions.NewCookieStore([]byte("TheVerySecretInformation"))

func home(rw http.ResponseWriter, req *http.Request) {

	session, err := golangSession.Get(req, "golangers")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Extracted cookie value : ", session.Values["name"], session.Values["Content"])
	sessionBytes, _ := json.Marshal(*session)
	fmt.Printf("Typeof (*session) : %T", *session)
	fmt.Println(*session, err)
	fmt.Println("Converted data", string(sessionBytes))
	rw.Write([]byte("Maps"))
}
func logout(rw http.ResponseWriter, req *http.Request) {
	session, err := golangSession.Get(req, "golangers")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Options = &sessions.Options{
		Path: "/v1/hygull",
		//Domain :""
		// MaxAge=0 means no 'Max-Age' attribute specified.
		// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'.
		// MaxAge>0 means Max-Age attribute present and given in seconds.
		MaxAge:   -1,
		Secure:   true,
		HttpOnly: true,
	}
	fmt.Println("Session deleted...")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":9000", nil)
}
