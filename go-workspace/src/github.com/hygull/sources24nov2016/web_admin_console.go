package main

import "io/ioutil"
import "fmt"
import "net/http"
import "strings"
import "github.com/gorilla/mux"
import "regexp"
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func home(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Redirecting to login...")
	http.Redirect(rw, req, "/high_goal/v1/account/login", 301)
}

func login(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Welcome to login.html")
	htmlBytes, err := ioutil.ReadFile("./html_css/login.html")
	if err != nil {
		fmt.Println("Can't read html file")
		return
	}
	rw.Write(htmlBytes)
}

func credentialsValidator(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Println("Connecting to mysql...")
	db, err := sql.Open("mysql", "root:admin@321@tcp(127.0.0.1:3306)/hygull")
	if err != nil {
		text := "Error while establishing connection with mysql..."
		fmt.Println(text)
		rw.Write([]byte(text))
		return
	}
	err = db.Ping()
	if err != nil {
		text := "Connection test did not succeed"
		fmt.Println(text)
		rw.Write([]byte(text))
		return
	}
	fmt.Println("Connected...")

	req.ParseForm()
	email := strings.TrimSpace(req.Form["email"][0])
	fmt.Println("Read email field...")
	if email == "" {
		text := "Email can't be blank..."
		fmt.Println(text)
		rw.Write([]byte("<center><h1 style='color:red;'>" + text + "</h1><a href='/high_goal/v1/account/login'>Retry</a></center>"))
		return
	}
	if !IsEmailValid(email) {
		text := "Email format is not valid..."
		fmt.Println(text)
		rw.Write([]byte("<center><h1 style='color:red;'>" + text + "</h1><a href='/high_goal/v1/account/login'>Retry</a></center>"))
		return
	}
	var (
		password, remember string
	)
	if len(req.Form["password"]) != 0 {
		password = strings.TrimSpace(req.Form["password"][0])
	}
	fmt.Println("Read password field...")
	if len(req.Form["remember"]) != 0 {
		remember = strings.TrimSpace(req.Form["remember"][0])
	}

	fmt.Println("Read remeber field...")
	fmt.Println("Entered details : ", email, password, remember)
	if password != "" {
		if (len(password) >= 8) && (len(password) <= 20) {
			/*blank*/
		} else {
			text := "Password length should vary from 8 to 20..."
			text2 := "<h2 style='color:maroon'>Back to Hygull login</h2>"
			fmt.Println(text)
			rw.Write([]byte("<center><h1 style='color:red;'>" + text + "</h1>" + text2 + "<a href='/high_goal/v1/account/login'>Retry</a></center>"))
			return
		}
	} else {
		text := "Blank password field...It is required..."
		text2 := "<h2 style='color:maroon'>Back to Hygull login</h2>"
		fmt.Println(text)
		rw.Write([]byte("<center><h1 style='color:red;'>" + text + "</h1>" + text2 + "<a href='/high_goal/v1/account/login'>Retry</a></center>"))
		return
	}

	var rows *sql.Rows
	fmt.Println("Executing select query...")
	rows, err = db.Query("select user_name from auth_users where email='" + email + "' and password='" + password + "';")
	if err != nil {
		text := "Error while executing select query..."
		fmt.Println(text)
		rw.Write([]byte(text))
		return
	}
	fmt.Println("Select query successfully executed...")
	if rows.Next() {
		var uname string
		rows.Scan(&uname)

		link := "/high_goal/v1/" + uname + "/new/posts"
		fmt.Println("Data Scanned...")
		fmt.Println("Redirecting to ", link)
		http.Redirect(rw, req, link, 301)
	} else {
		text := "This email and password combination does not exist..."
		text2 := "<h2 style='color:maroon'>check your email and password...<h2>"
		fmt.Println(text)
		rw.Write([]byte("<body bgcolor='#d2b48c'><center><h1 style='color:red;'>" + text + "</h1>" + text2 + "<a href='/high_goal/v1/account/login'>Retry</a></center></body>"))
		return
	}
}

func IsEmailValid(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/high_goal/v1/account/login", login)
	router.HandleFunc("/high_goal/v1/validate/credentials", credentialsValidator)
	http.ListenAndServe(":9000", router)
}
