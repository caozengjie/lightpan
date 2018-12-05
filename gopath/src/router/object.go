package router

import (
	"github.com/gobestsdk/gobase/httpserver"
	"net/http"
)

func object4d(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case http.MethodGet:
		object_get(resp, req)
	case http.MethodPost:
		object_post(resp, req)
	case http.MethodDelete:
		object_delete(resp, req)
	default:
		httpserver.Options(req, resp)
	}
}

func object_get(resp http.ResponseWriter, req *http.Request) {
	//uid:=getuid(req)

	//service.GetObject(uid)
}
func object_post(resp http.ResponseWriter, req *http.Request) {

}
func object_delete(resp http.ResponseWriter, req *http.Request) {

}
