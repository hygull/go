package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	"encoding/json"
	_ "strings"
)

func GetPosts(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	//AuthenticationKey := conf.AuthenticationKey
	ctx.SetContentType("application/json; charset=utf8")

	if !ctx.IsPost() { //Checking whether the data is being sent using POST method or not(If already implemented in Android then not required)
		ShowSuccessOrErrorAsJSON(ctx, "PostMethodNotFoundError", "The data is not being sent using POST method", 107, 7)
		fmt.Println("The data is not being sent using POST method")
		return
	}
	l := len(ctx.PostBody())
	if l == 0 { //If there's no post data then inform the user...len([]byte )==0
		ShowSuccessOrErrorAsJSON(ctx, "NoPostDataError", "You haven't sent the POST data", 104, 4)
		fmt.Println("You haven't send the POST data(as it is empty)")
		return
	}
	// fmt.Println("Length of PostBody() : ", l)

	// uId := strings.TrimSpace(string(ctx.FormValue("uid")))
	// uat := strings.TrimSpace(string(ctx.FormValue("uat")))
	// ty := strings.TrimSpace(string(ctx.FormValue("ty")))
	// cmtyId := strings.TrimSpace(string(ctx.FormValue("cmtyid")))

	/****************************************************************************************/
	// fmt.Println("Extracted URL Parameters : ", uId, uat, ty, cmtyId)
	// if !AreAllFieldsAreInValidForm(uId, uat, ty, cmtyId) {
	// 	ShowSuccessOrErrorAsJSON(ctx, "EmptyFieldOrKeyNameError", "You have to provide all fields (uid,uat,ty,cmtyid) except uid or check if  you have provided a wrong key uid", 103, 3)
	// 	return
	// }

	fmt.Println("Connecting to DB...")
	db, err := sql.Open("mysql", "root:admin@321@tcp(127.0.0.1:3306)/synkku_live")

	if err != nil {
		fmt.Println("Error in connection...")
		fmt.Println(err.Error())
		ShowSuccessOrErrorAsJSON(ctx, "DBConnectionError", "Error in connection with database", 104, 4)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error in connection test...")
		fmt.Println(err.Error())
		ShowSuccessOrErrorAsJSON(ctx, "DBConnectionTestError", "Error in connection with database", 112, 12)
		return
	}
	fmt.Println("Connected...")

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
	// compOp := ""

	// if ty == "1" {
	// 	compOp = ">"
	// } else {
	// 	compOp = "<"
	// }

	// queryString := "select posts.id, posts.user_id, auth_users.user_name,auth_users.profile_pic, " +
	// 	"posts.base_postid, posts.recent_postid,posts.recent_userid,posts.visibility_type, " +
	// 	"posts.community_id,posts.media_type,posts.text_message, posts.status,posts.post_type, " +
	// 	"posts.created_on, posts.updated_on from posts inner join auth_users on " +
	// 	"posts.user_id=auth_users.id where (posts.community_id=" + cmtyId + " and posts.updated_on" + compOp +
	// 	"'" + uat + "' AND posts.visibility_type=1);"
	queryString := "select posts.post_id, posts.user_id, users.first_name,users.profilepic_url, " +
		"posts.base_postid, posts.recent_postid,posts.recent_user_id,posts.visibility_type, " +
		"posts.community_id,posts.media_type,posts.text_message, posts.status,posts.post_type, " +
		"posts.created_on, posts.updated_on from posts inner join users on " +
		"posts.user_id=users.user_id limit 107;"

	fmt.Println(queryString)
	rows, err := db.Query(queryString)
	if err != nil {
		fmt.Println("Error in query execution")
		ShowSuccessOrErrorAsJSON(ctx, "QueryExecutionError", "Error in execution of query", 127, 27)
		return
	}
	found := false
	a := 0
	for rows.Next() {
		rows.Scan(&pid, &uid, &un, &p_pic, &bpid, &rpid, &ruid, &vbt, &cmtyid, &mt, &txt, &sts, &ptype, &c_at, &u_at)
		//RecentUsers=append(RecentUsers,un)
		// if pid == 108 {
		// 	fmt.Println("Txt:", txt)
		// 	txt = "This is mine..."
		// }
		PostArr = append(PostArr, PostData{pid, uid, un, p_pic, bpid, rpid, ruid, runame, vbt, cmtyid, mt, txt, sts, nlikes, ncmts, isliked, ptype, c_at, u_at})
		found = true
		a += 1
	}
	fmt.Println(">>>> a: ", a)
	if found == false {
		ShowSuccessOrErrorAsJSON(ctx, "PostsNotFound", "As of now, no one has posted here", 130, 30)
		fmt.Println("No posts found in this community")
		return
	}
	fmt.Fprintf(ctx, UniversalMessageCreatorAsJSON([]string{"success", "message", "p_arr"}, 1, "New posts are", PostArr))
	fmt.Println("Posts have been successfully displayed...")
}

/************* Construction of JSON messages for success/error ***********/

func ShowSuccessOrErrorAsJSON(ctx *fasthttp.RequestCtx, errType string, errDetails string, statusCode int, id int) {
	var jsonStrData []byte

	type ErrMsg struct {
		Message    string `json:"message"`
		Statuscode int    `json:"status code"`
		Id         int    `json:"id"`
	}
	type ErrMsgs struct {
		MessageDetails ErrMsg `json:"message details"`
	}
	msgsCollection := map[string]ErrMsgs{}
	msgsCollection[errType] = ErrMsgs{ErrMsg{errDetails, statusCode, id}}
	jsonStrData, _ = json.Marshal(msgsCollection)
	fmt.Fprintf(ctx, string(jsonStrData))
}

func UniversalMessageCreatorAsJSON(keys []string, values ...interface{}) string {
	jsonMapData := map[string]interface{}{}
	i := 0
	for _, v := range values {
		jsonMapData[keys[i]] = v
		i += 1
	}
	var jsonStrData []byte
	jsonStrData, _ = json.Marshal(jsonMapData)
	return string(jsonStrData)
}

// func AreAllFieldsAreInValidForm(postData ...string) bool {
// 	isAnyBlank := false
// 	for _, data := range postData {
// 		if data == "" {
// 			isAnyBlank = true
// 			break
// 		}
// 	}
// 	fmt.Println("Is Any Field blank : ", isAnyBlank)
// 	return !isAnyBlank //An efficient way to check for empty fields
// }

func main() {
	//To define a fasthttp router...
	fastRouter := fasthttprouter.New()
	fastRouter.POST("/v1/posts/all", GetPosts)
	fmt.Println("API is running....")
	fasthttp.ListenAndServe(":8080", fastRouter.Handler)

}
