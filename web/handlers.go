package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sxc/config"
)

type HomePage struct {
	Name string
}

type UserPage struct {
	Name string
}

func homeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uname, err1 := r.Cookie("username")
	session, err2 := r.Cookie("session")

	if err1 != nil || err2 != nil {
		p := &HomePage{Name: "songxiaocang"}
		t, e := template.ParseFiles("./template/home.html")
		if e != nil {
			log.Printf("parse home.html error: %v", e)
			return
		}
		t.Execute(w, p)
		return
	}

	if len(uname.Value) > 0 && len(session.Value) > 0 {
		http.Redirect(w, r, "/userhome", http.StatusFound)
	}

}

func userHomeHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	cname, err1 := r.Cookie("username")
	_, err2 := r.Cookie("session")

	if err1 != nil || err2 != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	fname := r.FormValue("username")

	var up *UserPage
	if len(cname.Value) > 0 {
		up = &UserPage{Name: cname.Value}
	} else if len(fname) > 0 {
		up = &UserPage{Name: fname}
	}

	t, e := template.ParseFiles("./template/userhome.html")
	if e != nil {
		log.Printf("parse userhome.html error: %v", e)
		return
	}

	t.Execute(w, up)

}

func apiHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method != http.MethodPost {
		res, _ := json.Marshal(ErrRequestNotRecognized)
		io.WriteString(w, string(res))
		return
	}

	res, _ := ioutil.ReadAll(r.Body)
	apiBody := &ApiBody{}
	if err := json.Unmarshal(res, apiBody); err != nil {
		res, _ := json.Marshal(ErrRequestBodyParseFailed)
		io.WriteString(w, string(res))
		return
	}

	request(apiBody, w, r)
	defer r.Body.Close()
}

func proxyVideoHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	u, _ := url.Parse("http://" + config.GetLBAddr() + ":9000/")
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w, r)
}

func proxyUploadHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	u, _ := url.Parse("http://" + config.GetLBAddr() + ":9000/")
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w, r)
}
