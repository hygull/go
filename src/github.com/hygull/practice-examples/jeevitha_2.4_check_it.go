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
    db, err := sql.Open("mysql", "root:"+conf.SynkkuDataBasePassword+"@tcp(127.0.0.1:3306)/"+conf.SynkkuDataBaseName)  //Table name
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
        fmt.Println("userId  details successfully retreived from DB....")
    }

    found := false
    for rows.Next() {
        found = true
    }

    if found == false {
        views.ShowSuccessOrErrorAsJSON(rw, "UserDoesNotExist", "This user id does not exist in the database", 101, 1)
        return
    } 
    rows2, e2 := db.Query("select user_name,email_id,date_of_birth,community_id,profilepic_url,profilevideo_url,login_type,account_type,created_on,updated_on from users")
   if e != nil {
        views.ShowSuccessOrErrorAsJSON(rw,"QueryExecutionError","Error in executing select query for community",127,27);
        return
    } else {
        fmt.Println("userId successfully retreived from DB....")
    }
    /*----------------------------------*/
    type CommunityDetail struct{
        User_id int `json:"user_id"`
    
    }
    GetArr := []GetuserDetails{}
    
    var user_name string
 var email_id string
var date_of_birth string
   var community_id int
   var profilepic_url string
   var profilevideo_url string
   var login_type int
   var account_type int
   var created_on string
   var updated_on string
    /*----------------------------------*/
    found = false
    for rows.Next() {
        rows.Scan(&user_name,&email_id,&date_of_birth,&community_id,&profilepic_url,&profilevideo_url,&login_type,&account_type,&created_on,&updated_on)
)
        GetArr=append(GetArr,GetuserDetails{user_name,email_id,date_of_birth,community_id,profilepic_url,profilevideo_url,login_type,account_type,created_on,updated_on)})
        found = true
    }
    if found==false{
        fmt.Println("No user found")
        views.ShowSuccessOrErrorAsJSON(rw,"NoUserFound","There is no any user",130,30)
        return
    }
    fmt.Println("Successfully got the users list")
    fmt.Fprintf(rw,views.UniversalMessageCreatorAsJSON([]string{"success","message","status"},1,"user details",1));
}

