package handlers

import (
	"fmt"
	"html/template"
	"net/http"
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
