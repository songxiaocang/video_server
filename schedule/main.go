package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video_server/schedule/taskrunner"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/video_del_rec/:vid_id", DelvidrecHandler)

	return router
}

func main() {
	go taskrunner.Start()
	handlers := RegisterHandlers()
	http.ListenAndServe(":9001", handlers)
}
