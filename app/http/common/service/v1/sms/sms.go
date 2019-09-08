package sms

import (
	"go-services/app/http/common/model/sms"
	"go-services/global/common/database/mysql"
)

type SmsService struct {
	Model sms.Sms
}

func (s *SmsService) Create() error {
	return mysql.DB.Create(&s.Model).Error
}

func Attributes(attributes map[string]interface{}) *sms.Sms {
	var smsModel sms.Sms
	mysql.DB.Where(attributes).First(&smsModel)
	return &smsModel
}

func Update(attributes map[string]interface{},values map[string]interface{}) error {
	return mysql.DB.Model(&sms.Sms{}).Where(attributes).Update(values).Error
}