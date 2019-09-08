package pet

import (
	"go-services/app/http/pets/model/pet"
	"go-services/global/common/database/mysql"
)

type PetService struct {
	Pet pet.Pet
}

//get attributes
func GetAttributes(attributes interface{}) *pet.Pet {
	var petModel pet.Pet
	mysql.DB.Where(attributes).First(&petModel)
	return &petModel
}

//create
func (p *PetService) Create() error {
	return mysql.DB.Create(&p.Pet).Error
}

//update
func Update(attributes interface{},values interface{}) error {
	var petModel pet.Pet
	return mysql.DB.Model(&petModel).Where(attributes).Update(values).Error
}