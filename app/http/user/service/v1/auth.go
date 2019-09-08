package v1

import (
	"go-services/app/http/user/model/user"
	"go-services/global/common/database/mysql"
)

type AuthService struct {
	UserModel user.User
}

//注册
func (user *AuthService) Register() error {
	return mysql.DB.Create(&user.UserModel).Error
}

//自定义修改
func SetAttributes(attributes interface{},values interface{}) error {
	var userModel user.User
	return mysql.DB.Model(&userModel).Where(attributes).Update(values).Error
}

//自定义查询
func GetAttributes(attributes interface{}) *user.User {
	var userModel user.User
	mysql.DB.Where(attributes).First(&userModel)
	return &userModel
}
