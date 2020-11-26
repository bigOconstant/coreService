package models

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
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
