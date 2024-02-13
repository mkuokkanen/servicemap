package server

import (
	"fmt"
	"net/http"
)

func accessLog(r *http.Request, funcName string) {
	fmt.Println("OP", r.Method, r.RequestURI, r.Proto, funcName)
}
