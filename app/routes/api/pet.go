package api

import (
	"go-services/global/common/server/app"
	"go-services/app/http/pets/api/v1/common"
	"go-services/app/http/pets/api/v1/pet"
	"go-services/app/middleware/auth"
)

//pet interface
func PetInterface() {

	p := app.Router.Group("pet")
	{
		p.Use(auth.Authentification)

		pcm := p.Group("common")
		{
			pcm.Any("upload",common.Upload)
			pcm.Any("identification",common.Identification)
		}

		pi := p.Group("pet")
		{
			pi.Any("register",pet.Register)
			pi.Any("edit",pet.Edit)
		}

	}

}