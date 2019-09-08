package category

import "time"

//category
type Category struct {
	Id 				int `json:"id"`
	Name 			string `json:"name"`
	Introduction	string `json:"introduction"`
	Feature			string `json:"feature"`
	CreatedAt  		time.Time `json:"created_at"`
	UpdatedAt 		time.Time `json:"updated_at"`
}

//category response
type CategoryResponse struct {
	Id 				string `json:"id"`
	Name 			string `json:"name"`
	Introduction	string `json:"introduction"`
	Feature			string `json:"feature"`
	CreatedAt  		time.Time `json:"created_at"`
}