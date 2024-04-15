package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = make(map[string]User)

func addnewuser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var entry User

	err := json.NewDecoder(r.Body).Decode(&entry)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	users[entry.Name] = entry

	w.WriteHeader(http.StatusCreated)
	return
}

func main() {
	http.HandleFunc("/createuser", addnewuser)
	fmt.Println("serve start")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
