package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

func login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := json.Marshal(map[string]string{
		"username": "caleb",
		"password": "password",
	})

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("http://user-service/login/", "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&body)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", getUser).Methods("GET")
	r.HandleFunc("/login", login).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", r))

}
