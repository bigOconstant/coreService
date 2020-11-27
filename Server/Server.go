package server

import (
	"bytes"
	"coreService/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	HOMEPAGE = `
	
	<table style="width:100%">
		<tr>
			<th>Route</th>
			<th>Verb</th>
			<th>Pupose</th>
		</tr>
		<tr>
			<td>/login/</td>
			<td>Post/</td>
			<td>Takes user name and password, returns a token</td>
		</tr>
	</table>

	`
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

	url := os.Getenv("USERSERVICEURL") + "/login/"

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))

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

func (m *Server) GetHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, HOMEPAGE)
}

func (m *Server) getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, HOMEPAGE)

}

func (m *Server) initServer() {
	fmt.Println("Initilizing Server")
	m.Router = mux.NewRouter()
}

func (m *Server) Serve() {
	m.initServer()
	m.Router.HandleFunc("/", m.GetHome).Methods("GET")
	m.Router.HandleFunc("/login", m.login).Methods("POST")
	log.Fatal(http.ListenAndServe(":3001", m.Router))

}
