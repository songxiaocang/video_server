package main

import (
	"net/http"
	"video_server/api/session"
)

var HEADER_FIELD_SESSIONID = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-User-Name"

func ValidUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_FIELD_SESSIONID)
	if len(sid) == 0 {
		return false
	}

	uname, flag := session.IsSessionExpired(sid)
	if flag {
		return false
	}

	if uname != "" {
		r.Header.Add(HEADER_FIELD_UNAME, uname)
		return true
	}
	return false
}

func ValidUser(r *http.Request) bool {
	uname := r.Header.Get(HEADER_FIELD_UNAME)
	if len(uname) == 0 {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}
	return true
}
