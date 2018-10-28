package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	files, e := template.ParseFiles("E:/videos/upoad.html")
	if e != nil {
		log.Printf("not found template: %v", e)
		return
	}
	files.Execute(w, nil)
}

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid_id")
	videoDir := VIDEO_DIR + vid

	file, e := os.Open(videoDir)
	if e != nil {
		log.Printf("open error: %v", e)
		sendErrorResponse(w, http.StatusInternalServerError, "internal error")
		return
	}

	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), file)

	defer file.Close()

	//上云
	//targetUrl := "http://xxx.aliyun.com"+p.ByName("vid_id")
	//http.Redirect(w,r,targetUrl,301)
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	/**
	1、文件大小校验
	2、获取文件
	3、读取文件
	4、写入本地目录
	5、响应
	*/
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)

	if e := r.ParseMultipartForm(MAX_UPLOAD_SIZE); e != nil {
		log.Printf("over by max upload size")
		sendErrorResponse(w, http.StatusBadRequest, "over by max upload size")
		return
	}

	file, _, e := r.FormFile("file")
	if e != nil {
		log.Printf("get file error: %v", e)
		sendErrorResponse(w, http.StatusInternalServerError, "internal error")
		return
	}

	data, err1 := ioutil.ReadAll(file)
	if err1 != nil {
		log.Printf("read file error: %v", err1)
		sendErrorResponse(w, http.StatusInternalServerError, "internal error")
		return
	}

	err2 := ioutil.WriteFile(VIDEO_DIR+p.ByName("vid_id"), data, 0666)
	if err2 != nil {
		log.Printf("write file error: %v", err2)
		sendErrorResponse(w, http.StatusInternalServerError, "internal error")
		return
	}

	//上云
	ossfn := "videos/" + p.ByName("vid_id")
	path := "E:\\videos\\" + p.ByName("vid_id")
	bn := "songxc-video2"
	result := UploadToOss(ossfn, path, bn)
	if !result {
		sendErrorResponse(w, http.StatusInternalServerError, "internal serive error")
		return
	}
	os.Remove(path)

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "upload success!!!")

}
