package main

import (
    "net/http"
    "log"
    "github.com/gorilla/mux"
    //"fmt"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Gorilla!\n"))
}

func main(){

		r := mux.NewRouter()
		//s := r.Host("{subdomain}.domain.com").Subrouter()
		// s.Path("/articles/{category}/{id:[0-9]+}").
		//   HandlerFunc(HelloHandler).
		//   Name("article")

		// // "http://news.domain.com/articles/technology/42"
		// url, err := r.Get("article").URL("subdomain", "news",
		//                                  "category", "technology",
		//                              "id", "42")
		// if err==nil{
		// 	fmt.Println("Serving : ",url) //url is of type url.URL
			r.HandleFunc("/",HelloHandler)
			log.Fatal(http.ListenAndServe(":8000", r))
			// } else{
			// 	fmt.Println("Error")
			// }   
		
}