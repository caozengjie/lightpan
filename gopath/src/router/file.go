package router

import (
	"github.com/gobestsdk/gobase/httpserver"
	"net/http"
)

func file(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case http.MethodGet:
		file_get(resp, req)
	case http.MethodPost:
		file_post(resp, req)
	case http.MethodDelete:
		file_delete(resp, req)
	default:
		httpserver.Options(req, resp)
	}
}

func file_get(resp http.ResponseWriter, req *http.Request) {
	//uid:=getuid(req)

	//service.GetObject(uid)
}
func file_post(resp http.ResponseWriter, req *http.Request) {

}
func file_delete(resp http.ResponseWriter, req *http.Request) {

}
