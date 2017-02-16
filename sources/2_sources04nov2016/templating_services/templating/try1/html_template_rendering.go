package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
	"strconv"
)

func AdminLogin(ctx *iris.Context) {

	ctx.Render("admin_login.html", struct{ Name string }{Name: "iris"})
}

func DeactivatePost(ctx *iris.Context) {
	/*****/
	ctx.SetContentType("text/html")
	fmt.Println("Connecting to DB...")
	db, err := sql.Open("mysql", "root:admin@321@tcp(127.0.0.1:3306)/synkku")

	if err != nil {
		fmt.Println("Error in connection...")
		fmt.Println(err.Error())

		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error in connection test...")
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Successfully connected to Database...")
	pid := ctx.Param("post_id")
	if pid == "" {
		fmt.Println("Didn't get any post id in URL parameter to deactivate that...")
		return
	}
	fmt.Println("Got post id : ", pid, " from URL")
	queryString := "update posts set status=2 where post_id=" + pid

	fmt.Println(queryString)
	stmt, err := db.Prepare(queryString)
	if err != nil {
		fmt.Println("Error in query execution")
		return
	}
	stmt.Exec()
	fmt.Println("Post with post id : ", pid, " deactivated")
}

func ShowPosts(ctx *iris.Context) {
	/*****/
	ctx.SetContentType("text/html")
	fmt.Println("Connecting to DB...")
	db, err := sql.Open("mysql", "root:admin@321@tcp(127.0.0.1:3306)/synkku")

	if err != nil {
		fmt.Println("Error in connection...")
		fmt.Println(err.Error())

		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error in connection test...")
		fmt.Println(err.Error())
		return
	}

	queryString := "select CONCAT(users.first_name,last_name),posts.post_id," +
		"posts.text_message,posts.created_on, posts.updated_on from posts inner join users on " +
		"posts.user_id=users.user_id where (posts.status=1) limit 10;"

	fmt.Println(queryString)
	rows, err := db.Query(queryString)
	if err != nil {
		fmt.Println("Error in query execution")
		return
	}
	found := false
	htmlText := ""
	for rows.Next() {
		var un, txt, c_at, u_at string
		var pid int
		rows.Scan(&un, &pid, &txt, &c_at, &u_at)
		pId := strconv.Itoa(pid)
		htmlText + "" +
			"<h2 style='color:lightgreen'> Post Id : " + pId + "</h2>" +
			"<h5 style='color:white'>" + txt + "</h5>" +
			"<h5 style='color:yellow'>Posted by : " + un + "</h5>" +
			"<h5 style='color:lightblue'>Posted on : " + c_at + "</h5>" +
			"<h5 style='color:lightblue'>Updated on  : " + u_at + "</h5>" +
			"<a href='/v1/post/deactivate/" + pId + "'><button type='button' class='btn btn-success'>deactivate</button></a>" +
			"<hr></div><div class='col-md-1'></div></div>"
		fmt.Println(htmlText)
		found = true
	}
	if !found {
		fmt.Println("No posts found...")
	}
	/*****/
	fmt.Println("Rendering posts.html")
	//ctx.Render("posts.html", struct{ Posts string }{Posts: htmlText})
	//ctx.MustRender("posts.html", struct{ Posts string }{Posts: htmlText})
	fmt.Println("\n", htmlContentHead+htmlText+htmlContentTail)
	ctx.HTML(iris.StatusOK, htmlContentHead+htmlText+htmlContentTail)
}
func main() {
	iris.Get("v1/account/login", AdminLogin)
	iris.Get("v1/posts/all", ShowPosts)
	iris.Get("/v1/post/deactivate/:post_id", DeactivatePost)

	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.SetContentType("text/html")
		ctx.Write("<body bgcolor='black'><center><h1 style='color:lightgreen;'>SynkkU</h1><h2 style='padding:50px;color:white'>Page not found</h2></center></body>")
		ctx.Log("%s", "Page not found...404 Error")
	})

	iris.OnError(iris.StatusInternalServerError, func(ctx *iris.Context) {
		ctx.Write("<body bgcolor='black'><center><h1 style='padding:330px;color:white'>Internal server error</h1></center></body>")
		ctx.Log("%s", "Internal server error")
	})
	iris.Listen(":8080")
}
