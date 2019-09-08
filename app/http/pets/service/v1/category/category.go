package category

import (
	"go-services/app/http/pets/model/category"
	"go-services/global/common/database/mysql"
)

type CategoryService struct {
	Category category.Category
}

//create category
func (c *CategoryService) Create() {
	mysql.DB.Create(&c.Category)
}