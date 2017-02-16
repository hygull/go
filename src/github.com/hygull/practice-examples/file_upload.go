package controllers

import (
	"bufio"
	_ "database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)
import (
	"synkku/conf"
	"synkku/views"
)

/*********** Uploading images to Amazon S3 server *****************/
func UploadFileToAmazonS3(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")

	if req.Method != "POST" { //Checking whether the data is being sent using POST method or not(If already implemented in Android then not required)
		views.ShowSuccessOrErrorAsJSON(rw, "PostMethodNotFoundError", "The data is not being sent using POST method", 105, 5)
		fmt.Println("The data/image is not being sent using POST method")
		return
	}

	req.ParseForm()
	//fmt.Printf("len(req.Form) : %d %T\n",len(req.Form),req.Form);
	// if len(req.Form) == 0 { //If there's no post data then inform the user
	// 	views.ShowSuccessOrErrorAsJSON(rw, "NoPostDataError", "You haven't sent the POST data", 104, 4)
	// 	fmt.Println("You haven't send the POST data")
	// 	return
	// }
	//fileToBeUploaded := "./images/computer tablet phone_0.jpg"
	//uploadHandler(rw,req)
	/*---------------------------------------------------------------------*/
	supportedFileFormats := []string{"jpg", "png", "gif", "bmp", "jpeg", "PNG", "JPG", "JPEG", "GIF", "BMP", "mp4", "avi"}
	file, header, err := req.FormFile("img")

	defer file.Close()

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		fmt.Println("GOPATH not found")
		return
	}
	fmt.Println("Great")
	fmt.Println(reflect.TypeOf(header.Filename), header.Filename)
	fmt.Println("Work")

	re := regexp.MustCompile("^[a-zA-Z0-9_.]*$")
	if !re.MatchString(header.Filename) {
		fmt.Println("The file name should be alphanumeric, only _ and . are permitted as special characters")
		return
	}

	two := strings.Split(header.Filename, ".")
	if len(two) < 2 {
		fmt.Println("File name is not in proper format")
		return
	}
	IsFileFormatOk := false
	for _, extension := range supportedFileFormats {
		if extension == two[len(two)-1] {
			IsFileFormatOk = true
			break
		}
	}
	if !IsFileFormatOk {
		MsgMap := make(map[string]string)
		fmt.Println(`Only "jpg", "png", "gif", "bmp","jpeg","PNG","JPG","JPEG","GIF","BMP","mp4","avi" files are allowed`)
		MsgMap["Message"] = "Only jpg,jpeg,png,gif,bmp,mp4,avi files are allowed"
		MsgStr, _ := json.Marshal(MsgMap)
		fmt.Fprintf(rw, string(MsgStr))
		return
	}

	str := ""
	for _, part := range two[:len(two)-1] {
		str += part
	}
	if str == "" {
		fmt.Println("There should be a valid file name")
		return
	}

	s := strings.Join(strings.Fields(time.Now().String()[0:19]), ":")
	chars := []string{" ", "-", ":"}
	for _, char := range chars {
		s = strings.Replace(s, char, "_", -1)
	}

	s = str + "_" + s + "__synkku."
	newFileName := s + two[len(two)-1] //strings.Join(strings.Fields(time.Now().String()[0:19])
	i := 1
	fmt.Println("File name decided...")
	_, err = os.Stat(gopath + "/src/tmp/uploaded/" + newFileName)
	for !os.IsNotExist(err) { //Does not exist
		s = newFileName
		newFileName += strconv.Itoa(i)
		i += 1
		_, err = os.Stat(gopath + "/src/tmp/uploaded/" + newFileName)
	}

	out, err2 := os.Create(gopath + "/src/tmp/uploaded/" + newFileName)
	if err2 != nil {
		fmt.Fprintf(rw, "Unable to create the file for writing. Check your write access privilege")
		return
	}

	defer out.Close()

	// write the content from POST to the file
	_, err2 = io.Copy(out, file)
	if err2 != nil {
		fmt.Fprintln(rw, err)
	}

	fmt.Println("File locally uploaded successfully : ")
	//fmt.Fprintf(w, header.Filename)
	/*---------------------------------------------------------------------*/
	// if !AreAllFieldsAreInValidForm(fileToBeUploaded) {
	// 	views.ShowSuccessOrErrorAsJSON(rw, "EmptyFieldError", "You have to specify the name of image (Make sure the image should be inside ./images/)", 103, 3)
	// 	return
	// }
	fmt.Println("Starting to upload media to AMAZON S3")
	AWSAuth := aws.Auth{
		AccessKey: "AKIAI4RV6AOJMQQE5UJQ",                     //"AKIAIGV32QO6ZUEXFLNA",   //change this to yours
		SecretKey: "8KOd3w63SL8BWXoC7N5cmR/HyJ9zq9WBG+gvZJCg", //"3tMlCfZZb9hsPtP7vBFFPiCXQjU5vEXT74ip5yJ0",
	}
	region := aws.APNortheast
	connection := s3.New(AWSAuth, region)

	bucket := connection.Bucket("tiiing1") // change this your bucket name
	//"tiiing1.s3-website-ap-northeast-1.amazonaws.com/"
	//path := "tiiing1.s3-website-ap-northeast-1.amazonaws.com/" // this is the target file and location in S3
	//path := "images/nature/"
	path := "images/" + newFileName /*Very important part*/
	fmt.Println("Processing")
	//fileToBeUploaded := "./images/computer tablet phone_0.jpg"
	//fileToBeUploaded:=req.Form["img"];
	fmt.Println("AMAZON S3 Uploading in process")
	file2, err := os.Open(gopath + conf.PathToTheFolderContainingImages + "/" + newFileName)
	//ile, header, err = req.FormFile("img")

	defer file2.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//defer file.Close()

	// fileInfo, _ := file.Stat()
	// var size int64 = fileInfo.Size()
	// bytes := make([]byte, size)

	fileInfo, _ := file2.Stat()
	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)
	fmt.Println("Size of file calculated...", size, "bytes")
	// read into buffer
	buffer := bufio.NewReader(file2)
	_, err = buffer.Read(bytes)
	fmt.Println("Bytes read")
	filetype := http.DetectContentType(bytes)
	fmt.Println("Content type detected")
	err = bucket.Put(path, bytes, filetype, s3.ACL("public-read"))

	if err != nil {
		//fmt.Println(err)
		fmt.Println(`Error in uploading in AMAZON S3`)
		views.ShowSuccessOrErrorAsJSON(rw, "ImageUploadFailed", "image upload failed, fixed the issues and retry", 135, 35)
		//os.Exit(1)
	} else {
		fmt.Println("Media kept inside bucket")

		views.ShowSuccessOrErrorAsJSON(rw, "ImageUploadSuccessful", "The media "+header.Filename+" successfully uploaded to AMAZON S3 as "+newFileName, 125, 25)

	} //fmt.Printf("[Synkku] Your file successfully Uploaded to %s with %v bytes to S3\n\n", path, size)
}
