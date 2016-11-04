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

func postsDisplayerForLoggedInUser(rw http.ResponseWriter, req *http.Request) {

}

func newPostsDisplayerLimit10(rw http.ResponseWriter, req *http.Request) {

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

	var title string
	var message string
	var postedBy string
	var postedOn string

	var posts string = ""
	var rows *sql.Rows
	query := "select auth_users.user_name, posts.title, posts.message, posts.posted_on from auth_users inner join posts on posts.user_id=auth_users.id;"
	fmt.Println(query)
	rows, err = db.Query(query)
	if err != nil {
		text := "Error while executing select query..."
		fmt.Println(text)
		rw.Write([]byte(text))
		return
	}
	fmt.Println("Select query successfully executed...")

	found := false
	for rows.Next() {
		if found == false {
			found = true
		}
		rows.Scan(&postedBy, &title, &message, &postedOn)
		posts += "<div style='border:1px solid #eee8aa;border-top:15px solid #ffb4e5;'><h1>" +
			title + "<h1><hr>" + "<h2>" + message + "</h2>" + "<h4>Posted by : " + postedBy + "</h4>" +
			"<h4>Posted on : " + postedOn + "</h4>" +
			"</div>"
	}
	if found == false {
		text := "No more posts are there"
		text2 := "<h2 style='color:maroon'>0 posts found<h2>"
		fmt.Println(text)
		rw.Write([]byte("<body bgcolor='#d2b48c'><center><h1 style='color:red;'>" + text + "</h1>" + text2 + "<a href='/high_goal/v1/account/login'>Retry</a></center></body>"))
		return
	}

	fmt.Println("Data Scanned...")
	rw.Write([]byte("<head><style>h1{color:white;}h4{color:lightgreen;} h2{color:#fa8072;}</style></head><body bgcolor='navy' style='padding-left:250px;padding-right:250px;'><center><h1 style='color:#f5deb3;'>New posts</h1>" + posts + "</center></body>"))

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/high_goal/v1/account/login", login)
	router.HandleFunc("/high_goal/v1/validate/credentials", credentialsValidator)
	router.HandleFunc("/high_goal/v1/{username}/posts", postsDisplayerForLoggedInUser)
	router.HandleFunc("/high_goal/v1/{username}/new/posts", newPostsDisplayerLimit10)
	http.ListenAndServe(":9000", router)
}
