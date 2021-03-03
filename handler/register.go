package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"../models"
)

// User represents Temporary User on Joining
var User map[string]bool

func init() {
	User = make(map[string]bool)
}

// Join HandleFunc
func Join(w http.ResponseWriter, r *http.Request) {
	// create decoder to read the body of request (readcloser)
	log.Print("Join function call")

	decoder := json.NewDecoder(r.Body)
	user := models.User{}
	// bind to user struct
	if err := decoder.Decode(&user); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	// check if user can join (already exist username)
	if isDuplicateUser(user.Username) {
		w.Write([]byte("duplicate"))
		return
	}
	// set the username to make it exist
	print("hiu", user.Email)
	User[user.Username] = true
	// set the email to make it exist
	User[user.Email] = true

	w.Write([]byte("ok"))
	return
}

func isDuplicateUser(user string) bool {
	if _, ok := User[user]; ok {
		return true
	}
	return false
}
