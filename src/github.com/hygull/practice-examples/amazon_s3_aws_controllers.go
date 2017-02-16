package controllers

import (
	"bufio"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
	_ "log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

import (
	"synkku/conf"
	"synkku/views"
)

/*******************Global declaration **********************************/
var SynkkuLoggedInUsersId = conf.SynkkuLoggedInUsersId


/******************Checking for User Login in the current system ********/
func isUserLoggedIn() bool {
	//-----------------1st way using DB------

	/*db,e:=sql.Open("mysql","root:admin@321@tcp(127.0.0.1:3306)/tiiingsession")
		defer db.Close();
		if e!=nil{
			panic(e.Error())
		}
		rows,err:=db.Query("select user_id from system_loggedin_user;");
		if err!=nil{
			panic(err.Error())
		}
	    isAnyLoggedIn:=false
	    for rows.Next(){
	    	isAnyLoggedIn=true
	    	rows.Scan(&global.TiiingUserId);
	    }
	    if isAnyLoggedIn{
	    	fmt.Println("The current logged in user id is : ",global.TiiingUserId);
	    	return true
	    }else{
	    	return false
	    }*/

	//-----------------2nd way using global variable-------
	if SynkkuLoggedInUsersId == -1 {
		return false
	} else {
		return true
	}
}


/************************ Login (Temp)***************************/
func LoginIntoAccount(rw http.ResponseWriter, req *http.Request) {
	//Not checking for POST method
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")

	if req.Method != "POST" && req.Method != "PUT" { //Checking whether the data is being sent using POST method or not(If already implemented in Android then not required)
		views.ShowSuccessOrErrorAsJSON(rw, "PutOrPostMethodNotFoundError", "The data is not being sent using POST method", 107, 7)
		fmt.Println("The data is not being sent using POST method")
		return
	}
	if isUserLoggedIn() {
		views.ShowSuccessOrErrorAsJSON(rw, "AlreadyLoggedIn", "2 users are not allowed to access the same server's port, First Logout then Login", 122, 22)
		return
	}
	req.ParseForm()

	if len(req.Form) == 0 { //If there's no post data then inform the user
		views.ShowSuccessOrErrorAsJSON(rw, "NoPostDataError", "You haven't sent the POST data", 104, 4)
		fmt.Println("You haven't send the POST data")
		return
	}
	tempUserId := strings.TrimSpace(req.Form["us_id"][0])
	tempEmail := strings.TrimSpace(req.Form["eml"][0])
	//password is required, Implementation required

	if !areAllFieldsAreInValidForm(tempEmail, tempUserId) {
		views.ShowSuccessOrErrorAsJSON(rw, "EmptyFieldError", "You have to select all the fileds", 103, 3)
		return
	}
	fmt.Println("Connecting to DB...")
	//db, err := sql.Open("mysql", conf.DBUserName+":"+conf.DBPassword+"@tcp"+"("+conf.DBHost+":"+conf.DBPort+")/"+conf.DBName)
	db, err := sql.Open("mysql", conf.DBUserName+":"+conf.DBPassword+"@tcp"+"("+conf.DBHost+":"+conf.DBPort+")/"+conf.DBName)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("Connected")
	rows, err := db.Query("select  user_id, email_id from users where user_id=" + tempUserId + " AND email_id= '" + tempEmail + "';")
	if err != nil {
		panic(err.Error())
	}

	found := false
	for rows.Next() {
		found = true
	}
	if found {
		SynkkuLoggedInUsersId, _ = strconv.Atoi(tempUserId) //Global userId set to keep track about user login
		fmt.Println("Your Logged in User ID : ", SynkkuLoggedInUsersId)
		//views.ShowSuccessOrErrorAsJSON(rw, "LoginSuccessful", "You have successfully logged in", 121, 21) //No need to return from here
		http.Redirect(rw, req, "/", 301)
	} else {
		views.ShowSuccessOrErrorAsJSON(rw, "UserDoesNotExist", "This user does not exist, Try again", 101, 1)
		return
	}
}

/************************ Login *****************************************/
func LogoutFromAccount(rw http.ResponseWriter, req *http.Request) {
	//Not checking for POST method
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")

	if SynkkuLoggedInUsersId == -1 {
		views.ShowSuccessOrErrorAsJSON(rw, "LoginRequired", "You need to login to access the site. Visit http:127.0.0.1:9000/v1/feeds/login", 120, 20)
		return
	} else {
		SynkkuLoggedInUsersId = -1
		views.ShowSuccessOrErrorAsJSON(rw, "LoggedOut", "You successfully logged out", 122, 22)
	}
}



/******************Checking for empty values (Using variadic functions)**/
func areAllFieldsAreInValidForm(postData ...string) bool {
	isAnyBlank := false
	for _, data := range postData {
		if data == "" {
			isAnyBlank = true
			break
		}
	}
	fmt.Println(!isAnyBlank)
	return !isAnyBlank //An efficient way to check for empty fields
}


f
/********************************************************************************/
func DeactivateACityOffer(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	if !isUserLoggedIn() {
		views.ShowSuccessOrErrorAsJSON(rw, "LoginRequired", "You need to login to access the site. Visit http:127.0.0.1:9000/v1/feeds/login", 120, 20)
		return
	}

	if req.Method != "POST" { //Checking whether the data is being sent using POST method or not(If already implemented in Android then not required)
		views.ShowSuccessOrErrorAsJSON(rw, "PostMethodNotFoundError", "The data is not being sent using POST method", 105, 5)
		fmt.Println("The data is not being sent using POST method")
		return
	}

	req.ParseForm()
	//fmt.Printf("len(req.Form) : %d %T\n",len(req.Form),req.Form);
	if len(req.Form) == 0 { //If there's no post data then inform the user
		views.ShowSuccessOrErrorAsJSON(rw, "NoPostDataError", "You haven't sent the POST data", 104, 4)
		fmt.Println("You haven't send the POST data")
		return
	}

	userId := req.Form["us_id"][0]
	cityOfferId := req.Form["ctyofrid"][0]

	if !areAllFieldsAreInValidForm(userId, cityOfferId) {
		views.ShowSuccessOrErrorAsJSON(rw, "EmptyFieldError", "You have to select all the fileds", 103, 3)
		return
	}
	//If all the data are not blank
	fmt.Println("Got userId : " + userId + " and cityOfferId : " + cityOfferId)
	fmt.Println("Preparing to add post details to Database ")
	db, err := sql.Open("mysql", conf.DBUserName+":"+conf.DBPassword+"@tcp"+"("+conf.DBHost+":"+conf.DBPort+")/"+conf.DBName) //Table name
	fmt.Println("Connected...")
	if err != nil {
		fmt.Println("This is the error : ", err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Println("Ping Error : ", err.Error())
	} else {
		fmt.Println("Connection is available...")
	}

	rows, e := db.Query("select user_id from users where user_id=" + userId + ";")
	if e != nil {
		panic(e)
	} else {
		fmt.Println("userId successfully retreived from DB....")
	}

	found := false
	for rows.Next() {
		found = true
	}
	if found == false {
		views.ShowSuccessOrErrorAsJSON(rw, "UserDoesNotExist", "This user id does not exist in the database", 101, 1)
	} else {
		rows, e = db.Query("select status from users where user_id=" + userId + " AND id=" + cityOfferId + ";")
		if e != nil {
			panic(e)
		} else {
			fmt.Println("status successfully retreived from DB(cityoffers table)....")
		}
		found = false
		var status int
		updated := false
		for rows.Next() {
			found = true //Execute only once
			rows.Scan(&status)
		}

		if found == true {
			if status == 1 { //If status is 1 then we will change it to 2
				stmt, err2 := db.Prepare("update cityoffers set status = 2;")
				fmt.Println("status updated in cityoffers table")
				updated = true //If I will not use this variable then I need to write the same code here and the following status==2 if section
				if err2 == nil {
					stmt.Exec()
				}
			} else {
				if status == 2 { //If status is 2 then we will change it to 1
					stmt, err2 := db.Prepare("update cityoffers set status = 1;")
					fmt.Println("status updated in cityoffers table") //Code duplication here, change is required
					updated = true
					if err2 == nil {
						stmt.Exec()
					}
				} else { //If niether 1 nor 2 then we will an error message
					views.ShowSuccessOrErrorAsJSON(rw, "CityStatusFoundOutOfRange", "CityOffer status should be either 1 or 2, worng entry in DB, Update required", 117, 17)
					return //If return will not be here then It will allow the next if to execute
				}
			}
			if updated == true {
				views.ShowSuccessOrErrorAsJSON(rw, "CityOfferDeactivated", "City offer successfully de-activated", 118, 18)
			}
		} else {
			views.ShowSuccessOrErrorAsJSON(rw, "CityOfferStatusNotFound", "City status does not exist for this user and cityId", 116, 16)
		}
	}
}

/*********** Uploading images to Amazon S3 server *****************/
func UploadImageToAmazonS3(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	if !isUserLoggedIn() {
		views.ShowSuccessOrErrorAsJSON(rw, "LoginRequired", "You need to login to access the site. Visit http:127.0.0.1:9000/v1/feeds/login", 120, 20)
		return
	}

	if req.Method != "POST" { //Checking whether the data is being sent using POST method or not(If already implemented in Android then not required)
		views.ShowSuccessOrErrorAsJSON(rw, "PostMethodNotFoundError", "The data is not being sent using POST method", 105, 5)
		fmt.Println("The data/image is not being sent using POST method")
		return
	}

	req.ParseForm()
	//fmt.Printf("len(req.Form) : %d %T\n",len(req.Form),req.Form);
	if len(req.Form) == 0 { //If there's no post data then inform the user
		views.ShowSuccessOrErrorAsJSON(rw, "NoPostDataError", "You haven't sent the POST data", 104, 4)
		fmt.Println("You haven't send the POST data")
		return
	}
	//fileToBeUploaded := "./images/computer tablet phone_0.jpg"
	fileToBeUploaded := req.Form["img"][0]
	if !areAllFieldsAreInValidForm(fileToBeUploaded) {
		views.ShowSuccessOrErrorAsJSON(rw, "EmptyFieldError", "You have to specify the name of image (Make sure the image should be inside ./images/)", 103, 3)
		return
	}

	AWSAuth := aws.Auth{
		AccessKey: "AKIAI4RV6AOJMQQE5UJQ", // change this to yours
		SecretKey: "8KOd3w63SL8BWXoC7N5cmR/HyJ9zq9WBG+gvZJCg",
	}
	region := aws.APNortheast
	connection := s3.New(AWSAuth, region)

	bucket := connection.Bucket("tiiing1") // change this your bucket name
	//"tiiing1.s3-website-ap-northeast-1.amazonaws.com/"
	//path := "tiiing1.s3-website-ap-northeast-1.amazonaws.com/" // this is the target file and location in S3
	//path := "images/nature/"
	path := "images/all"
	fmt.Println("Processing")
	//fileToBeUploaded := "./images/computer tablet phone_0.jpg"
	//fileToBeUploaded:=req.Form["img"];
	fmt.Println("Uploading in process")
	file, err := os.Open(conf.PathToTheFolderContainingImages + "/" + fileToBeUploaded)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)

	// read into buffer
	buffer := bufio.NewReader(file)
	_, err = buffer.Read(bytes)

	filetype := http.DetectContentType(bytes)

	err = bucket.Put(path, bytes, filetype, s3.ACL("public-read"))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	views.ShowSuccessOrErrorAsJSON(rw, "ImageUploadSuccessful", "The image successfully uploaded to AMAZON S3", 125, 25)
	fmt.Printf("[Synkku] Your file successfully Uploaded to %s with %v bytes to S3\n\n", path, size)
}









/*********** Uploading images to Amazon S3 server *****************/
func DownloadImageFromAmazonS3(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	if !isUserLoggedIn() {
		views.ShowSuccessOrErrorAsJSON(rw, "LoginRequired", "You need to login to access the site. Visit http:127.0.0.1:9000/v1/feeds/login", 120, 20)
		return
	}

	if req.Method != "POST" { //Checking whether the data is being sent using POST method or not(If already implemented in Android then not required)
		views.ShowSuccessOrErrorAsJSON(rw, "POSTMethodNotFoundError", "The data is not being requested using GET method", 112, 12)
		fmt.Println("The data is not being requested using GET method")
		return
	}

	req.ParseForm()
	//fmt.Printf("len(req.Form) : %d %T\n",len(req.Form),req.Form);
	if len(req.Form) == 0 { //If there's no post data then inform the user
		views.ShowSuccessOrErrorAsJSON(rw, "NoPostDataError", "You haven't sent the GET data", 109, 9)
		fmt.Println("You haven't send the POST data")
		return
	}
	//fileToBeUploaded := "./images/computer tablet phone_0.jpg"
	fileToBeDownloaded := req.Form["img"][0]
	targetFileName := req.Form["target_file_name"][0]
	if !areAllFieldsAreInValidForm(fileToBeDownloaded) {
		views.ShowSuccessOrErrorAsJSON(rw, "EmptyFieldError", "You have to specify the name of image (Make sure the image should be inside ./images/)", 103, 3)
		return
	}

	AWSAuth := aws.Auth{
		AccessKey: "AKIAI4RV6AOJMQQE5UJQ", // change this to yours
		SecretKey: "8KOd3w63SL8BWXoC7N5cmR/HyJ9zq9WBG+gvZJCg",
	}
	region := aws.APNortheast
	connection := s3.New(AWSAuth, region)

	bucket := connection.Bucket("tiiing1") // change this your bucket name
	//"tiiing1.s3-website-ap-northeast-1.amazonaws.com/"
	//path := "tiiing1.s3-website-ap-northeast-1.amazonaws.com/" // this is the target file and location in S3
	//path := "images/nature/"
	path := "images/all" //Don't specify the path like this images/any/
	fmt.Println("Processing")
	//fileToBeUploaded := "./images/computer tablet phone_0.jpg"
	//fileToBeUploaded:=req.Form["img"];
	downloadBytes, err := bucket.Get(path)

	if err != nil {
		views.ShowSuccessOrErrorAsJSON(rw, "KeyNotFoundError", "The specified key does not exist", 127, 27)
		os.Exit(1)
	}
	targetFileName = time.Now().String()[0:19] + targetFileName
	downloadFile, err := os.Create(conf.PathToFolderForStoringDownloadedImages + "/" + targetFileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer downloadFile.Close()

	downloadBuffer := bufio.NewWriter(downloadFile)
	downloadBuffer.Write(downloadBytes)

	io.Copy(downloadBuffer, downloadFile)

	fmt.Printf("Downloaded from S3 and saved to download.jpg. \n\n")
	views.ShowSuccessOrErrorAsJSON(rw, "ImageDownloadSuccessful", "The image successfully downloaded from AMAZON S3", 125, 25)
	fmt.Printf("[Synkku] Your file successfully Downloaded to %s from AMAZON S3", path)
}
