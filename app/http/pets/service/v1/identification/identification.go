package identification

import (
	"go-services/app/http/pets/model/identification"
	"go-services/global/common/database/mysql"
)

type IdentificationService struct {
	Identification identification.Identification
}

//创建
func(i *IdentificationService) Create() error {
	return mysql.DB.Create(&i.Identification).Error
}

//edit
func SetAttributes(attributes interface{},values interface{}) error {
	var identificationModel identification.Identification
	return mysql.DB.Model(&identificationModel).Where(attributes).Update(values).Error
}