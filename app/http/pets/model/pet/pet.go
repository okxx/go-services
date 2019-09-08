package pet

import "time"

type Pet struct {
	Id 				int `json:"id"`
	Cid 			int `json:"cid"`//category id [种类]
	UserId 			int `json:"user_id"`// 用户id
	Nickname		string `json:"nickname"`
	Avatar			string `json:"avatar"`
	CreatedAt 		time.Time `json:"created_at"`
	UpdatedAt 		time.Time `json:"updated_at"`
}

//宠物公共返回值
type PetResponse struct {
	Id 				int `json:"id"`
	Cid 			int `json:"cid"`//category id [种类]
	UserId 			int `json:"user_id"`// 用户id
	Nickname		string `json:"nickname"`
	Avatar			string `json:"avatar"`
}