package v

import "github.com/astaxie/beego/validation"


func V(cs interface{}) bool {
	valid := validation.Validation{}
	if ok,_ := valid.Valid(cs); !ok {
		return false
	}
	return true
}