func(){
	
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

	if !areAllFieldsAreInValidForm(usId, uAt, ty, comId) {
			views.ShowSuccessOrErrorAsJSON(rw, "EmptyFieldError", "You have to select all the fileds", 103, 3)
			return
	}

	userId:=req.Form["uid"]
	uAt:=req["u_at"]
	ty:=req.Form["ty"]
	cmytyId:=req.Form["cmtyid"]

	fmt.Println("Connecting to DB...")
	db,err:=sql.Open("mysql","root:admin@321@tcp(127.0.0.1:3306)/synkku?charset=utf8");
	if err!=nil{
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


type PostData struct{
  Pid int `json:"pid"`
  Uid int    `json:"uid"`
  Un string `json:"un"`
  Ppic string `json:"p_pic"`
  Bpid int `jspn:"bpid"`
  Rpid int `json:"rpid"`
  Ruid int `json:"ruid"`
  Runame string `json:"runame"`
  Vbt int `json:"vbt"`
  Cmtyid int `json:"cmtyid"`
  Mt int `json:"mt"`
  Txt int `json:"txt"`
  Sts int `json:"sts"`
  Nlikes int `json:"nlikes"`
  Ncmts int `json:"ncmts"`
  Isliked int `json:"isliked"`
  Ptype int `json:"p_type"`
  Cat string `json:"c_at"`
  Uat string `json:u_at`
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
  var txt int
  var sts int 
  var nlikes int //Remaining
  var ncmts int  //Remaining
  var isliked int //Remaining
  var ptype int 
  var c_at string 
  var  u_at string 
  
  PostArr:=[]PostData{}
  RecentUsers:=[]string{}
  compOp:=""

  if ty=="1"{
  	compOp=">"
  }else{
  	compOp="<"
  }

  queryString="select posts.post_id, posts.user_id, users.user_name,users.profilepic_url, "+
  "posts.base_postid, posts.recent_postid,posts.recent_user_id,posts.visibility_type, "+
  "posts.community_id,posts.media_type,posts.text_message, posts.status,posts.post_type, "+
  "posts.created_on, posts.updated_on from posts inner join users on "+
  "posts.recent_user_id=users.user_id where (posts.community_id="+cmtyid+"and posts.updated_on"+compOp+
  "'"+uAt+"' AND posts.visibility_type=1) order by posts.updated_on desc limit 50;"
  
  rows,err:=db.Query(queryString);
  found:=false
  for rows.Next(){
	rows.Scan(&pid,,&uid,&un,&p_pic,&bpid,&rpid,&ruid,&vbt,&cmtyid,&mt,&txt,&sts,&ptype,&c_at,&u_at);
	RecentUsers=append(RecentUsers,un)
  	PostArr=append(PostArr,PostData{pid,uid,un,p_pic,bpid,rpid,ruid,"",vbt,cmtyid,mt,txt,sts,0,0,0,ptype,c_at,u_at})
  	found=true;
  }

  if found==false{
  	views.ShowSuccessOrErrorAsJSON(rw,"PostsNotFound","As of now, no one has posted here",130,30);
  	fmt.Println("No posts found in this community");
	return
  }
  fmt.Fprintf(rw,views.UniversalMessageCreatorAsJSON([]string{"success","message","p_arr"},1,"New posts are",PostArr))
  fmt.Println("Posts has been successfully displayed")		
}
