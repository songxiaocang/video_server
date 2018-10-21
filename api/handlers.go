package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"video_server/api/defs"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//io.WriteString(w,"Create User handler")
	/**
	1、接收数据序列化
	2、插入用户
	3、响应数据反序列化
	*/
	request, _ := ioutil.ReadAll(r.Body)

	ubody := &defs.UserCredential{}
	err := json.Unmarshal(request, ubody)
	if err != nil {
		//sendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		sendErrorResponse(w, &defs.ErrorRequestBodyParseFailed)
		return
	}
	err2 := dbops.AddUserCredential(ubody.Username, ubody.Pwd)
	if err2 != nil {
		sendErrorResponse(w, &defs.ErrorDBError)
		return
	}
	uuid := session.GenerateNewSessionId(ubody.Username)
	signUp := &defs.SignUp{Success: true, SessionId: uuid}
	res, e := json.Marshal(signUp)
	if e != nil {
		sendErrorResponse(w, &defs.ErrorInternalFaults)
		return
	}

	sendNormalResponse(w, string(res), 201)

}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}
