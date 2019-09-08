package animal

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
	"go-services/global/common/cache/redis"
)

type baidu struct {
	ApplicationName			string
	AppId 					string
	AppKey					string
	AppSecret 				string
}
var Driver = &baidu{}

const (
	TOKEN_URL  					= "https://aip.baidubce.com/oauth/2.0/token"// GET TOKEN API
	DISTINGUISH_URL  	= "https://aip.baidubce.com/rest/2.0/image-classify/v1/animal"
	ACCESS_TOKEN  		= "baidu_ai_access_token"// REDIS TOKEN KEYS
)

type accessToken struct {
	RefreshToken 		string `json:"refresh_token"`
	ExpiresIn			int    `json:"expires_in"`
	SessionKey			string `json:"session_key"`
	AccessToken			string `json:"access_token"`
	Scope				string `json:"scope"`
	SessionSecret		string `json:"session_secret"`
}
var AccessTokenDriver = &accessToken{}

type distinguish struct {
	LogId	int						`json:"log_id"`
	Result	[]struct{
		Score 			string	`json:"score"`
		Name			string	`json:"name"`
		BaikeInfo		struct 	{
			BaikeUrl	string	`json:"baike_url"`
			ImageUrl	string	`json:"image_url"`
			Description string	`json:"description"`
		} `json:"baike_info"`
	} `json:"result"`
}
var DistinguishDriver = &distinguish{}


//return access token
func GetAccessToken() (error) {
	conn := redis.Server.Get()
	defer conn.Close()
	exists,_ := redis.Bool(conn.Do("EXISTS",ACCESS_TOKEN))
	if !exists {
		//todo 获取access token
		resp, err := http.PostForm(TOKEN_URL, url.Values{
			"grant_type": {"client_credentials"},// 必须参数，固定为client_credentials；
			"client_id": {Driver.AppKey},// 必须参数，应用的API Key；
			"client_secret": {Driver.AppSecret},// 必须参数，应用的Secret Key；
		})
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil
		}
		jsonErr := json.Unmarshal(body,AccessTokenDriver)
		if jsonErr != nil {
			return jsonErr
		}

		if _,err := redis.Bool(conn.Do("SET",ACCESS_TOKEN,AccessTokenDriver,AccessTokenDriver.ExpiresIn)); err != nil {
			return err
		}
	}
	value,err := redis.Bytes(conn.Do("GET",ACCESS_TOKEN))
	if err != nil {
		return err
	}
	resultErr := json.Unmarshal(value,AccessTokenDriver)
	if resultErr != nil {
		return resultErr
	}
	return nil
}


//todo 识别
func Distinguish(imageBase64 string) (error) {
	resp, err := http.PostForm(DISTINGUISH_URL + "?access_token=" + AccessTokenDriver.AccessToken, url.Values{
		"image": {imageBase64},//图像数据，base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px,支持jpg/png/bmp格式。注意：图片需要base64编码、去掉编码头后再进行urlencode。
		"top_num": {"6"},//返回预测得分top结果数，默认为6
		"baike_num": {"1"},//返回百科信息的结果数，默认不返回
	})
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body,&DistinguishDriver);err != nil {
		return err
	}
	return nil
}