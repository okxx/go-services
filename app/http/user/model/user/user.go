package user

import "time"

//user table schema
type User struct {
	Id 			int  		`json:"id"`
	Account		string  	`json:"account"`
	Password 	string  	`json:"password"`
	Salting		string  	`json:"salting"`
	Ticket 		string  	`json:"ticket"`
	Nickname 	string  	`json:"nickname"`
	Status 		int     	`json:"status"`
	CreatedAt	time.Time 	`json:"created_at"`
	UpdatedAt  	time.Time	`json:"updated_at"`
}

//用户公共返回值
type UserResponse struct {
	Id 			int  		`json:"id"`
	Account		string  	`json:"account"`
	Nickname 	string  	`json:"nickname"`
	Status 		int     	`json:"status"`
}

//user login response
type UserLoginResponse struct {
	Account		string  	`json:"account"`
	Ticket 		string  	`json:"ticket"`
	Token  		string		`json:"token"`
}