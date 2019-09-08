package common

import (
	"go-services/global/common/server/request"
	"go-services/global/common/server/app"
)

func getPetList() {


	rpc := request.Rpc{}
	rpc.Services = map[string]interface{}{
		"pet.common.getList" : map[string]interface{}{"userId" : app.User.Id},
	}
	rpc.Sync()
}