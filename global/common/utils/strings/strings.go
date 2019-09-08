package strings

import (
	"math/rand"
	"time"
)

var (
	A 	= 	"123456789"
	B 	= 	"0123456789"
	C 	= 	"ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	D 	= 	"abcdefghijklmnopqrstuvwxyz"
	E 	= 	"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	F 	= 	"0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// 随机字符串
func Random(t int, l int) string {
	str := F
	switch t {
		case 1:
			str = A
		case 2:
			str = B
		case 3:
			str = C
		case 4:
			str = D
		case 5:
			str = E
		default:
			str = F
	}
	bytes := []byte(str)
	result := []byte("")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}