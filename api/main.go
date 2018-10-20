package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type MiddleWareHandler struct {
	R *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler{
	m = &MiddleWareHandler{R:r}
	return m
}

//ServeHTTP(ResponseWriter, *Request)
func(m *MiddleWareHandler) ServerHTTP(w http.ResponseWriter,r *http.Request){
	//check user session
	ValidUserSession(r)

	m.R.ServeHTTP(w,r)

}

func RegisterHandlers() *httprouter.Router{
	router := httprouter.New()
	router.POST("/user",CreateUser)
	router.POST("/user/:user_name",Login)
	return router
}

func main(){
	r := RegisterHandlers()
	handler := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000",handler)
}
