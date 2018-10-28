package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", homeHandler)
	router.POST("/", homeHandler)
	router.GET("/userhome", userHomeHandler)
	router.POST("/userhome", userHomeHandler)
	//api透传
	router.POST("/api", apiHandler)
	//proxy转发
	router.GET("/videos/:vid_id", proxyVideoHandler)

	router.POST("/upload/:vid_id", proxyUploadHandler)

	//读取指定目录前端模板
	router.ServeFiles("/statics/*filepath", http.Dir("./template"))

	return router

}

func main() {
	handlers := RegisterHandlers()
	http.ListenAndServe(":8082", handlers)
}
