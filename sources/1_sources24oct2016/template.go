package main

import "html/template"
import "net/http"
import "os"

//import "github.com/justinas/nosurf"
import "fmt"

func main() {
	http.HandleFunc("/v1", home)
	http.ListenAndServe(":8000", nil)
}

var html = `
			<!doctype html>
			<html>
			<head>
			<title>{{.title}}</title>
			<style>
				h1{
					color:lightgreen;
				}
				h2{
					color:white;
				}
			</style>
			</head>
			<body bgcolor='navy'>
			<center>
			<h1>{{ .article_heading }}</h1>
			<h2>{{ .article_content }}</h2>
			<img src={{.image_link}} height="50%" width="50%" alt="IMAGE_SHOULD_BE_THERE">
			</center>
			</body>
			</html>
		  `
var templ = template.Must(template.New("t1").Parse(html))

func home(rw http.ResponseWriter, req *http.Request) {
	context := make(map[string]string)
	if len(os.Args) == 5 {
		context["title"] = os.Args[1]
		context["article_heading"] = os.Args[2]
		context["article_content"] = os.Args[3]
		context["image_link"] = os.Args[4]
	} else {
		if len(os.Args) == 1 {
			context["title"] = "Go"
			context["article_heading"] = "Maps in Go"
			context["article_content"] = "Map is a key value based data structure where keys are mapped to some particular values or some other data structure"
			context["image_link"] = "http://4.bp.blogspot.com/-yfcJ9QhZ6EE/VZpKaZDHgaI/AAAAAAAADqw/k4pC1NSgXrY/s1600/railway-track-1920x965.jpg"
		} else {
			fmt.Println("(1) Either you have to provide TITLE OF HTML, HEADING OF ARTICLE, CONTENT OF ARTICLE\n")
			fmt.Println("(2) Or you have not to provide any as default will be displayed in this case")
			return
		}
	}
	templ.Execute(rw, context)
}
