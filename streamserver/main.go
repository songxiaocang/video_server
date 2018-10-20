package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type middlewareHandler struct {
	r *httprouter.Router
	cl *connLimiter
}

func newMiddlewareHandler(router *httprouter.Router,cc int) http.Handler{
	m := middlewareHandler{}
	m.r = router
	m.cl = newConnLimiter(cc)
	return m
}

func (m middlewareHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	if !m.cl.getConn() {
		log.Printf("too many request count")
		sendErrorResponse(w,http.StatusBadRequest,"too many request count")
		return
	}

	m.r.ServeHTTP(w,r)
	defer m.cl.releaseConn()
}

func registerHandler() *httprouter.Router{
	router := httprouter.New()
	router.GET("/videos/:vid_id",streamHandler)
	router.POST("/upload/:vid_id",uploadHandler)
	router.GET("/testpage",testPageHandler)

	return router
}


func main(){
	router := registerHandler()
	mh := newMiddlewareHandler(router, 2)
	http.ListenAndServe(":9000",mh)
}
