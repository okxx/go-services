package ticket

import (
	"go-services/global/common/utils/auth/md5"
	"go-services/global/common/utils/strings"
)

/*
	TODO Build Ticket
 */
func New(key string) string {
	md5Key 	:= md5.EncodeMD5(key)
	randStr := md5.EncodeMD5(strings.Random(6,22))
	return md5.EncodeMD5(randStr+md5Key)
}