package router

import (
	"github.com/light4d/object4d/model"
	"github.com/light4d/object4d/service"

	"encoding/json"
	"errors"
	"github.com/gobestsdk/gobase/httpserver"
	"github.com/gobestsdk/gobase/log"
	"net/http"
)

func getuid(req *http.Request) string {
	c, err := req.Cookie("token")
	if err != nil {
		return ""
	}
	return service.Checktoken(c.Value)
}

func checktoken(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")

	result := model.CommonResp{}
	c, err := req.Cookie("token")
	if err != nil {
		result.Code = -2
		result.Error = "nedd token"
		Endresp(result, resp)
		return
	}
	if service.Checktoken(c.Value) == "" {
		result.Code = -2
		result.Error = errors.New("need token")
		Endresp(result, resp)
		return
	} else {
		switch req.URL.Path {
		case "/user":
			user(resp, req)
		case "/group":
			group(resp, req)
		case "/group/owner":
			group_setowner(resp, req)
		case "/group/user":
			group_user(resp, req)
		}
		return
	}
}
func Endresp(result model.CommonResp, resp http.ResponseWriter) {
	log.Info(log.Fields{
		"resp": result,
	})
	httpserver.Header(resp)

	r, _ := json.Marshal(result)
	resp.Write(r)
}
