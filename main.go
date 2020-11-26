package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	USERSERVICEURL = "jetsonnano"
)

type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	guy := user{"caleb", "caleb@gmail.com"}
	json.NewEncoder(w).Encode(&guy)

}

//Logintoken token
type LoginToken struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

//Login struct
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func login(w http.ResponseWriter, r *http.Request) {

	var output Login
	err := json.NewDecoder(r.Body).Decode(&output)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(output)

	requestBody, err := json.Marshal(map[string]string{
		"username": output.Username,
		"password": output.Password,
	})

	resp, err := http.Post(USERSERVICEURL+"/login/", "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var token LoginToken = LoginToken{}

	jsonInput, err := strconv.Unquote(string(body))
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(&token)
	}

	err = json.Unmarshal([]byte(jsonInput), &token)

	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(&token)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&token)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", getUser).Methods("GET")
	r.HandleFunc("/login", login).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", r))

}
