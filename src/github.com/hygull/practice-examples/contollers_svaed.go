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

/******************Home page Handler(Temporary)**************************/
func HomePage(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	if !isUserLoggedIn() {
		views.ShowSuccessOrErrorAsJSON(rw, "LoginRequired", "You need to login to access the site. Visit http:127.0.0.1:9000/v1/feeds/login", 120, 20)
		return
	}
	type Link struct {
		Link string `json:"Url"`
		Id   int    `json:"Id"`
	}
	type Links struct {
		UserPage Link `json:"UrlDetails"`
	}
	var jsonStrData []byte
	linksCollection := map[string]Links{}
	linksCollection["NewPostSubmissionLink"] = Links{Link{"http://127.0.0.1:9000/v1/feeds/newpost", 1}}
	linksCollection["PostUpdationLink"] = Links{Link{"http://127.0.0.1:9000/v1/feeds/updatepost", 2}}
	linksCollection["PostsListingLink"] = Links{Link{"http://127.0.0.1:9000/v1/feeds/getpostslist", 3}}
	linksCollection["ProblemReportLink"] = Links{Link{"http://127.0.0.1:9000/v1/feeds/problemreport", 4}}
	linksCollection["FeedbackMessageLink"] = Links{Link{"http://127.0.0.1:9000/v1/feeds/feedbackmsg", 5}}
	linksCollection["CityStatusChangingLink"] = Links{Link{"http://127.0.0.1:9000/v1/feeds/deactivatecityoffer", 6}}
	jsonStrData, _ = json.Marshal(linksCollection)
	views.ShowDataAsJSON(rw, string(jsonStrData))
} //End of function

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

/********************** Middleware (Token validator) ***********************/
func isAppSignInKeyCorrect(appSignInKeyMd5 string, authenticationToken string, email string) bool {
	reversed_email := []byte(email) //string is immutable in Go
	//email reversing for loop
	for front, last := 0, len(reversed_email)-1; front < last; front, last = front+1, last-1 {
		reversed_email[front], reversed_email[last] = reversed_email[last], reversed_email[front]
	}
	reversed_email_and_auth_token := string(reversed_email) + authenticationToken

	fmt.Printf("Reversed email and auth token      :  %v (%T)", reversed_email_and_auth_token, reversed_email_and_auth_token)
	fmt.Println()
	fmt.Println("\nYour app generated app_sign_key : ", appSignInKeyMd5)

	md5_of_reversed_email_and_auth_token := getMD5Hash(reversed_email_and_auth_token)
	fmt.Println("Server generated app_sign_key   : ", md5_of_reversed_email_and_auth_token)

	return md5_of_reversed_email_and_auth_token == appSignInKeyMd5
}

/*************************START (User login)******************************************/

func getStoken(userId int, authenticationToken string) string {
	unixTime := strconv.FormatInt(time.Now().UTC().Unix(), 10)
	tempStoken := strconv.Itoa(userId) + authenticationToken + unixTime //<userId><authenticationToken><unixTime>
	stoken := getMD5Hash(tempStoken)                                    //md5("<userId><authenticationToken><unixTime>")
	return stoken
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func IsEmailValid(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}

/*************************END of callee of UserLogin2 (User login)***********************************************/

/********************** User Login ******************************************************/
func UserLogin(rw http.ResponseWriter, req *http.Request) {
	const AuthenticationKey = "eZ$21#@54>4074W8Ndkf**WE32awe2376THWEKm"
	//Not checking for POST method
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")

	if req.Method != "POST" { //Checking whether the data is being sent using POST method or not(If already implemented in Android then not required)
		views.ShowSuccessOrErrorAsJSON(rw, "PostMethodNotFoundError", "The data is not being sent using POST method", 107, 7)
		fmt.Println("The data is not being sent using POST method")
		return
	}
	// if isUserLoggedIn() {
	// 	views.ShowSuccessOrErrorAsJSON(rw, "AlreadyLoggedIn", "2 users are not allowed to access the same server's port, First Logout then Login", 122, 22)
	// 	return
	// }
	req.ParseForm()

	if len(req.Form) == 0 { //If there's no post data then inform the user
		views.ShowSuccessOrErrorAsJSON(rw, "NoPostDataError", "You haven't sent the POST data", 104, 4)
		fmt.Println("You haven't send the POST data")
		return
	}
	un := strings.TrimSpace(req.Form["un"][0])                        //username(optional)
	umail := strings.TrimSpace(req.Form["umail"][0])                  //p_pic(required), ltype, atype & cmtyid is not required for login, I need to check
	p_pic := strings.TrimSpace(req.Form["p_pic"][0])                  //From google  (optional)
	cmtyid := strings.TrimSpace(req.Form["cmtyid"][0])                //community id (required)
	ltype := strings.TrimSpace(req.Form["ltype"][0])                  //1 for google (required)
	atype := strings.TrimSpace(req.Form["atype"][0])                  //Student, Job Seeker, Student, Employed (required)
	app_sign_in_key := strings.TrimSpace(req.Form["appsigninkey"][0]) //app sign in key while first login (required)

	fmt.Println("Extracted URL Parameters : ", un, umail, p_pic, cmtyid, ltype, atype, app_sign_in_key)
	if !areAllFieldsAreInValidForm(umail, app_sign_in_key, cmtyid, ltype, atype) {
		views.ShowSuccessOrErrorAsJSON(rw, "EmptyFieldError", "You have to select all the fileds,profile pic & username should be empty", 103, 3)
		return
	}

	if !IsEmailValid(umail) {
		views.ShowSuccessOrErrorAsJSON(rw, "InvalidEmailFormatError", "The enetered email is not valid", 133, 33)
		fmt.Println("Email is not valid")
		return
	}
	if !isAppSignInKeyCorrect(app_sign_in_key, AuthenticationKey, umail) { //Call to a Middleware
		views.ShowSuccessOrErrorAsJSON(rw, "InvalidAppSignKey", "App signin key is invalid", 132, 32)
		fmt.Println("The token is invalid")
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
	rows, err := db.Query("select  user_name, email_id from users where email_id= '" + umail + "' AND active=1;")
	if err != nil {
		panic(err.Error())
	}

	found := false
	for rows.Next() {
		found = true
	}
	if found {
		rows, err = db.Query("select user_id, profilepic_url, profilevideo_url, community_id, account_type from users" +
			" where email_id='" + umail + "' ;") //I need to check about stoken
		fmt.Println("Details extracted from DB")
		var ppic, pvid, stoken string
		var cmtid, act, uid int
		stoken = getStoken(uid, AuthenticationKey)
		rows.Next()
		rows.Scan(&uid, &ppic, &pvid, &cmtid, &act)

		stmt, err := db.Prepare("update users set stoken='" + stoken + "' , stoken_updated_on='" + time.Now().String()[0:19] + "' where email_id='" + umail + "';")
		if err != nil {
			fmt.Println("Error in execution of update query")
			views.ShowSuccessOrErrorAsJSON(rw, "UpdateQueryExecutionError", "Error in execution of update query", 134, 34)
			return
		}
		stmt.Exec()
		fmt.Println("Details updated as the user is already having an account")
		fmt.Fprintf(rw, views.UniversalMessageCreatorAsJSON([]string{"success", "message", "uid", "p_pic", "p_vid", "cmtyid", "atype", "stoken"}, 1, "You successfully logged in", uid, ppic, pvid, cmtid, act, stoken))
		//http.Redirect(rw, req, "/", 301)
	} else {
		timeNow := time.Now().String()[0:19]

		query := "insert into users(user_name, email_id, date_of_birth, city_id, gender, " +
			" marital_status, studied_at, employed_at, community_id, profilepic_url, profilevideo_url, login_type, " +
			"account_type,stoken , stoken_updated_on, details_updated_on, updated_on, active, " +
			" deleted_on ) values(" + "'" + un + "' , '" + umail + "' , '1900-01-01',0, 0, 0, 'Default School', 'None', " +
			cmtyid + ", '" + p_pic + "', '', 0, 1, '' , '1900-01-01 12:00:00', '" + timeNow + "' , '" + timeNow +
			"' , 1, '1900-01-01 12:00:00');"
		fmt.Println(query)
		fmt.Println("Preparing to insert new details on the Database")
		stmt, err := db.Prepare(query)
		if err != nil {
			fmt.Println("Error in query execution")
			views.ShowSuccessOrErrorAsJSON(rw, "QueryExecutionError", "Error in executing the insert query", 127, 27)
			return
		}
		stmt.Exec()
		fmt.Println("New details inserted into Database(A new user login)")

		rows, err = db.Query("select user_id, profilepic_url, profilevideo_url, community_id, account_type from users" +
			" where email_id='" + umail + "' ;") //I need to check about stoken
		fmt.Println("Details extracted from DB")
		var ppic, pvid, stoken string
		var cmtid, act, uid int
		stoken = getStoken(uid, AuthenticationKey)
		rows.Next()
		rows.Scan(&uid, &ppic, &pvid, &cmtid, &act)

		stmt, err = db.Prepare("update users set stoken='" + stoken + "' , stoken_updated_on='" + time.Now().String()[0:19] + "' where email_id='" + umail + "';")
		if err != nil {
			fmt.Println("Error in execution of update query")
			views.ShowSuccessOrErrorAsJSON(rw, "UpdateQueryExecutionError", "Error in execution of update query", 134, 34)
			return
		}
		stmt.Exec()
		fmt.Println("New user details updated")
		fmt.Fprintf(rw, views.UniversalMessageCreatorAsJSON([]string{"success", "message", "uid", "p_pic", "p_vid", "cmtyid", "atype", "stoken"}, 1, "You successfully logged in", uid, ppic, pvid, cmtid, act, stoken))
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

/*********************Storing new post details to database **************/
func StoreNewPostToDb(rw http.ResponseWriter, req *http.Request) {
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

	userId := req.Form["uid"][0] //strconv.Itoa(SynkkuLoggedInUsersId)
	visibilityType := req.Form["vbt"][0]
	communityId := req.Form["cmtyid"][0]
	mediaType := req.Form["mt"][0]
	textMessage := req.Form["txt"][0]
	pType := req.Form["p_type"][0] //1=new, 2=shared
	basePostId := req.Form["bpid"][0]
	var recentUserId, recentPostId string //We will get this data from DB
	//recentUserId := req.Form["re_us_id"][0] /*This should come from DB*/
	//recentPostId := req.Form["re_po"][0]    /*This should come from DB*/

	if !areAllFieldsAreInValidForm(userId, visibilityType, communityId, mediaType, textMessage,
		pType, basePostId) {
		views.ShowSuccessOrErrorAsJSON(rw, "EmptyFieldError", "You have to select all the fileds", 103, 3)
		return
	}
	//If all the data are not blank
	fmt.Println("Storing details: ", len(userId), visibilityType, communityId, mediaType, textMessage, pType, basePostId)

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
		fmt.Println("All the details of users successfully retreived from DB....")
	}

	found := false
	for rows.Next() {
		found = true
	}
	if found == false {
		views.ShowSuccessOrErrorAsJSON(rw, "UserDoesNotExist", "This user id does not exist in the database", 101, 1)
	} else {
		fmt.Println("Preparing query")

		//fmt.Println("insert into posts(user_id,visibility_type,community_id,media_type,text_message,status,base_postid,recent_postid,recent_user_id,created_on,updated_on) values(" + userId + "," + visibilityType + "," + communityId + "," + mediaType + ",'" + textMessage + "'," + pType + "," + basePostId + "," + recentPostId + "," + recentUserId + ",'19920514000509','19920514000509');")  //00000000000000 is not a valid one, it causes the SQL to not perform the jod in right way
		currentTime := time.Now().String()[0:19]
		r, e := db.Query("select user_id,post_id from posts order by created_on desc limit 1;")
		if e != nil {
			fmt.Println("Error in execution for getting user_id & post_id query")
			return
		}
		found = false
		if r.Next() {
			r.Scan(&recentUserId, &recentPostId)
			found = true
		}
		recentIntPostId, _ := strconv.Atoi(recentPostId)

		if found == true {
			fmt.Println("Recent post_id, user_id found")
			stmt, err2 := db.Prepare("insert into posts(base_postid,recent_postid,recent_user_id,user_id,visibility_type,community_id,media_type,post_type,text_message,created_on,updated_on) values(" + basePostId + "," + recentPostId + "," + recentUserId + ", " +
				userId + "," + visibilityType + ",'" + communityId + "'," + mediaType + "," + pType + ",'" + textMessage + "', '" +
				currentTime + "','" + currentTime + "');") //The current date & time of system will be the value of created_on and updated_on
			stmt.Exec()
			if err2 != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Query successfully executed")
				fmt.Println("1 post successfully stored into DataBase")
				//views.ShowSuccessOrErrorAsJSON(rw, "PostSubmissionSucceeded", "Post successfully posted", 102, 2)
				fmt.Fprintf(rw, views.UniversalMessageCreatorAsJSON([]string{"success", "message", "pid", "u_at", "c_at"}, 1, "Posted successfully", recentIntPostId+1, currentTime, currentTime))
			}
		} else {
			fmt.Println("Recent post_id, user_id does not exist")
			stmt, err2 := db.Prepare("insert into posts(base_postid,recent_postid,recent_user_id,user_id,visibility_type,community_id,media_type,post_type,text_message,created_on,updated_on) values(" + basePostId + "," + " 1 " + "," + userId + ", " +
				userId + "," + visibilityType + ",'" + communityId + "'," + mediaType + "," + pType + ",'" + textMessage + "', '" +
				currentTime + "','" + currentTime + "');")
			stmt.Exec()
			if err2 != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Query successfully executed")
				fmt.Println("1 post successfully stored into DataBase")
				//views.ShowSuccessOrErrorAsJSON(rw, "PostSubmissionSucceeded", "Post successfully posted", 102, 2)
				fmt.Fprintf(rw, views.UniversalMessageCreatorAsJSON([]string{"success", "message", "pid", "u_at", "c_at"}, 1, "Posted successfully", recentIntPostId+1, currentTime, currentTime))
			}
		}

	}
} //End of function storeNewPostToDb

/*********************Storing new post details to database **************/
func UpdatePost(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	if !isUserLoggedIn() {
		views.ShowSuccessOrErrorAsJSON(rw, "LoginRequired", "You need to login to access the site. Visit http:127.0.0.1:9000/v1/feeds/login", 120, 20)
		return
	}

	fmt.Println("Req method : ", req.Method)
	if req.Method != "POST" && req.Method != "PUT" { //Checking whether the data is being sent using POST method or not(If already implemented in Android then not required)
		views.ShowSuccessOrErrorAsJSON(rw, "PutOrPostMethodNotFoundError", "The data is not being sent using POST method", 107, 7)
		fmt.Println("The data is not being sent using POST method")
		return
	}

	req.ParseForm()

	if len(req.Form) == 0 { //If there's no post data then inform the user
		views.ShowSuccessOrErrorAsJSON(rw, "NoPutOrPostDataError", "You haven't provided the PUT/Updating information", 109, 9)
		fmt.Println("You haven't send the PUT/Updating data")
		return
	}

	userId := strings.TrimSpace(req.Form["uid"][0])
	textMsg := strings.TrimSpace(req.Form["txt"][0]) //To remove white spaces at end, before updating the data in DB
	status := strings.TrimSpace(req.Form["sts"][0])  //It's not required to chech this for whitespaces, as this is number, it depends
	postId := strings.TrimSpace(req.Form["pid"][0])

	fmt.Println("POST DATA TO UPDATE : ", userId, status, postId)
	if !areAllFieldsAreInValidForm(userId, postId, status) { //Send data to variadic function for checking the emptiness
		views.ShowSuccessOrErrorAsJSON(rw, "EmptyFieldError", "You have to fill all 3 the fileds for updation", 108, 8)
		return
	}

	fmt.Println("Preparing to update post details in Database ")
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

	rows, e := db.Query("select post_id from posts where post_id=" + postId + " AND user_id=" + userId + ";") //This
	if e != nil {
		panic(e)
	} else {
		fmt.Println("postId successfully retreived from DB....")
	}

	found := false
	for rows.Next() {
		found = true
	}

	if found == false {
		views.ShowSuccessOrErrorAsJSON(rw, "PostDoesNotExist", "This post id does not exist in the database", 110, 10)
	} else {
		fmt.Println("Preparing query")

		//fmt.Println("insert into posts(user_id,visibility_type,community_id,media_type,text_message,status,base_postid,recent_postid,recent_user_id,created_on,updated_on) values(" + userId + "," + visibilityType + "," + communityId + "," + mediaType + ",'" + textMessage + "'," + pType + "," + basePostId + "," + recentPostId + "," + recentUserId + ",'19920514000509','19920514000509');")  //00000000000000 is not a valid one, it causes the SQL to not perform the jod in right way
		updatedTime := time.Now().String()[0:19]
		if status == "1" {
			if textMsg != "" {
				stmt, err2 := db.Prepare("update posts set text_message='" + textMsg + "', status=" + //The current date & time of system will be the vale of updated_on
					status + ", updated_on='" + updatedTime + "' where post_id=" + postId + " AND user_id=" + userId + ";")

				if err2 != nil {
					fmt.Println(err.Error())
					return
				} else {
					stmt.Exec()
					fmt.Println("Post details successfully executed")
					//views.ShowSuccessOrErrorAsJSON(rw, "PostUpdationSucceeded", "Post successfully updated", 111, 11)
					fmt.Fprintf(rw, views.UniversalMessageCreatorAsJSON([]string{"success", "message"}, 1, "Updated successfully"))
				}
			} else {
				views.ShowSuccessOrErrorAsJSON(rw, "BlankMessageFieldError", "Please specify the new message", 125, 25)
			}
		} else { /*status==2*/
			if status == "2" {
				if textMsg == "" { /*Execute query to delete the post*/
					stmt, err := db.Prepare("delete from posts where user_id=" + userId + " AND post_id=" + postId + " ;")
					if err != nil {
						fmt.Println("Error in executing the DELETE Query")
						return
					} else {
						stmt.Exec()
						fmt.Println("1 post successfully deleted from the database, post_id: ", postId)
						fmt.Fprintf(rw, views.UniversalMessageCreatorAsJSON([]string{"success", "message"}, "1", "Updated/Deleted successfully"))
					}
				} else {
					views.ShowSuccessOrErrorAsJSON(rw, "MessageSpecificationError", "As you have specified status=2(to delete the post), So message will be blank", 124, 24)
				}
			} else {
				views.ShowSuccessOrErrorAsJSON(rw, "BadStatusError", "The status should be either 1(active) or 2(inactive/delete)", 123, 23)
			}
		}
	}
}

/*********************Storing new post details to database **************/
func GetPostsList(rw http.ResponseWriter, req *http.Request) {

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	if !isUserLoggedIn() {
		views.ShowSuccessOrErrorAsJSON(rw, "LoginRequired", "You need to login to access the site. Visit http:127.0.0.1:9000/v1/feeds/login", 120, 20)
		return
	}
	if req.Method != "POST" { //Checking whether the data is being sent using POST method or not(If already implemented in Android then not required)
		views.ShowSuccessOrErrorAsJSON(rw, "POSTMethodNotFoundError", "The data is not being requested using POST method", 112, 12)
		fmt.Println("The data is not being requested using POST method")
		return
	}

	req.ParseForm()

	if len(req.Form) == 0 { //If there's no post data then inform the user
		views.ShowSuccessOrErrorAsJSON(rw, "NoPOSTDataError", "You haven't provided the fetching information", 109, 9)
		fmt.Println("You haven't provided the fetching data")
		return
	}
	userId := req.Form["uid"][0]
	uAt := req.Form["u_at"][0]
	ty := req.Form["ty"][0]
	cmtyId := req.Form["cmtyid"][0]

	if !areAllFieldsAreInValidForm(userId, uAt, ty, cmtyId) {
		views.ShowSuccessOrErrorAsJSON(rw, "EmptyFieldError", "You have to select all the fileds", 103, 3)
		return
	}

	fmt.Println("Connecting to DB...")
	db, err := sql.Open("mysql", "root:admin@321@tcp(127.0.0.1:3306)/synkku?charset=utf8")
	if err != nil {
		fmt.Println("")
		return
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Println("Ping Error : ", err.Error())
	} else {
		fmt.Println("Connection is available...")
	}

	type PostData struct {
		Pid     int    `json:"pid"`
		Uid     int    `json:"uid"`
		Un      string `json:"un"`
		Ppic    string `json:"p_pic"`
		Bpid    int    `jspn:"bpid"`
		Rpid    int    `json:"rpid"`
		Ruid    int    `json:"ruid"`
		Runame  string `json:"runame"`
		Vbt     int    `json:"vbt"`
		Cmtyid  int    `json:"cmtyid"`
		Mt      int    `json:"mt"`
		Txt     string `json:"txt"`
		Sts     int    `json:"sts"`
		Nlikes  int    `json:"nlikes"`
		Ncmts   int    `json:"ncmts"`
		Isliked int    `json:"isliked"`
		Ptype   int    `json:"p_type"`
		Cat     string `json:"c_at"`
		Uat     string `json:"u_at"`
	}
	var pid int
	var uid int
	var un string
	var p_pic string
	var bpid int
	var rpid int
	var ruid int
	var runame string //Remaining
	var vbt int
	var cmtyid int
	var mt int
	var txt string
	var sts int
	var nlikes int  //Remaining
	var ncmts int   //Remaining
	var isliked int //Remaining
	var ptype int
	var c_at string
	var u_at string

	PostArr := []PostData{}
	//RecentUsers:=[]string{}
	compOp := ""

	if ty == "1" {
		compOp = ">"
	} else {
		compOp = "<"
	}

	queryString := "select posts.post_id, posts.user_id, users.user_name,users.profilepic_url, " +
		"posts.base_postid, posts.recent_postid,posts.recent_user_id,posts.visibility_type, " +
		"posts.community_id,posts.media_type,posts.text_message, posts.status,posts.post_type, " +
		"posts.created_on, posts.updated_on from posts inner join users on " +
		"posts.recent_user_id=users.user_id where (posts.community_id=" + cmtyId + " and posts.updated_on" + compOp +
		"'" + uAt + "' AND posts.visibility_type=1) order by posts.updated_on desc limit 5;"
	fmt.Println(queryString)
	rows, err := db.Query(queryString)
	if err != nil {
		fmt.Println("Error in query execution")
		views.ShowSuccessOrErrorAsJSON(rw, "QueryExecutionError", "Error in execution of query", 127, 27)
		return
	}
	found := false
	for rows.Next() {
		rows.Scan(&pid, &uid, &un, &p_pic, &bpid, &rpid, &ruid, &vbt, &cmtyid, &mt, &txt, &sts, &ptype, &c_at, &u_at)
		//RecentUsers=append(RecentUsers,un)
		PostArr = append(PostArr, PostData{pid, uid, un, p_pic, bpid, rpid, ruid, runame, vbt, cmtyid, mt, txt, sts, nlikes, ncmts, isliked, ptype, c_at, u_at})
		found = true
	}

	if found == false {
		views.ShowSuccessOrErrorAsJSON(rw, "PostsNotFound", "As of now, no one has posted here", 130, 30)
		fmt.Println("No posts found in this community")
		return
	}
	fmt.Fprintf(rw, views.UniversalMessageCreatorAsJSON([]string{"success", "message", "p_arr"}, 1, "New posts are", PostArr))
	fmt.Println("Posts have been successfully displayed...")
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

/*********************Report a problem***********************************/
func FeedbackMessageSender(rw http.ResponseWriter, req *http.Request) {
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

	userId := req.Form["uid"][0]
	feedbackMessage := req.Form["febk_mes"][0]

	if !areAllFieldsAreInValidForm(userId, feedbackMessage) {
		views.ShowSuccessOrErrorAsJSON(rw, "EmptyFieldError", "You have to select all the fileds", 103, 3)
		return
	}
	//If all the data are not blank
	fmt.Println("Validating entered details: ", userId, feedbackMessage)
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
		fmt.Println("userId and postId details successfully retreived from DB....")
	}

	found := false
	for rows.Next() {
		found = true
	}

	if found == false {
		views.ShowSuccessOrErrorAsJSON(rw, "UserDoesNotExist", "This user id does not exist in the database", 101, 1)
	} else {
		now := time.Now().String()[0:19]
		stmt, err2 := db.Prepare("insert into feedback(user_id, fback_message, created_on, updated_on) values(" + userId + ", '" + feedbackMessage + "' , '" + now + "', '" + now + "');")
		if err2 != nil {
			fmt.Println("Problem during insertion of feedback")
			views.ShowSuccessOrErrorAsJSON(rw, "FeedbackInsertionError", "Error in inserting the feedback to DB", 128, 28)
			return
		}
		stmt.Exec()
		feedBackId := 1
		found = false
		rows, err3 := db.Query("select fback_id from  feedback order by created_on desc limit 1;")
		if err3 != nil {
			views.ShowSuccessOrErrorAsJSON(rw, "QueryExecutionError", "There's an error in select query execution", 127, 27)
			fmt.Println("Error in execution of select query")
			return
		}
		for rows.Next() {
			rows.Scan(&feedBackId)
			found = true
		}

		fmt.Println("Feedback Query successfully executed")
		fmt.Println("Feedback successfully sent")
		//views.ShowSuccessOrErrorAsJSON(rw, "ProblemReportSent", "Probelm Report successfully sent", 115, 15)
		fmt.Fprintf(rw, views.UniversalMessageCreatorAsJSON([]string{"success", "message", "febk_id"}, 1, "We received your message, We will get back you soon", feedBackId))
	}
} //End of function feedbackMessage

/*********************Report a problem***********************************/
func ProblemReportSender(rw http.ResponseWriter, req *http.Request) {
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

	userId := req.Form["uid"][0]
	problemMessage := req.Form["prbl_mes"][0]

	if !areAllFieldsAreInValidForm(userId, problemMessage) {
		views.ShowSuccessOrErrorAsJSON(rw, "EmptyFieldError", "You have to select all the fileds", 103, 3)
		return
	}
	//If all the data are not blank
	fmt.Println("Validating entered details: ", userId, problemMessage)
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
		fmt.Println("All the details of user successfully retreived from DB....")
	}

	found := false
	for rows.Next() {
		found = true
	}
	if found == false {
		views.ShowSuccessOrErrorAsJSON(rw, "UserDoesNotExist", "This user id does not exist in the database", 101, 1)
	} else {
		now := time.Now().String()[0:19]
		stmt, err2 := db.Prepare("insert into reportaproblem(user_id, problem, created_on, updated_on) values(" + userId + ", '" + problemMessage + "' , '" + now + "', '" + now + "');")
		if err2 != nil {
			fmt.Println("Problem during insertion of feedback")
			views.ShowSuccessOrErrorAsJSON(rw, "ProblemReportInsertionError", "Error in inserting the problem report to DB", 126, 26)
			return
		}
		stmt.Exec()
		recentPostId := 1
		found = false
		rows, err3 := db.Query("select post_id from posts order by created_on desc limit 1;")
		if err3 != nil {
			views.ShowSuccessOrErrorAsJSON(rw, "QueryExecutionError", "There's an error in select query execution", 127, 27)
			fmt.Println("Error in execution of select query")
			return
		}
		for rows.Next() {
			rows.Scan(&recentPostId)
			found = true
		}

		fmt.Println("Report Query successfully executed")
		fmt.Println("Problem reported successfully")
		//views.ShowSuccessOrErrorAsJSON(rw, "ProblemReportSent", "Probelm Report successfully sent", 115, 15)
		fmt.Fprintf(rw, views.UniversalMessageCreatorAsJSON([]string{"success", "message", "rpid"}, 1, "We received your message, We will get back you soon", recentPostId))
	}
} //End of function reportAProblem

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

/**********************Getting communities list *********************/
func DisplayCommunities(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	if !isUserLoggedIn() {
		views.ShowSuccessOrErrorAsJSON(rw, "LoginRequired", "You need to login to access the site. Visit http:127.0.0.1:9000/v1/feeds/login", 120, 20)
		return
	}
	fmt.Println("Req method : ", req.Method)
	if req.Method != "POST" { //Checking whether the data is being sent using POST method or not(If already implemented in Android then not required)
		views.ShowSuccessOrErrorAsJSON(rw, "POSTMethodNotFoundError", "The data is not being requested using POST method", 112, 12)
		fmt.Println("The data is not being requested using POST method")
		return
	}

	req.ParseForm()
	if len(req.Form) == 0 { //If there's no post data then inform the user
		views.ShowSuccessOrErrorAsJSON(rw, "NoPOSTDataError", "You haven't provided the fetching information", 109, 9)
		fmt.Println("You haven't provided the fetching data")
		return
	}

	usId := req.Form["uid"][0]
	if !areAllFieldsAreInValidForm(usId) {
		views.ShowSuccessOrErrorAsJSON(rw, "EmptyFieldError", "You have to select all the fileds", 103, 3)
		return
	}

	fmt.Println("Preparing to extract posts from Database")
	db, err := sql.Open("mysql", conf.DBUserName+":"+conf.DBPassword+"@tcp"+"("+conf.DBHost+":"+conf.DBPort+")/"+conf.DBName) //Table name
	fmt.Println("Connected...")
	if err != nil {
		fmt.Println("This is the error : ", err.Error())
	}
	defer db.Close()
	//make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Println("Ping Error : ", err.Error())
	} else {
		fmt.Println("Connection is available...")
	}
	rows, e := db.Query("select user_id from users where user_id=" + usId + ";")
	if e != nil {
		panic(e)
	} else {
		fmt.Println("userId and postId details successfully retreived from DB....")
	}

	found := false
	for rows.Next() {
		found = true
	}

	if found == false {
		views.ShowSuccessOrErrorAsJSON(rw, "UserDoesNotExist", "This user id does not exist in the database", 101, 1)
		return
	}
	rows, e = db.Query("select id, name from community;")
	if e != nil {
		views.ShowSuccessOrErrorAsJSON(rw, "QueryExecutionError", "Error in executing select query for community", 127, 27)
		return
	} else {
		fmt.Println("userId successfully retreived from DB....")
	}
	/*----------------------------------*/
	type CommunityDetail struct {
		ComId   int    `json:"cmtyid"`
		ComName string `json:"cmtyn"`
	}
	CommArr := []CommunityDetail{}
	var commId int
	var commName string
	/*----------------------------------*/
	found = false
	for rows.Next() {
		rows.Scan(&commId, &commName)
		CommArr = append(CommArr, CommunityDetail{commId, commName})
		found = true
	}
	if found == false {
		fmt.Println("No community found")
		views.ShowSuccessOrErrorAsJSON(rw, "NoCommunityFound", "There is no any community", 130, 30)
		return
	}
	fmt.Println("Successfully got the community list")
	fmt.Fprintf(rw, views.UniversalMessageCreatorAsJSON([]string{"success", "message", "cmty_arr"}, 1, "List of communities", CommArr))
}

/************************* To like a post *************************/
func LikeAPost(rw http.ResponseWriter, req *http.Request) {
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
	dataLength := len(req.Form)
	if dataLength == 0 { //If there's no post data then inform the user
		views.ShowSuccessOrErrorAsJSON(rw, "NoPostDataError", "You haven't sent the POST data", 104, 4)
		fmt.Println("You haven't send the POST data")
		return
	} else {
		if dataLength != 3 {
			views.ShowSuccessOrErrorAsJSON(rw, "UrlParamsError", "Some parameters are missed", 131, 31)
			fmt.Println("Please specify all the URL parameters")
			return
		}
	}

	uid := req.Form["uid"][0]
	pid := req.Form["pid"][0]
	isliked := req.Form["isliked"][0]

	if !areAllFieldsAreInValidForm(uid, pid, isliked) {
		views.ShowSuccessOrErrorAsJSON(rw, "EmptyFieldError", "You have to select all the fileds", 103, 3)
		return
	}
	fmt.Println("POST DATA FOR LIKES : ", uid, pid, isliked)
	fmt.Println("Preparing to update the likes table")
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

	rows, e := db.Query("select user_id from users where user_id=" + uid + ";")
	if e != nil {
		panic(e)
	} else {
		fmt.Println("All the details of users successfully retreived from DB....")
	}

	found := false
	for rows.Next() {
		found = true
	}
	if found == false {
		views.ShowSuccessOrErrorAsJSON(rw, "UserDoesNotExist", "This user id does not exist in the database", 101, 1)
	} else {
		fmt.Println("Preparing query")

		//fmt.Println("insert into posts(user_id,visibility_type,community_id,media_type,text_message,status,base_postid,recent_postid,recent_user_id,created_on,updated_on) values(" + userId + "," + visibilityType + "," + communityId + "," + mediaType + ",'" + textMessage + "'," + pType + "," + basePostId + "," + recentPostId + "," + recentUserId + ",'19920514000509','19920514000509');")  //00000000000000 is not a valid one, it causes the SQL to not perform the jod in right way
		r, e := db.Query("select user_id,post_id from likes where user_id=" + uid + " AND post_id=" + pid + ";")
		if e != nil {
			fmt.Println("Error in execution for getting user_id & post_id query")
			return
		}
		found = false
		if r.Next() {
			found = true
		}

		var now, updated_now string //To store the current system time
		var likes_count int         //To store total count
		var c_at string

		if found == false { //If data not found then insert
			now = time.Now().String()[0:19]
			stmt, e2 := db.Prepare("insert into likes(user_id, is_liked, post_id, liked_on, updated_on) values(" + uid + ", " + isliked + ", " + pid + ", " + "'" + now + "' , '" + now + "'" + ");")
			if e2 != nil {
				views.ShowSuccessOrErrorAsJSON(rw, "QueryExecutionError", "Error in execution of insert query", 127, 27)
				return
			}
			stmt.Exec()
		} else { //If found then update the is_liked
			updated_now = time.Now().String()[0:19]
			stmt, e2 := db.Prepare("update likes set is_liked=" + isliked + ", updated_on='" + updated_now + "' where user_id=" + uid + " AND post_id=" + pid + ";")
			if e2 != nil {
				views.ShowSuccessOrErrorAsJSON(rw, "QueryExecutionError", "Error in execution of insert query", 127, 27)
				return
			}
			stmt.Exec()
		} //else END
		if isliked == "1" {
			fmt.Println("Successfully liked the post")
		} else {
			fmt.Println("Successfully unliked the post")
		}
		rows, e = db.Query("select count(is_liked) as total_like_count from likes where post_id=" + pid + " AND is_liked=1;")
		if e != nil {
			fmt.Println("Error in execution of select query while accessing likes table during counting likes")
			views.ShowSuccessOrErrorAsJSON(rw, "QueryExecutionError", "Error in executing select query during counting likes", 127, 27)
			return
		}
		rows.Next()
		rows.Scan(&likes_count) //Storing the total likes count to likes_count variable
		fmt.Println("Updating the likes related to user_id : ", uid, " and post_id:", pid)
		rows, e = db.Query("select liked_on from likes where user_id=" + uid + " AND post_id=" + pid + ";")
		if e != nil {
			fmt.Println("Error in execution of select query while accessing likes table during getting created_on date")
			views.ShowSuccessOrErrorAsJSON(rw, "QueryExecutionError", "Error in executing select query during getting created_on date", 127, 27)
			return
		}
		rows.Next()      //Get
		rows.Scan(&c_at) //Store

		if found == true {
			fmt.Fprintf(rw, views.UniversalMessageCreatorAsJSON([]string{"success", "message", "tc", "c_at", "u_at"}, 1, "Your like data updated successfully", likes_count, c_at, updated_now))
		} else {
			fmt.Fprintf(rw, views.UniversalMessageCreatorAsJSON([]string{"success", "message", "tc", "c_at", "u_at"}, 1, "Your like data updated successfully", likes_count, now, now))
		}
	}
}

/******************** Get media files   *************************/
func GetMediaFiles(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	if req.Method != "POST" { //Checking whether the data is being sent using POST method or not(If already implemented in Android then not required)
		views.ShowSuccessOrErrorAsJSON(rw, "PostMethodNotFoundError", "The data is not being sent using POST method", 105, 5)
		fmt.Println("The data is not being sent using POST method")
		return
	}

	req.ParseForm()
	dataLength := len(req.Form)
	if dataLength == 0 { //If there's no post data then inform the user
		views.ShowSuccessOrErrorAsJSON(rw, "NoPostDataError", "You haven't sent the POST data", 104, 4)
		fmt.Println("You haven't send the POST data")
		return
	}

	post_id := req.Form["pid"][0]
	user_id := req.Form["uid"][0]
	if !areAllFieldsAreInValidForm(post_id, user_id) {
		views.ShowSuccessOrErrorAsJSON(rw, "EmptyFieldError", "You have to select all the fileds", 103, 3)
		return
	} else { //Start of else
		fmt.Println("Connecting to MySQL...")
		fmt.Println("Preparing to update the likes table")
		db, err := sql.Open("mysql", conf.DBUserName+":"+conf.DBPassword+"@tcp"+"("+conf.DBHost+":"+conf.DBPort+")/"+conf.DBName) //Table name
		fmt.Println("Connected...")
		if err != nil {
			fmt.Println("This is the error : ", err.Error())
		}
		defer db.Close()

		err = db.Ping()
		if err != nil {
			fmt.Println("Ping Error : ", err.Error()) //To make sure connection is available
		} else {
			fmt.Println("Connection started")
		}

		fmt.Println("sql is running")

		rows, e := db.Query("select post_id,user_id from media where post_id = " + post_id + " AND user_id = " + user_id + ";")
		if e != nil {
			panic(e)
		} else {
			fmt.Println("All the details of users successfully retreived from db")
		}
		found := false
		for rows.Next() {
			found = true
			break
		}

		if found == true {
			rows2, e2 := db.Query("select media_id,media_type,media_url,status,created_on,updated_on from media where user_id=" + user_id + " AND post_id=" + post_id + " AND status=1;")

			if e2 != nil {

				fmt.Println("error in query")
			} else {
				fmt.Println("All the details of users successfully retreived from db")

			}
			fmt.Println("ROWS DATA :\n", rows2)

			type MediaData struct {
				Mid   int    `json:"mid"`
				Mtype int    `json:"mtype"`
				Murl  string `json:"murl"`
				Sts   int    `json:"sts"`
				Cat   string `json:"c_at"`
				Uat   string `json:"u_at"`
			}
			MediaArr := []MediaData{}

			var media_id int
			var media_type int
			var media_url string
			var status int
			var created_on string
			var updated_on string

			fmt.Println("Getting media details from DB...")
			for rows2.Next() {
				rows2.Scan(&media_id, &media_type, &media_url, &status, &created_on, &updated_on)
				MediaArr = append(MediaArr, MediaData{media_id, media_type, media_url, status, created_on, updated_on})
			}
			fmt.Fprintf(rw, views.UniversalMessageCreatorAsJSON([]string{"success", "message", "me_arr"}, 1, "Media files are here", MediaArr))
		} else {
			fmt.Println(rw, "MediaNotFound", "This media id does not exist")
			//This check does not require...it's for to make sure for proper data input while testing the API
			views.ShowSuccessOrErrorAsJSON(rw, "MediaNotFound", "The media id related to the specified user does not exist in the database", 132, 32)
			return
		}
	}
}

/******************** Get comments *******************************************/
func GetComments(rw http.ResponseWriter, req *http.Request) {

}

/******************** Get media files   **************************************/
func GetListOfFriendRequest(rw http.ResponseWriter, req *http.Request) {

}

/******************** Comment a post *****************************************/
func CommentAPost(rw http.ResponseWriter, req *http.Request) {

}

/******************** Get notifications  *************************************/
func GetNotifications(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	if !isUserLoggedIn() {
		views.ShowSuccessOrErrorAsJSON(rw, "LoginRequired", "You need to login to access the site. Visit http:127.0.0.1:9000/v1/feeds/login", 120, 20)
		return
	}
	fmt.Println("Req method : ", req.Method)
	if req.Method != "POST" { //Checking whether the data is being sent using POST method or not(If already implemented in Android then not required)
		views.ShowSuccessOrErrorAsJSON(rw, "POSTMethodNotFoundError", "The data is not being requested using POST method", 112, 12)
		fmt.Println("The data is not being requested using POST method")
		return
	}

	req.ParseForm()
	if len(req.Form) == 0 { //If there's no post data then inform the user
		views.ShowSuccessOrErrorAsJSON(rw, "NoPOSTDataError", "You haven't provided the fetching information", 109, 9)
		fmt.Println("You haven't provided the fetching data")
		return
	}

	userid := req.Form["us_id"][0]
	updatedon := req.Form["u_at"][0]
	ty := req.Form["ty"][0]

	if !areAllFieldsAreInValidForm(userid, updatedon, ty) {
		views.ShowSuccessOrErrorAsJSON(rw, "EmptyFieldError", "You have to select all the fileds", 103, 3)
		return
	}

	fmt.Println("Preparing to extract posts from Database")
	db, err := sql.Open("mysql", conf.DBUserName+":"+conf.DBPassword+"@tcp"+"("+conf.DBHost+":"+conf.DBPort+")/"+conf.DBName) //Table name

	fmt.Println("Connected...")
	if err != nil {
		fmt.Println("This is the error : ", err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("Ping Error : ", err.Error())
	} else {
		fmt.Println("Connection is available...")
	}

	rows, e := db.Query("select user_id from users where user_id=" + userid + ";")
	if e != nil {
		panic(e)
	} else {
		fmt.Println("select query for userId successfully executed....")
	}

	found := false
	for rows.Next() {
		found = true
	}
	if found == false {
		views.ShowSuccessOrErrorAsJSON(rw, "UserDoesNotExist", "This user id does not exist in the database", 101, 1)
		return
	}
	type Likes struct {
		Nid    int    `json:"nid"`
		Pid    int    `json:"pid"`
		Ntype  int    `json:"ntype"`
		Ncount int    `json:"ncount"` //notififcations count
		Uid    int    `json:"uid"`
		Uname  string `json:"uname"`
		Ppic   string `json:"p_pic"`
		Cat    string `json:"c_at"`
		Uat    string `json:"u_at"`
	}

	var uname string
	var p_pic string
	var usid int
	var postid int
	var id int
	var c_at string
	var u_at string
	var likecount int
	var ncount int

	rows2, e := db.Query("select users.user_name ,users.profilepic_url,newtable.user_id,newtable.post_id,newtable.created_on,newtable.updated_on , " +
		"newtable.id, newtable.likecount  from users inner join ((select posts.user_id,posts.post_id, " +
		"notifications.created_on , notifications.id, notifications.updated_on,notifications.likecount from posts inner join notifications on " +
		"posts.post_id=notifications.post_id ) as newtable) on (users.user_id=newtable.user_id AND " +
		" users.user_id=" + userid + ");")
	if e != nil {
		views.ShowSuccessOrErrorAsJSON(rw, "QueryExecutionError", "Error in execution in query", 127, 27)
		return
	}

	NotificationsArr := []Likes{}
	for rows2.Next() {
		rows2.Scan(&uname, &p_pic, &usid, &postid, &c_at, &u_at, &id, &likecount)
		rows3, e := db.Query("select count(*) from notifications where post_id=" + strconv.Itoa(postid) + ";")
		if e != nil {
			views.ShowSuccessOrErrorAsJSON(rw, "QueryExecutionError", "Erro in execution in query", 127, 27)
			return
		}
		rows3.Next()
		rows3.Scan(&ncount)
		NotificationsArr = append(NotificationsArr, Likes{id, postid, ncount, likecount, usid, uname, p_pic, c_at, u_at})
	}
	fmt.Fprintf(rw, views.UniversalMessageCreatorAsJSON([]string{"success", "message", "ntf_arr"}, 1, "List in notifications", NotificationsArr))
	fmt.Println("Notifications list successfully displayed")
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
