package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

func home(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Rendering ./hello.html")
	//ttp.ServeFile(rw, req, "./hello.html")
	rw.Write([]byte("Hello"))
}
func login(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Rendering ./login.html")
	http.ServeFile(rw, req, "./login.html")
}

func showBooks(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")

	queryParams := mux.Vars(req)
	category := queryParams["category"]
	fmt.Println("You want to see the ", category, " books...Ok...wait...")

	db, err := sql.Open("mysql", "root:admin@321@tcp(127.0.0.1:3306)/library")
	if err != nil {
		fmt.Println("Error in connection with db")
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error in connection with db...connection is unavailable")
		return
	}
	fmt.Println("Connected...")
	q := "select name from books where category='" + category + "';"
	fmt.Println("Executing : ", q)
	var rows *sql.Rows
	rows, err = db.Query(q)
	fmt.Println("Data fetched from database")
	if err != nil {
		fmt.Println("Error in execution of mysql query...")
		return
	}
	fmt.Printf("%T\n", rows)
	var book, booksName, htmlTmp string
	htmlTmp = "<!doctype html><body bgcolor='navy'><center><h1 style='color:white'>" + category + " books</h1><br>"
	for rows.Next() {
		rows.Scan(&book)
		booksName += "<h3 style='color:lightgreen'>" + book + "</h3>"
		fmt.Println(book)
	}

	fmt.Println("Details successfully extracted...")

	// 	htmlText := `<!DOCTYPE html>
	// <html>
	// <head>
	//   <meta name='viewport' content='width=device-width, initial-scale=1'>
	//   <link rel='stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css'>
	//   <script src='https://ajax.googleapis.com/ajax/libs/jquery/1.12.4/jquery.min.js'></script>
	//   <script src='https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js'></script>
	// </head>
	// <style>
	//   h1,h3{
	//     color:white;
	//   }
	//   .container .dropdown ul li a{
	//     color:green;
	//   }
	// </style>
	// <body>

	// <div class='container'>
	//   <center><h1> ` + category + `books </h1><br>` +
	// 		booksName +
	// 		`</center></div></body></html>`
	//rw.Write([]byte(`<h1>Hello Bro</h1>` + "<h2>Great</h2>"))
	rw.Write([]byte(htmlTmp + booksName + "</center></body>"))
}

func main() {
	port := "8001"
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/account/login", login)
	router.HandleFunc("/v1/library/books/{category}", showBooks)
	router.Handle("/hello", http.FileServer(http.Dir("./")))

	fmt.Println("Server is listening on the port ", port)

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		fmt.Println("Error in connection with the specified port : ", port)
		fmt.Println("After...Bye")
		return
	}
}
