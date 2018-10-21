package main

import (
	"encoding/json"
	"io"
	"net/http"
	"video_server/api/defs"
)

func sendErrorResponse(w http.ResponseWriter, resp *defs.ErrResponse) {
	w.WriteHeader(resp.HttpSC)

	res, e := json.Marshal(&resp.Error)
	if e == nil {
		io.WriteString(w, string(res))
	}
}

func sendNormalResponse(w http.ResponseWriter, resp string, httpSC int) {
	w.WriteHeader(httpSC)
	io.WriteString(w, resp)
}
