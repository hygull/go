  package main

  import("database/sql"
  _ "github.com/go-sql-driver/mysql"
  "net/http"
  _"fmt"
  )

func ListPosts(rw http.ResponsWriter, req *http.Request){
  db,err:=sql.Open("mysql","root:admin@321@tcp(127.0.0.1:3306)/synkku?charset=utf8");
  if err!=nil{
    fmt.Println("DB Exit")
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
  var runame string 
  var vbt int
  var cmtyid int
  var mt int 
  var txt int
  var sts int 
  var nlikes int 
  var ncmts int 
  var isliked int
  var ptype int 
  var c_at string 
  var  u_at string 

  userId:=req.Form["uid"]
  uAt:=req["u_at"]
  ty:=req.Form["ty"]
  cmytyid:=req.Form["cmtyid"]

  rows,err:=db.Query("select posts.post_id, posts.base_postid, posts.recent_post_id,posts.recent_user_id,users.user_name,posts.visibility_type,posts.community_id,posts.text_message from posts inner join users on posts.recent_user_id=users.user_id")

}


func main(){
  http.HandleFunc("/v1/feeds/getpostslist",ListPosts);
  http.ListenAndServe(":9000",nil);
}
/*
select posts.post_id, posts.base_postid, posts.recent_postid,posts.recent_user_id,users.user_name,posts.visibilitposts.community_id,posts.text_message from posts inner join users on posts.recent_user_id=users.user_id where posts.community_id=3 and posts.updated_on<'2016-10-11';

select posts.post_id, posts.base_postid, posts.recent_postid,posts.recent_user_id,users.user_name,posts.visibility_type,posts.community_id,posts.text_message, posts.status,posts.created_on, posts.updated_on from posts inner join users on posts.recent_user_id=users.user_id where (posts.community_id=3 and posts.updated_on<'2016-09-30') order by posts.updated_on desc;

*/

