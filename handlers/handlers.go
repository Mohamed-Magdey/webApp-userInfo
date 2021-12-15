package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"webApp/userInfo/model"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

//Index page handler
func IndexFunc(w http.ResponseWriter, r *http.Request) {
	au := model.ShowAllUsers()
	t, err := template.ParseFiles("templates/indexPage.html")
	checkError(err)
	t.Execute(w, au)
}

//handler to show user with id input
func ShowUserFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/showUserPage.html")
		t.Execute(w, nil)

	} else {

		id, err := strconv.Atoi(r.FormValue("id"))
		checkError(err)
		var alUsrs model.AllUsers
		file, err := os.OpenFile("list.json", os.O_RDONLY, 0666)
		checkError(err)
		b, err := ioutil.ReadAll(file)
		checkError(err)
		json.Unmarshal(b, &alUsrs.Users)

		var allID []int
		for _, usr := range alUsrs.Users {
			allID = append(allID, usr.Id)
		}
		for _, usr := range alUsrs.Users {
			if !model.IsValueInSlice(allID, id) {
				http.Redirect(w, r, "/showuser/notsuccededshow/", http.StatusFound)
				return
			}
			if usr.Id != id {
				continue
			} else {
				t, err := template.ParseFiles("templates/showUserPage.html")
				checkError(err)
				t.Execute(w, usr)
			}

		}
	}
}

func ShowUser(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/showUser.html")
}

func NotSuccededShow(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/notSuccededShow.html")
}
