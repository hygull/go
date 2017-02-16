package main

import "fmt"
import "net/http"

func Home(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Println("Rendering ./hello.html")
	//http.ServeFile(rw, req, "./hello.html")
	rw.Write([]byte("<h1 style='color:blue'>Hello</h1>"))
}
func Login(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Rendering ./login.html")
	http.ServeFile(rw, req, "./login.html")
}
func main() {
	port := "9001"
	http.HandleFunc("/", Home)
	http.HandleFunc("/account/login/", Login)
	http.Handle("/hello", http.FileServer(http.Dir("./")))
	fmt.Println("Server is listening on the port ", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error in connection with the specified port : ", port)
		fmt.Println("After...Bye")
		return
	}
}
