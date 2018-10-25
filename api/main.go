package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video_server/api/session"
)

type MiddleWareHandler struct {
	R *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := &MiddleWareHandler{R: r}
	//m.R=r
	return m
}

//ServeHTTP(ResponseWriter, *Request)
func (m *MiddleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//check user session
	ValidUserSession(r)

	m.R.ServeHTTP(w, r)

}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)

	router.GET("/user/:user_name", GetUserInfo)
	router.POST("/user/:user_name/videos", AddNewVideo)
	router.GET("/user/:user_name/videos", ListAllVideos)
	router.DELETE("/user/:user_name/videos/:vid_id", DeleteVideo)
	router.POST("/videos/:vid_id/comments", PostComment)
	router.GET("/videos/:vid_id/comments", ShowComments)
	return router
}

func Prepare() {
	session.LoadSessionFromDB()
}

func main() {
	Prepare()
	r := RegisterHandlers()
	handler := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", handler)
}
