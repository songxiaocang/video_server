package main

import (
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter,httpSC int,errMsg string){
	w.WriteHeader(httpSC)
	io.WriteString(w,errMsg)
}
