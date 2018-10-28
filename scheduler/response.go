package main

import (
	"io"
	"net/http"
)

func sendResponse(w http.ResponseWriter, httpSC int, msg string) {
	w.WriteHeader(httpSC)
	io.WriteString(w, msg)
}
