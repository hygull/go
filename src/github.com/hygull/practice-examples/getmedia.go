package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	//"strconv"
)

func checkEmptyValues(params ...string) bool {
	isEmpty := false
	for _, str := range params {
		if str == "" {
			isEmpty = true
			break //terminate loop immediately
		}
	}
	return isEmpty
}
func GetMediaFiles(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")

	req.ParseForm()
	if len(req.Form) == 0 {
		fmt.Println("empty fields")
		return
	}
	fmt.Println("Hiii")

	post_id := req.Form["p_id"][0]
	user_id := req.Form["u_id"][0]

	fmt.Println(post_id, user_id)
	isEmpty := checkEmptyValues(user_id, post_id)
	fmt.Println("Hi")

	if isEmpty == true {
		fmt.Println("Empty value") //JSON message into browser
	} else { //Start of else
		fmt.Println("Connecting to MySQL...")
		db, err := sql.Open("mysql", "root:jeevitha@tcp(127.0.0.1:3306)/TiiingDb")
		if err != nil {
			fmt.Println("This is the error : ", err.Error())
		}
		defer db.Close()

		err = db.Ping()
		if err != nil {
			fmt.Println("Ping Error : ", err.Error())
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

			//var Username string
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

			for rows2.Next() {

				rows2.Scan(&media_id, &media_type, &media_url, &status, &created_on, &updated_on)
				fmt.Println(media_id, media_type, media_url, status, created_on, updated_on)
				MediaArr = append(MediaArr, MediaData{media_id, media_type, media_url, status, created_on, updated_on})
			}
			fmt.Fprintf(rw, UniversalMessageCreatorAsJSON([]string{"success", "message", "me_arr"}, 1, "Media files are here", MediaArr))
		} else {
			//fmt.Println("User not found")
			errorMap := map[string]string{}
			errorMap["success"] = "0"
			errorMap["message"] = "media not found"
			jsonStrData2 := []byte{}
			jsonStrData2, e = json.Marshal(errorMap)
			fmt.Fprintf(rw, string(jsonStrData2))
		}
	}
}

func handleRequests() {
	newRouter := mux.NewRouter().StrictSlash(true)
	//newRouter.HandleFunc("/", homePage)
	newRouter.HandleFunc("/tiiing/feeds/GetMediaFiles", GetMediaFiles)

	if len(os.Args) != 2 {
		fmt.Println("You forgot to give the port number")
		return
	}
	log.Fatal(http.ListenAndServe(":"+os.Args[1], newRouter))
}

func main() {

	fmt.Println("welcome to my synkku database")
	fmt.Println("Getting mediaurl from database")

	handleRequests()
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
