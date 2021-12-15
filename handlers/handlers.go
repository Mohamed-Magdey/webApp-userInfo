package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
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
	fmt.Println(r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/showUserPage.html")
		t.Execute(w, nil)

	} else {

		id, err := strconv.Atoi(r.FormValue("id"))
		checkError(err)
		var alUsrs model.AllUsers
		file, err := os.OpenFile("list.json", os.O_RDONLY, 0666)
		checkError(err)
		defer file.Close()
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

//function to check correct user adding input (regular expression and non-empty field input)
func checkFormValue(w http.ResponseWriter, r *http.Request, forms ...string) (res bool, errStr string) {
	for _, form := range forms {
		reg, _ := regexp.Compile("^[a-zA-Z]+$")
		m := reg.MatchString(r.FormValue(form))

		if r.FormValue(form) == "" {
			return false, "All forms must be completed"
		}
		if !m {
			return false, "Use only english letters if firstname,lastname forms"
		}

	}
	return true, ""
}

//function to add user
func AddNewUserFunc(w http.ResponseWriter, r *http.Request) {

	//creating new instance and checking method
	newUser := &model.User{}
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/addNewUser.html")
		t.Execute(w, nil)

	} else {
		resBool, errStr := checkFormValue(w, r, "firstname", "lastname")
		if !resBool {
			t, err := template.ParseFiles("templates/notSucceded.html")
			checkError(err)
			t.Execute(w, errStr)

			return
		}
		newUser.FirstName = r.FormValue("firstname")
		newUser.LastName = r.FormValue("lastname")
		var err error
		newUser.Balance, err = strconv.ParseFloat(r.FormValue("balance"), 64)
		checkError(err)

		//open file
		file, err := os.OpenFile("list.json", os.O_RDWR, 0644)
		checkError(err)
		defer file.Close()

		//read file and unmarshall json file to slice of users
		b, err := ioutil.ReadAll(file)
		checkError(err)
		var alUsrs model.AllUsers
		err = json.Unmarshal(b, &alUsrs.Users)
		checkError(err)
		max := 0

		//generation of id(last id at the json file+1)
		for _, usr := range alUsrs.Users {
			if usr.Id > max {
				max = usr.Id
			}
		}
		id := max + 1
		newUser.Id = id

		//appending newUser to slice of all Users and rewrite json file
		alUsrs.Users = append(alUsrs.Users, newUser)
		newUserBytes, err := json.MarshalIndent(&alUsrs.Users, "", " ")
		checkError(err)
		ioutil.WriteFile("list.json", newUserBytes, 0666)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)

	}
}

//serving file with error (add function:empty field input or uncorrect input)
func NotSucceded(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/notSucceded.html")
}
