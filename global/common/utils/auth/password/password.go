package password

import (
	"go-services/global/common/utils/auth/md5"
)

//key
const secretKey  = "[1iri1*9iri9.]"

/*
	TODO Build Password
 */
func New(password,salting string,createdAt int) string {
	if salting == "" {
		md5.EncodeMD5(password)
	}
	return md5.EncodeMD5(md5.EncodeMD5(md5.EncodeMD5(salting) + password + string(createdAt)) + secretKey)
}
