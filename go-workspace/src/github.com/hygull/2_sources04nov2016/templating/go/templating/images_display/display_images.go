package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
)

func ShowImages(ctx *iris.Context) {
	/*****/
	ctx.SetContentType("text/html")
	fmt.Println("Connecting to DB...")
	db, err := sql.Open("mysql", "root:admin@321@tcp(127.0.0.1:3306)/media")

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

	queryString := "select title, url from images;"

	fmt.Println(queryString)
	rows, err := db.Query(queryString)
	if err != nil {
		fmt.Println("Error in query execution")
		return
	}
	found := false
	htmlText := "<!doctype html><body><center>"
	for rows.Next() {
		var title, url string

		rows.Scan(&title, &url)
		htmlText += "<h3 style='color:green'>" + title + "</h3>" +
			"<img  src='" + url + "' height='10%' width='10%'/><br>"
		fmt.Println(htmlText)
		found = true
	}
	htmlText += "</center></body>"

	if !found {
		fmt.Println("No posts found...")
	}

	ctx.HTML(iris.StatusOK, htmlText)
}

func ShowImagesAsJSON(ctx *iris.Context) {
	/*****/
	ctx.SetContentType("text/html")
	fmt.Println("Connecting to DB...")
	db, err := sql.Open("mysql", "root:admin@321@tcp(127.0.0.1:3306)/media")

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

	queryString := "select title, url from images;"

	fmt.Println(queryString)
	rows, err := db.Query(queryString)
	if err != nil {
		fmt.Println("Error in query execution")
		return
	}
	found := false

	type Image struct {
		Title string `json:"title"`
		Url   string `json:"url"`
	}
	var imagesArr []Image
	//imagesMap := map[string]interface{}{}
	for rows.Next() {
		var title, url string
		rows.Scan(&title, &url)
		imagesArr = append(imagesArr, Image{title, url})
		found = true
	}

	if !found {
		fmt.Println("No posts found...")
	}
	//imagesMap["images"] = imagesArr
	ctx.JSON(200, iris.Map{"iamges": imagesArr})
}

func main() {
	iris.Get("/v1/images/all", ShowImages)
	iris.Get("/v1/images/as_json", ShowImagesAsJSON)

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
