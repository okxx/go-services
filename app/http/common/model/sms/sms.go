package sms

import "time"

const (
	//表名
	_smsTableName 				= "tbl_sms"
)


// sms table schema
type Sms struct {
	Id 			int `json:"id"`
	Status 		int `json:"status"`
	Account 	string `json:"account"`
	Mode 		string `json:"mode"`
	Sms			string `json:"sms"`
	Effective	int    `json:"effective"`
	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
}

//sms response struct
type SmsInfoRes struct {
	Id 			int `json:"id"`
	Status 		int `json:"status"`
	Account 	string `json:"account"`
	Mode 		string `json:"mode"`
	Sms			string `json:"sms"`
	Effective	string `json:"effective"`
	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
}

//发送验证码返回值
type SmsSendResponse struct {
	Sms 		string	`json:"sms"`
	Hint 		string	`json:"hint"`
}

func (m *SmsSendResponse) Serializer() SmsSendResponse{
	return SmsSendResponse{
		Sms: m.Sms,
		Hint: m.Hint,
	}
}

func (sms *Sms) TableName() string {
	return _smsTableName
}
