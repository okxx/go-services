package pet

import (
	"github.com/gin-gonic/gin"
	"go-services/global/common/utils/io"
	"go-services/global/common/utils/validation/v"
	"go-services/global/common/utils/io/o"
	"go-services/app/http/pets/service/v1/pet"
	"go-services/global/common/server/app"
	pet2 "go-services/app/http/pets/model/pet"
	"go-services/global/common/utils/time/time"
)

// @Summary 编辑宠物
// @Accept json
// @Tags 宠物接口
// @Security Bearer
// @Produce  json
// @Param nickname query string true "宠物名称"
// @Param avatar query string true "宠物头像"
// @Resource Name
// @Router /pet/pet/edit [POST]
// @Success 200 {object} pet.PetResponse
func Edit(c *gin.Context) {
	id 			:= io.App.Context.Query("id")
	nickname 	:= io.App.Context.Query("nickname")
	avatar		:= io.App.Context.Query("avatar")
	if !v.V(struct {Nickname string `valid:"Required"`;Avatar string `valid:"Required"`}{Nickname:nickname,Avatar:avatar}) {
		io.App.Response(o.C_ERROR,o.M_INVALID_PARAMS,nil)
		return
	}

	//检查宠物是否存在
	petModel := pet.GetAttributes(map[string]interface{}{"id" : id, "user_id" : app.User.Id})
	if petModel.Id < 1 {
		io.App.Response(o.C_ERROR,o.M_PET_NOT_EXIST,nil)
		return
	}

	//检查重名
	petModel = pet.GetAttributes(map[string]interface{}{"nickname" : nickname, "user_id" : app.User.Id})
	if petModel.Id > 0 {
		io.App.Response(o.C_ERROR,o.M_PET_EXISTED,nil)
		return
	}

	attributes 	:= map[string]interface{}{"id" : id, "user_id" : app.User.Id}
	values 		:= map[string]interface{}{"nickname" : nickname,"avatar":avatar,"updated_at": time.CurrentTimeNow()}
	if err := pet.Update(attributes,values); err != nil {
		io.App.Response(o.C_ERROR,o.M_EDIT_PET_ERROR,nil)
		return
	}

	response := pet2.PetResponse{
		Id			: petModel.Id,
		Cid			: petModel.Cid,
		UserId		: app.User.Id,
		Nickname	: nickname,
		Avatar		: avatar,
	}
	io.App.Response(o.C_SUCCESS,o.M_SUCCESS,response)
}