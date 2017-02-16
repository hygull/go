package main

import (
	"fmt"
	"github.com/kataras/iris"
)

var Topics = func(ctx *iris.Context) {
	SetContentTypeAndShowPath(ctx)
	ctx.Write("<center><h1 style='color:green;'>Topics</h1>" +
		"<h2 style='color:blue;'>Beginners topics</h2>" +
		"<h2 style='color:blue;'>Advanced topics</h2></center>")
}

func SetContentTypeAndShowPath(ctx *iris.Context) {
	ctx.SetContentType("text/html")
	fmt.Println("Visiting : ", ctx.PathString())
}

func GolangResources(ctx *iris.Context) {
	SetContentTypeAndShowPath(ctx)
	ctx.Write("<center><h1 style='color:green;'>Resourcces on web</h1>" +
		"<a href='https://github.com/kataras/iris'>Great iris documentation on github</a><br>" +
		"<a href='https://docs.iris-go.com'>docs..iris-go.com</a><br></center>")
}

func GetXwwwFormUrlEncodedData(ctx *iris.Context) {
	SetContentTypeAndShowPath(ctx)
	name := ctx.FormValueString("name")
	lang := ctx.FormValueString("lang")
	ctx.Write("<h1 style='color:green'>%s</h1><h1 style='color:green'>%s</h1>", name, lang)
}

func showMap(ctx *iris.Context) {
	ctx.SetContentType("application/json")
	type User struct {
		Name   string
		Age    int
		Branch string
		Degree string
	}
	userMap := iris.Map{}
	userMap["Name"] = "Rishikesh"
	userMap["Age"] = 24
	userMap["Branch"] = "CSE"
	userMap["Degree"] = "BTech"
	fmt.Println()
}
func main() {

	iris.Get("/", func(ctx *iris.Context) {
		SetContentTypeAndShowPath(ctx)
		ctx.Write("<center><h1 style='color:green;'>Golangers</h1></center>")
	})

	golang := func(ctx *iris.Context) {
		SetContentTypeAndShowPath(ctx)
		ctx.Write("<center><h1 style='color:green;'>Lets learn Golang<h1></center>")
	}

	call := func(ctx *iris.Context) {
		ctx.SetContentType("application/json")
		fmt.Println("Visiting ", ctx.PathString())
		ctx.JSONP(iris.StatusOK, "User", struct{ Name string }{"Rishikesh"})
	}

	user := func(ctx *iris.Context) {
		ctx.SetContentType("application/json")
		fmt.Println("Visiting ", ctx.PathString())
		ctx.JSONP(iris.StatusOK, "User", iris.Map{"name": "Rishikesh Agrawani", "branch": "CSE", "interested_in": "Programming, reading tutorials"})
	}

	user2 := func(ctx *iris.Context) {
		ctx.SetContentType("application/json")
		fmt.Println("Visiting ", ctx.PathString())
		ctx.JSON(iris.StatusOK, iris.Map{"name": "Rishikesh Agrawani", "branch": "CSE", "interested_in": "Programming, reading tutorials"})
	}

	iris.Get("/golang/topics", Topics)
	iris.Get("/golang", golang)
	iris.Get("/call", call)
	iris.Get("/user", user)
	iris.Get("/user2", user2)
	iris.HandleFunc("GET", "/golang/resources", GolangResources)
	iris.HandleFunc("POST", "/post/get", GetXwwwFormUrlEncodedData)

	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.SetContentType("text/html")
		ctx.Write("<body bgcolor='black'><center><h1 style='padding:330px;color:white'>Page not found</h1></center></body>")
		ctx.Log("%s", "Page not found...404 Error")
	})

	iris.OnError(iris.StatusInternalServerError, func(ctx *iris.Context) {
		ctx.Write("<body bgcolor='black'><center><h1 style='padding:330px;color:white'>Internal server error</h1></center></body>")
		ctx.Log("%s", "Internal server error")
	})
	fmt.Println("Server is running on port 8080")
	iris.Listen(":8080")
}
