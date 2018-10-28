package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sxc/config"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{}
}

func request(b *ApiBody, w http.ResponseWriter, r *http.Request) {
	//var resp *http.Response
	//var err error

	u, _ := url.Parse(b.Url)
	u.Host = config.GetLBAddr() + ":" + u.Port()
	newUrl := u.String()

	switch b.Method {
	case http.MethodGet:
		req, _ := http.NewRequest("GET", newUrl, nil)
		req.Header = r.Header
		resp, err := httpClient.Do(req)
		if err != nil {
			log.Printf("httpclient error: %v", err)
			return
		}
		normalResponse(w, resp)
	case http.MethodPost:
		req, _ := http.NewRequest("POST", newUrl, bytes.NewBuffer([]byte(b.ReqBody)))
		req.Header = r.Header
		resp, err := httpClient.Do(req)
		if err != nil {
			log.Printf("httpclient error: %v", err)
			return
		}
		normalResponse(w, resp)
	case http.MethodDelete:
		req, _ := http.NewRequest("Delete", newUrl, nil)
		req.Header = r.Header
		resp, err := httpClient.Do(req)
		if err != nil {
			log.Printf("httpclient error: %v", err)
			return
		}
		normalResponse(w, resp)
	default:
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "bad api request")

	}
}

func normalResponse(w http.ResponseWriter, r *http.Response) {
	res, e := ioutil.ReadAll(r.Body)
	if e != nil {
		resp, _ := json.Marshal(ErrInternalFaults)
		w.WriteHeader(500)
		io.WriteString(w, string(resp))
		return
	}

	w.WriteHeader(r.StatusCode)
	io.WriteString(w, string(res))
}
