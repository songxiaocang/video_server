package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"video_server/api/dbops"
	"video_server/api/defs"
	"video_server/api/session"
	"video_server/api/utils"
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
	res, _ := ioutil.ReadAll(r.Body)
	log.Printf("%s", res)
	ubody := &defs.UserCredential{}
	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, &defs.ErrorRequestBodyParseFailed)
		return
	}

	//validate the request body
	uname := p.ByName("user_name")
	log.Printf("login url uname : %s", uname)
	log.Printf("login body uname : %s", ubody.Username)
	if uname != ubody.Username {
		sendErrorResponse(w, &defs.ErrorNotAuthUser)
		return
	}

	log.Printf("ubody username: %s", ubody.Username)
	pwd, e := dbops.GetUserCredential(uname)
	log.Printf("login pwd: %s", pwd)
	log.Printf("login body pwd: %s", ubody.Pwd)
	if e != nil || len(pwd) != 0 || pwd != ubody.Pwd {
		sendErrorResponse(w, &defs.ErrorNotAuthUser)
		return
	}

	sid := session.GenerateNewSessionId(ubody.Username)
	si := &defs.SignIn{Success: true, SessionId: sid}
	if resp, err := json.Marshal(si); err != nil {
		sendErrorResponse(w, &defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}

}

func GetUserInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !ValidUser(w, r) {
		log.Printf("not auth user")
		return
	}

	name := ps.ByName("user_name")
	user, e := dbops.GetUser(name)
	if e != nil {
		log.Printf("db error: %v", e)
		sendErrorResponse(w, &defs.ErrorDBError)
		return
	}

	userInfo := &defs.UserInfo{Id: user.Id}
	if resp, err := json.Marshal(userInfo); err != nil {
		sendErrorResponse(w, &defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}

}

func AddNewVideo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !ValidUser(w, r) {
		log.Printf("not auth user")
		return
	}

	res, _ := ioutil.ReadAll(r.Body)
	vibody := &defs.NewVideo{}
	if err := json.Unmarshal(res, vibody); err != nil {
		sendErrorResponse(w, &defs.ErrorRequestBodyParseFailed)
		return
	}

	info, e := dbops.AddVideoInfo(vibody.AuthorId, vibody.Name)
	if e != nil {
		sendErrorResponse(w, &defs.ErrorDBError)
		return
	}

	if resp, err := json.Marshal(info); err != nil {
		sendErrorResponse(w, &defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 201)
	}

}

func ListAllVideos(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !ValidUser(w, r) {
		log.Printf("not auth user")
		return
	}

	name := ps.ByName("user_name")

	//todo
	videos, err := dbops.ListAllVideos(name, 0, utils.GetCurrentTimestmapSec())

	if err != nil {
		sendErrorResponse(w, &defs.ErrorDBError)
		return
	}

	vs := &defs.VideosInfo{Videos: videos}
	if resp, err := json.Marshal(vs); err != nil {
		sendErrorResponse(w, &defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}

func DeleteVideo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !ValidUser(w, r) {
		log.Printf("not auth user")
		return
	}

	vid := ps.ByName("vid_id")
	err := dbops.DelVideoInfo(vid)
	if err != nil {
		log.Printf("db error: %v", err)
		sendErrorResponse(w, &defs.ErrorDBError)
		return
	}

	//向streamserver发起 del video的请求 ,异步执行
	go utils.SendDelVideoRequest(vid)

	sendNormalResponse(w, "delete success", 204)
}

func PostComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !ValidUser(w, r) {
		log.Printf("not auth user")
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)
	ncBody := &defs.Comment{}
	//ncBody := &defs.NewComment{}
	if err := json.Unmarshal(reqBody, ncBody); err != nil {
		log.Printf("%v", err)
		sendErrorResponse(w, &defs.ErrorRequestBodyParseFailed)
		return
	}

	vid := ps.ByName("vid_id")
	err := dbops.AddComment(vid, ncBody.Author, ncBody.Content)
	if err != nil {
		log.Printf("addComment error: %v", err)
		sendErrorResponse(w, &defs.ErrorDBError)
	} else {
		sendNormalResponse(w, "ok", 201)
	}
}

func ShowComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !ValidUser(w, r) {
		log.Printf("not auth user")
		return
	}

	vid := ps.ByName("vid_id")
	comments, err := dbops.ListComments(vid, 0, utils.GetCurrentTimestmapSec())
	if err != nil {
		sendErrorResponse(w, &defs.ErrorDBError)
		return
	}

	cs := &defs.Comments{Comments: comments}
	if resp, err := json.Marshal(cs); err != nil {
		sendErrorResponse(w, &defs.ErrorDBError)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}

}
