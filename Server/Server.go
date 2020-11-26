package server

import (
	"bytes"
	"coreService/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	USERSERVICEURL = "http://jetsonnano.local"

)
type Server struct {
	Router *mux.Router
}

func (m *Server) login(w http.ResponseWriter, r *http.Request) {

	var output models.Login
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
	var token models.LoginToken = models.LoginToken{}

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

func (m *Server) getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	guy := models.User{"welcome", "to the api"}
	json.NewEncoder(w).Encode(&guy)

}

func (m *Server) initServer(){
	fmt.Println("Initilizing Server")
	m.Router = mux.NewRouter()
}


func (m *Server) Serve(){
	m.initServer()
	m.Router.HandleFunc("/", m.getUser).Methods("GET")
	m.Router.HandleFunc("/login", m.login).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", m.Router))
	
}
