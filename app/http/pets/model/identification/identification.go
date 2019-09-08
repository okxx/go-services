package identification

import "time"

//识别
type Identification struct {
	Id			string	`json:"id"`
	UserId		int		`json:"user_id"`
	Image 		string	`json:"image"`
	Last 		string	`json:"last"`
	ResultList	string	`json:"result_list"`
	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt	time.Time `json:"updated_at"`
}

//识别同意返回结构
type IdentificationResponse struct {
	Id			string	`json:"id"`
	UserId		int		`json:"user_id"`
	Image 		string	`json:"image"`
	Last 		string	`json:"last"`
	CreatedAt 	time.Time `json:"created_at"`
}

//上传识别图片的返回值
type UploadIdentificationResponse struct {
	K 		string `json:"image"`
	Path 	string `json:"path"`
}