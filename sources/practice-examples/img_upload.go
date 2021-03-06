
package main

import (
    _"bytes"
    "fmt"
    "io"
    _"io/ioutil"
    _"mime/multipart"
    "net/http"
    "os"
)

func postFile(w http.ResponseWriter,r * http.Request){
       r.ParseMultipartForm(32 << 20)
        file, handler, err := r.FormFile("uploadfile")
        if err != nil {
            fmt.Println(err)
            return
        }
        defer file.Close()
        fmt.Fprintf(w, "%v", handler.Header)
        f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
            fmt.Println(err)
            return
        }
        defer f.Close()
        io.Copy(f, file)
}

// sample usage
func main() {
    http.HandleFunc("/upload/",postFile)
    
    http.ListenAndServe(":9000",nil) 
}