/************************ User Login ****************************/
func UserLogin2(rw http.ResponseWriter, req *http.Request) {
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
	un := req.Form["un"][0]
	umail := req.Form["umail"][0] //p_pic, ltype, atype & cmtyid is not required for login, I need to check

	if !areAllFieldsAreInValidForm(un, umail) {
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
	rows, err := db.Query("select  user_name, email_id from users where user_name='" + un + "' AND email_id= '" + umail + "';")
	if err != nil {
		panic(err.Error())
	}

	found := false
	for rows.Next() {
		found = true
	}
	if found {
		rows, err = db.Query("select user_id, profilepic_url, profilevideo_url, community_id, account_type, stoken from users" +
			" where user_name='" + un + "' AND " + " email_id='" + umail + "' ;") //I need to check about stoken
		var ppic, pvid, stoken string
		var cmtid, act, uid int
		rows.Next()
		rows.Scan(&uid, &ppic, &pvid, &cmtid, &act, &stoken)
		fmt.Fprintf(rw, views.UniversalMessageCreatorAsJSON([]string{"success", "message", "uid", "p_pic", "p_vid", "cmtyid", "atype", "stoken"}, 1, "You successfully loggen in", uid, ppic, pvid, cmtid, act, stoken))
		//http.Redirect(rw, req, "/", 301)
	} else {
		views.ShowSuccessOrErrorAsJSON(rw, "UserDoesNotExist", "This user does not exist, Try again", 101, 1)
		return
	}
}
