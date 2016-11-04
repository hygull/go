package main

import "fmt"

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

	rows2, e:= db.Query("select users.user_name ,users.profilepic_url,newtable.user_id,newtable.post_id,newtable.created_on,newtable.updated_on , " +
		"newtable.id, newtable.likecount  from users inner join ((select posts.user_id,posts.post_id, " +
		"notifications.created_on , notifications.id, notifications.updated_on,notifications.likecount from posts inner join notifications on " +
		"posts.post_id=notifications.post_id ) as newtable) on (users.user_id=newtable.user_id AND " +
		" users.user_id=" + userid + ");")
	if e!=nil{
		view.ShowSuccessOrErrorAsJSON(rw,"QueryExecutionError","Error in execution in query",127,27)
		return
	}

	NotificationsArr := []Likes{}
	for rows2.Next() {
		rows2.Scan(&uname, &p_pic, &usid, &postid, &c_at, &u_at, &id, &likecount)
		rows3,e:=db.Query("select count(*) from notifications where post_id="+postid+";")
		if e!=nil{
			views.ShowSuccessOrErrorAsJSON(rw,"QueryExecutionError","Erro in execution in query",127,27)
			return
		}
		rows3.Next();
		rows3.Scan(&ncount)
		NotificationsArr = append(NotificationsArr, Likes{id, postid, ncount, likecount, usid, uname, p_pic, c_at, u_at})
	}
	fmt.Fprintf(rw, views.UniversalMessageCreatorAsJSON([]string{"success","message","ntf_arr"},1,"List in notifications",NotificationsArr)
	fmt.Println("Notifications list successfully displayed")
}
