package cookie

import (
	"net/http"
	"go-services/global/common/utils/io"
)

//write info to cookie
func Set(cookies map[string]string) {
	for k,v := range cookies {
		http.SetCookie(io.App.Context.Writer,&http.Cookie{Name : k, Value : v, Path : "/", HttpOnly: true,})
	}
}