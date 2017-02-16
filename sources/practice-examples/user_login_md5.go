/********************** Middleware (Token validator) ***********************/
func isAppSignInKeyCorrect(appSignInKeyMd5 string,authenticationToken string,email string) bool{
	reversed_email:=[]byte(email); //string is immutable in Go
	//email reversing for loop
	for front,last:=0,len(reversed_email)-1;front<last;front,last=front+1,last-1{
		reversed_email[front],reversed_email[last]=reversed_email[last],reversed_email[front];
	}
	reversed_email_and_auth_token:=string(reversed_email)+authenticationToken;

	fmt.Printf("Your entered app_sign_key       :  %v (%T)",reversed_email_and_auth_token,reversed_email_and_auth_token)
	fmt.Println()
	fmt.Println("Your app generated app_sign_key : ",appSignInKeyMd5)

	md5_of_reversed_email_and_auth_token:=getMD5Hash(reversed_email_and_auth_token)
	fmt.Println("Server generated app_sign_key   : ",md5_of_reversed_email_and_auth_token)

	return md5_of_reversed_email_and_auth_token==appSignInKeyMd5
}

func getStoken(userId int, authenticationToken string) string{
	unixTime:=strconv.FormatInt(time.Now().UTC().Unix(), 10)
	tempStoken:= strconv.Itoa(userId)+authenticationToken+unixTime //<userId><authenticationToken><unixTime>
	stoken:=getMD5Hash(tempStoken)			 //md5("<userId><authenticationToken><unixTime>")
	return stoken
}

func getMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}

func IsEmailValid(email string) bool {
    Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
    return Re.MatchString(email);
}

/********************** User Login ****************************************/
func UserLogin(rw http.ResponseWriter, req *http.Request) {
	const AuthenticationKey="eZ$21#@54>4074W8Ndkf**WE32awe2376THWEKm"
	//Not checking for POST method
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")

	if req.Method != "POST"{ //Checking whether the data is being sent using POST method or not(If already implemented in Android then not required)
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
	un := strings.TrimSpace(req.Form["un"][0])		   //username(optional)
	umail := strings.TrimSpace(req.Form["umail"][0])  //p_pic(required), ltype, atype & cmtyid is not required for login, I need to check
	p_pic :=strings.TrimSpace(req.Form["p_pic"][0])   //From google  (optional)
	cmtyid :=strings.TrimSpace(req.Form["cmtyid"][0]) //community id (required)
	ltype:=strings.TrimSpace(req.form["ltype"][0])    //1 for google (required)
	atype:=strings.TrimSpace(req.Form["atype"][0])    //Student, Job Seeker, Student, Employed (required)
	app_sign_in_key:=strings.TrimSpace(req.Form["appsigninkey"][0]) //app sign in key while first login (required)


	if !areAllFieldsAreInValidForm(umail,app_sign_in_key,cmtyid,ltype,actype) {
		views.ShowSuccessOrErrorAsJSON(rw, "EmptyFieldError", "You have to select all the fileds,profile pic should be empty", 103, 3)
		return
	}

	if !isAppSignInKeyCorrect(app_sign_in_key,AuthenticationKey,umail){ //Call to a Middleware
		views.ShowSuccessOrErrorAsJSON(rw,"Invalid Token","The token is invalid",132,32)
		fmt.Println("The token is invalid")
		return
	}

	if !IsEmailValid(umail){
		view.ShowSuccessOrErrorAsJSON(rw,"InvalidEmailError","The enetered email is not valid",133,33)
		fmt.Println("Email is not valid")
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
	rows, err := db.Query("select  user_name, email_id from users where user_name='" + un + "' AND email_id= '" + umail + "' AND status=1;")
	if err != nil {
		panic(err.Error())
	}

	found := false
	for rows.Next() {
		found = true
	}
	if found {
		rows, err = db.Query("select user_id, profilepic_url, profilevideo_url, community_id, account_type from users" +
			" where user_name='" + un + "' AND " + " email_id='" + umail + "' ;") //I need to check about stoken

		var ppic, pvid, stoken string
		var cmtid, act, uid int
		stoken=getStoken(uid,AuthenticationKey);
		rows.Next()
		rows.Scan(&uid, &ppic, &pvid, &cmtid, &act)

		stmt,err:=db.Prepare("update users set stoken='"+stoken+"' , stoken_updated_on='"+time.Now().String()[0:19]+"' where email_id='"+umail+"'");
		if err!=nil{
			fmt.Println("Error in execution of update query")
			views.ShowSuccessOrErrorAsJSON(rw,"UpdateQueryExecutionError","Error in execution of update query",134,34)
			return
		}
		fmt.Fprintf(rw, views.UniversalMessageCreatorAsJSON([]string{"success", "message", "uid", "p_pic", "p_vid", "cmtyid", "atype", "stoken"}, 1, "You successfully logged in", uid, ppic, pvid, cmtid, act, stoken))
		//http.Redirect(rw, req, "/", 301)
	}else{
		views.ShowSuccessOrErrorAsJSON(rw, "UserDoesNotExist", "This user does not exist, Try again", 101, 1)
		return
	}
}