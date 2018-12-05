package router

import (
	"encoding/json"
	"errors"

	"github.com/gobestsdk/gobase/httpserver"
	"github.com/light4d/object4d/model"
	"github.com/light4d/object4d/service"
	"io/ioutil"
	"net/http"
)

func group_user(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case http.MethodGet:
		group_user_get(resp, req)
	case http.MethodPost:
		group_user_post(resp, req)
	case http.MethodPut:
		group_user_put(resp, req)
	case http.MethodDelete:
		group_user_delete(resp, req)
	default:
		httpserver.Options(req, resp)
	}
}

func group_user_get(resp http.ResponseWriter, req *http.Request) {
	result := model.CommonResp{}
	filter := httpserver.Getfilter(req)
	if g, have := filter["id"].(string); have {
		us, err := service.GetGroupuser(g)
		if err != nil {
			result.Code = -1
			result.Error = err.Error()
		} else {
			result.Result = us
		}
	}
	if u, have := filter["user"].(string); have {
		gs, err := service.GetUsergroup(u)
		if err != nil {
			result.Code = -1
			result.Error = err.Error()
		} else {
			result.Result = gs
		}
	}

	Endresp(result, resp)
}

func group_user_post(resp http.ResponseWriter, req *http.Request) {
	result := model.CommonResp{}

	us := make([]string, 0)
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
		Endresp(result, resp)
		return
	}
	err = json.Unmarshal(body, &us)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
		Endresp(result, resp)
		return
	}
	id := req.URL.Query().Get("id")
	if id == "" {
		result.Code = -1
		result.Error = errors.New("id不能为空")
		Endresp(result, resp)
		return
	}
	uid := getuid(req)
	err = service.AddGroupusers(uid, id, us)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
		Endresp(result, resp)
		return
	}
	Endresp(result, resp)
}

func group_user_put(resp http.ResponseWriter, req *http.Request) {
	result := model.CommonResp{}

	us := make([]string, 0)
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
		Endresp(result, resp)
		return
	}
	err = json.Unmarshal(body, &us)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
		Endresp(result, resp)
		return
	}
	id := req.URL.Query().Get("id")
	if id == "" {
		result.Code = -1
		result.Error = errors.New("id不能为空")
		Endresp(result, resp)
		return
	}
	uid := getuid(req)
	err = service.ResetGroupusers(uid, id, us)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
		Endresp(result, resp)
		return
	}
	Endresp(result, resp)
}

func group_user_delete(resp http.ResponseWriter, req *http.Request) {
	result := model.CommonResp{}

	us := make([]string, 0)
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
		Endresp(result, resp)
		return
	}
	err = json.Unmarshal(body, &us)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
		Endresp(result, resp)
		return
	}
	id := req.URL.Query().Get("id")
	if id == "" {
		result.Code = -1
		result.Error = errors.New("id不能为空")
		Endresp(result, resp)
		return
	}
	uid := getuid(req)
	err = service.DeleteGroupusers(uid, id, us)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
		Endresp(result, resp)
		return
	}
	Endresp(result, resp)
}
