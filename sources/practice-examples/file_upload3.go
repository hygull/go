 package main

 import (
 	"fmt"
 	"io"
 	"net/http"
 	"os"
 	_"time"
 	_"strings"
 	"strconv"
 )

 func uploadHandler(w http.ResponseWriter, r *http.Request) {

 	// the FormFile function takes in the POST input id file
 	file, header, err := r.FormFile("file")

 	if err != nil {
 		fmt.Fprintln(w, err)
 		return
 	}

 	defer file.Close();
 	gopath:=os.Getenv("GOPATH");
 	if gopath==""{
 		fmt.Println("GOPATH not found");
 		return
 	}
 	newFileName:="Media"//strings.Join(strings.Fields(time.Now().String()[0:19])
 	i:=1

 	_,err=os.Stat(gopath+"/src/tmp/uploaded/"+newFileName);
 	for !os.IsNotExist(err){//Does not exist
 		newFileName="Media"+strconv.Itoa(i)
		i+=1
		_,err=os.Stat(gopath+"/src/tmp/uploaded/"+newFileName)
	}

	out, err2 := os.Create(gopath+"/src/tmp/uploaded/"+newFileName)
	if err2 != nil {
		 		fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
		 		return
	}

 	defer out.Close()

 	// write the content from POST to the file
 	_, err2 = io.Copy(out, file)
 	if err2 != nil {
 		fmt.Fprintln(w, err)
 	}

 	fmt.Fprintf(w, "File uploaded successfully : ")
 	fmt.Fprintf(w, header.Filename)
 }

 func main() {
 	http.HandleFunc("/", uploadHandler)
 	http.ListenAndServe(":8080", nil)
 }