package common

import (
	"net/http"
	"io/ioutil"
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-services/global/common/third/qiniu"
	"go-services/global/common/utils/io/o"
	"go-services/global/common/utils/time/time"
	"go-services/global/common/utils/validation/v"
	"go-services/app/http/pets/service/v1/identification"
	"go-services/global/common/third/baidu/Identification/animal"
	identification2 "go-services/app/http/pets/model/identification"
	"go-services/global/common/server/app"
	"go-services/global/common/utils/io"
	"go-services/app/http/pets/service/v1/category"
)

// @Summary 上传宠物图片
// @Accept json
// @Tags 宠物接口
// @Security Bearer
// @Produce  json
// @Param image query file true "图片"
// @Resource Name
// @Router /pet/common/upload [POST]
// @Success 200 {object} pet.PetResponse
func Upload(c *gin.Context) {
	f, h, err := c.Request.FormFile("image")
	if err != nil {
		io.App.Response(o.C_ERROR, o.M_ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
		return
	}
	defer f.Close()
	file, _ := h.Open() // 这里获得的实际就是一个io,通过源码看到这个open方法最终返回的是一个结构体,其内部包含了 io.Reader的接口
	if filename, err := qiniu.UploadImage(file, h.Size); err != nil { // 通过h.size 即可获得文件大小
		io.App.Response(o.C_ERROR, o.M_ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		return
	} else {
		service := identification.IdentificationService{}
		service.Identification.UserId 		= app.User.Id
		service.Identification.Image		= filename
		service.Identification.CreatedAt	= time.CurrentTimeNow()
		service.Identification.UpdatedAt	= time.CurrentTimeNow()
		if err := service.Create();err != nil {
			io.App.Response(o.C_ERROR,o.M_ERROR_UPLOAD_SAVE_IMAGE_FAIL,nil)
			return
		}
		response := identification2.UploadIdentificationResponse{K:filename,Path: qiniu.Driver.DefaultUrl+filename}
		io.App.Response(o.C_SUCCESS,o.M_SUCCESS,response)
	}
}


// @Summary 识别宠物
// @Accept json
// @Tags 宠物接口
// @Security Bearer
// @Produce  json
// @Param k query string true "upload 接口返回的 K"
// @Resource Name
// @Router /pet/common/identification [POST]
// @Success 200 {object} pet.PetResponse
func Identification(c *gin.Context) {
	k := c.PostForm("k")
	if !v.V(struct {K string `valid:"Required"`}{K:k}) {
		io.App.Response(o.C_ERROR,o.M_INVALID_PARAMS,nil)
		return
	}
	fileUrl := qiniu.Driver.DefaultUrl+k
	res,err := http.Get(fileUrl)
	if err != nil {
		io.App.Response(o.C_ERROR,o.M_ERROR_UPLOAD_CHECK_IMAGE_FAIL,nil);return
	}
	imageByte,err := ioutil.ReadAll(res.Body)
	if err != nil {
		io.App.Response(o.C_ERROR,o.M_READ_FILE_FAIL,nil);return
	}
	if err := animal.GetAccessToken();err != nil {
		io.App.Response(o.C_ERROR,o.M_SERVER_ERROR,nil);return
	}
	if err := animal.Distinguish(base64.StdEncoding.EncodeToString(imageByte));err != nil {
		io.App.Response(o.C_ERROR,o.M_ERROR_IMAGE_DISTINGUISH_FAIL,nil);return
	}
	result := animal.DistinguishDriver.Result
	firstName := result[0].Name
	if firstName == "非动物" {
		io.App.Response(o.C_ERROR,o.M_IS_NOT_ANIMAL,nil)
		return
	}

	//创建分类 todo 功能完成之后使用消息队列
	categoryTimeNow := time.CurrentTimeNow()
	for _,v := range result {
		categoryService := category.CategoryService{}
		categoryService.Category.Name 		= v.Name
		categoryService.Category.CreatedAt 	= categoryTimeNow
		categoryService.Category.UpdatedAt 	= categoryTimeNow
		categoryService.Create()
	}

	last, _ 	:= json.Marshal(result[0])
	list, _ 	:= json.Marshal(result)
	attributes 	:= map[string]interface{}{"image" : k,"user_id" : app.User.Id}
	values		:= map[string]interface{}{
		"last" 			: last,
		"result_list"	: list,
	}
	if err := identification.SetAttributes(attributes,values); err != nil {
		io.App.Response(o.C_ERROR,o.M_OPERATION_FAIL,nil)
		return
	}
	io.App.Response(o.C_SUCCESS,o.M_SUCCESS,gin.H{
		"list" : animal.DistinguishDriver.Result,
	})
}

