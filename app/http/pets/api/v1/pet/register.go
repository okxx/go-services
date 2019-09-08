package pet

import (
	"github.com/gin-gonic/gin"
	"go-services/global/common/utils/io"
	"go-services/global/common/utils/validation/v"
	"go-services/global/common/utils/io/o"
	"go-services/app/http/pets/service/v1/pet"
	"go-services/global/common/server/app"
	"go-services/global/common/utils/time/time"
	pet2 "go-services/app/http/pets/model/pet"
	"strconv"
)

// @Summary 注册宠物
// @Accept json
// @Tags 宠物接口
// @Security Bearer
// @Produce  json
// @Param cid query string true "category id"
// @Param nickname query string true "宠物名称"
// @Param avatar query string true "宠物头像"
// @Resource Name
// @Router /pet/pet/register [POST]
// @Success 200 {object} pet.PetResponse
func Register(c *gin.Context) {
	cid 		:= io.App.Context.Query("cid")
	nickname 	:= io.App.Context.Query("nickname")
	avatar		:= io.App.Context.Query("avatar")
	if !v.V(struct {Cid string `valid:"Required"`;Nickname string `valid:"Required"`;Avatar string `valid:"Required"`}{Cid:cid,Nickname:nickname,Avatar:avatar}) {
		io.App.Response(o.C_ERROR,o.M_INVALID_PARAMS,nil)
		return
	}

	attributes := map[string]interface{}{"nickname" : nickname, "user_id" : app.User.Id}
	petModel := pet.GetAttributes(attributes)
	if petModel.Id > 0 {
		io.App.Response(o.C_ERROR,o.M_PET_EXISTED,nil)
		return
	}

	categoryId,_ := strconv.Atoi(cid)
	service := pet.PetService{}
	service.Pet.Cid			= categoryId
	service.Pet.UserId		= app.User.Id
	service.Pet.Nickname 	= nickname
	service.Pet.Avatar		= avatar
	service.Pet.CreatedAt	= time.CurrentTimeNow()
	service.Pet.UpdatedAt	= time.CurrentTimeNow()
	if err := service.Create(); err != nil {
		io.App.Response(o.C_ERROR,o.M_REGISTER_PET_ERROR,nil)
		return
	}

	response := pet2.PetResponse{
		Id			: service.Pet.Id,
		Cid			: categoryId,
		UserId		: app.User.Id,
		Nickname	: nickname,
		Avatar		: avatar,
	}
	io.App.Response(o.C_SUCCESS,o.M_SUCCESS,response)
}