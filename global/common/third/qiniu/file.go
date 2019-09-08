package qiniu

import (
	"github.com/qiniu/api.v7/storage"
	"github.com/qiniu/api.v7/auth/qbox"
	"context"
	"io"
	"strconv"
	"encoding/base64"
	"fmt"
	"go-services/global/common/utils/strings"
	"go-services/global/common/utils/time/time"
	"go-services/global/common/utils/auth/md5"
)

type qiniu struct {
	AccessKey 				string
	SecretKey 				string
	Bucket 					string
	Access 					string
	HotlinkPreventionKey 	string
	DefaultUrl				string
}
var Driver = &qiniu{}

func config() storage.Config {
	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	return cfg
}

func UploadImage(localFile io.Reader, size int64) (string, error) {
	firstName 	:= md5.EncodeMD5(strings.Random(5,25) + strconv.Itoa(time.CurrentTime()))
	lastName 	:= md5.EncodeMD5(strings.Random(6,25) + strings.Random(6,25))
	filename 	:= firstName+lastName// 这里是计算了一个md5(time())的字符串作为文件名
	putPolicy := storage.PutPolicy{
		Scope: Driver.Bucket,
	}
	mac := qbox.NewMac(Driver.AccessKey, Driver.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cig := config()
	formUploader := storage.NewFormUploader(&cig)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}
	err := formUploader.Put(context.Background(), &ret, upToken, filename, localFile, size, &putExtra)
	if err != nil {
		return "", err
	}
	return ret.Key, nil
}

// EncodedEntry 生成URL Safe Base64编码的 Entry
func EncodedEntry(key string) string {
	entry := fmt.Sprintf("%s:%s", Driver.Bucket, key)
	return base64.URLEncoding.EncodeToString([]byte(entry))
}
