package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func home(rw http.ResponseWriter, req *http.Request) {
	//fmt.Println("Rendering ./hello.html")
	//ttp.ServeFile(rw, req, "./hello.html")
	//rw.Write([]byte("Hello"))
	http.Redirect(rw, req, "/v1/account/login/", 301)
}
func login(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Rendering ./login.html")
	http.ServeFile(rw, req, "./login.html")
}
func showBookDetails(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")

	queryParams := mux.Vars(req)
	bookName := queryParams["book_name"]
	bookName = strings.Replace(bookName, "_", " ", -1)

	fmt.Println("You want to see the details of '", bookName, "' ...Ok...wait...")

	db, err := sql.Open("mysql", "root:admin@321@tcp(127.0.0.1:3306)/library")
	if err != nil {
		fmt.Println("Error in connection with db")
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error in connection with db...connection is unavailable")
		return
	}
	fmt.Println("Connected...")
	q := "select * from books where name='" + bookName + "';"
	fmt.Println("Executing : ", q)
	var rows *sql.Rows
	rows, err = db.Query(q)
	fmt.Println("Data fetched from database")
	if err != nil {
		fmt.Println("Error in execution of mysql query...")
		return
	}

	startHtmlText := "<!doctype html><head><style>h1,td{color:white;}td{padding:5px;color:white;border:1px solid lightgreen;font-family:tahoma;font-size:18px;}</style></head><body bgcolor='navy'><center><h1>" + bookName + "</h1>"
	middleHtmlText := ""
	lastHtmlText := "</center></body>"
	var id int
	var name, author, edition, publication, availability, typ, category string //type is predefined one...so typ is here
	var price float64
	//found := false
	for rows.Next() {
		rows.Scan(&id, &name, &price, &author, &edition, &publication, &availability, &typ, &category)
		middleHtmlText += "<h2><table style='border-collapse:collapse;border:1px solid white;'>" +
			"<tr><td>Book Id</td><td>" + strconv.Itoa(id) + "</td></tr>" +
			"<tr><td>Book Name</td><td>" + name + "</td></tr>" +
			"<tr><td>Author</td><td>" + author + "</td></tr>" +
			"<tr><td>Author</td><td>" + strconv.FormatFloat(price, 'f', 2, 64) + "</td></tr>" +
			"<tr><td>Edition</td><td>" + edition + "</td></tr>" +
			"<tr><td>Publication</td><td>" + publication + "</td></tr>" +
			"<tr><td>Availability</td><td>" + availability + "</td></tr>" +
			"<tr><td>Type</td><td>" + typ + "</td></tr>" +
			"<tr><td>Category</td><td>" + category + "</td></tr>"
	}
	middleHtmlText += "</table><br><a href='/v1/library/books/programming' style='color:orange;'>Back</a>"
	htmlText := startHtmlText + middleHtmlText + lastHtmlText
	rw.Write([]byte(htmlText))
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
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error in connection with db...connection is unavailable")
		return
	}
	fmt.Println("Connected...")
	q := "select id,name from books where category='" + category + "';"
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
	var id int //It will be used to attach in new url
	htmlTmp = "<!doctype html><head><style type='text/css'>a:hover{color:white;}</style></head><body bgcolor='navy'><center><h1 style='color:white;'>" + category + " books</h1><br>"
	for rows.Next() {
		rows.Scan(&id, &book)
		newNameForBook := strings.Replace(book, " ", "_", -1)

		//booksName += "<a style='color:lightgreen'>" + book + "</h3>"
		booksName += "<a href='/v1/library/books/programming/" + newNameForBook + "/details' style='color:lightgreen;font-size:20px;font-family:open sans;text-decoration:none'>" + book + "</a><br>"
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
	rw.Write([]byte(htmlTmp + booksName + "<br><br><a href='/v1/library/books/admin/console' style='font-size:20px;color:orange;'>Back to Admin console</a></center></body>"))
}

func addCategory(rw http.ResponseWriter, req *http.Request) { //DB store
	fmt.Println("Rendering ./add_category.html")
	//http.ServeFile(rw, req, "./add_books.html")
}

func showAdminConsole(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Rendering ./admin_console.html")
	http.ServeFile(rw, req, "./admin_console.html")
}

func fillBookDetails(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Rendering ./add_books.html")
	http.ServeFile(rw, req, "./add_books.html")
}

func fillBookCategory(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Rendering ./add_category.html")
	http.ServeFile(rw, req, "./add_category.html")
}
func addNewBooks(rw http.ResponseWriter, req *http.Request) { //DB Store
	// fmt.Println("Rendering ./admin_console.html")
	// http.ServeFile(rw, req, "./admin_console.html")

	rw.Header().Set("Content-Type", "text/html; charset=utf-8")

	if req.Method != "POST" { //Checking whether the data is being sent using POST method or not(If already implemented in Android then not required)
		rw.Write([]byte("<h1 style='color:red'>You have to use POST method for sending data from front end</h1>"))
		fmt.Println("The data is not being sent using POST method")
		return
	}

	req.ParseForm()

	if len(req.Form) == 0 { //If there's no post data then inform the user
		rw.Write([]byte("<h1 style='color:red'>You have not filled any details about book</h1>"))
		fmt.Println("You haven't send the POST data")
		return
	}
	_, b := req.Form["book_name"]
	if b == false {
		fmt.Println("Golden ...")
		return
	} else {
		fmt.Println(b, "Golden Juise")
	}
	fmt.Println(b)
	if len(req.Form["book_name"]) == 0 {
		fmt.Println("name blank")
		return
	}
	bookName := strings.TrimSpace(req.Form["book_name"][0]) //username(optional)
	fmt.Println(bookName)

	price := strings.TrimSpace(req.Form["price"][0]) //p_pic(required), ltype, atype & cmtyid is not required for login, I need to check
	fmt.Println(price)

	author := strings.TrimSpace(req.Form["author"][0]) //From google  (optional)
	fmt.Println(author)

	edition := strings.TrimSpace(req.Form["edition"][0]) //community id (required)
	fmt.Println(edition)
	publication := strings.TrimSpace(req.Form["publication"][0]) //1 for google (required)
	fmt.Println(publication)

	availability := strings.TrimSpace(req.Form["availability"][0]) //Student, Job Seeker, Student, Employed (required)
	fmt.Println(availability)

	typ := strings.TrimSpace(req.Form["type"][0]) //app sign in key while first login (required)
	fmt.Println(typ)

	category := strings.TrimSpace(req.Form["category"][0])
	fmt.Println(category)

	if !areAllFieldsAreInValidForm(bookName, publication, price, author, availability, edition, typ, category) {
		rw.Write([]byte("</h1 style='color:red'>All the fields are required</h1>"))
		return
	}
	db, err := sql.Open("mysql", "root:admin@321@tcp(127.0.0.1:3306)/library")
	if err != nil {
		fmt.Println("Error in connection with db")
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error in connection with db...connection is unavailable")
		return
	}
	fmt.Println("Connected...")

	fmt.Println("Entered details : ", bookName, price, author, edition, publication, availability, typ, category)
	q := "insert into books (name, price, author, edition , publication,availability,type, category) values('" +
		bookName + "', " + price + ", '" + author + "' , '" + edition + "' , '" + publication + "' , '" +
		availability + "' , '" + typ + "' , '" + category + "');"
	stmt, err := db.Prepare(q)
	if err != nil {
		fmt.Println("Error in execution of query")
		return
	}
	stmt.Exec()
	fmt.Println("Book details successfully inserted")
	rw.Write([]byte("<h2 style='color:green'>Book details successfully inserted</h2>"))
}

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

func IsEmailValid(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}

func main() {
	port := "8002"
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/v1/account/login", login)
	router.HandleFunc("/v1/library/books/{category}", showBooks)
	router.HandleFunc("/v1/library/books/{category}/{book_name}/details", showBookDetails)
	router.HandleFunc("/v1/library/books/fill/details", fillBookDetails)
	router.HandleFunc("/v1/library/books/fill/categroy", fillBookCategory)
	router.HandleFunc("/v1/library/books/add/new", addNewBooks)
	router.HandleFunc("/v1/library/books/add/category", addCategory)
	router.HandleFunc("/v1/library/books/admin/console", showAdminConsole)
	//router.Handle("/hello", http.FileServer(http.Dir("./")))

	fmt.Println("Server is listening on the port ", port)

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		fmt.Println("Error in connection with the specified port : ", port)
		fmt.Println("After...Bye")
		return
	}
}
