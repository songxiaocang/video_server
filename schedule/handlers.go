package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"

)

func delvidrecHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	vid := p.ByName("video_id")

	if len(vid) == 0 {
		log.Printf("video_id is null: %v",vid)
		sendResponse(w,400,"video_id is null")
		return
	}


	err := dbops.addVideoRec(vid)

	if err!=nil {
		log.Printf("internal error: %v",err)
		sendResponse(w,500,"internal error")
		return
	}

	sendResponse(w,200,"operation success")
}
