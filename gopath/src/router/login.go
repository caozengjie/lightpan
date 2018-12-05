package router

import (
	"encoding/json"

	"io/ioutil"

	"github.com/light4d/object4d/model"
	"net/http"

	"github.com/gobestsdk/gobase/httpserver"
	"github.com/light4d/object4d/service"
)

func login(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case http.MethodPost:
		login_post(resp, req)
	default:
		httpserver.Options(req, resp)
	}
}

func login_post(resp http.ResponseWriter, req *http.Request) {
	result := model.CommonResp{}
	m := map[string]string{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
		Endresp(result, resp)
		return
	}
	err = json.Unmarshal(body, &m)
	if err != nil {
		result.Code = -1
		result.Error = err
		Endresp(result, resp)
		return
	}
	t, err := service.Login(m["id"], m["password"])
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
		Endresp(result, resp)
		return
	}
	result.Result = struct {
		UserID string
		Token  string
	}{UserID: m["id"], Token: t}
	cookie := http.Cookie{Name: "token", Value: t, Path: "/"}
	http.SetCookie(resp, &cookie)
	Endresp(result, resp)
}
