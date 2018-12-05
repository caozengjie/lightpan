package router

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/light4d/object4d/model"
	"net/http"

	"github.com/gobestsdk/gobase/httpserver"
	"github.com/light4d/object4d/service"
)

func user(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case http.MethodGet:
		user_get(resp, req)
	case http.MethodPost:
		user_post(resp, req)
	case http.MethodPut:
		user_put(resp, req)
	case http.MethodDelete:
		user_delete(resp, req)
	default:
		httpserver.Options(req, resp)
	}
}
func user_get(resp http.ResponseWriter, req *http.Request) {
	result := model.CommonResp{}
	filter := httpserver.Getfilter(req)
	filter["type"] = ""
	us, err := service.SearchUser(filter)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
	} else {
		result.Result = us
	}
	Endresp(result, resp)
}
func user_post(resp http.ResponseWriter, req *http.Request) {
	result := model.CommonResp{}

	user := model.User{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
		Endresp(result, resp)
		return
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
		Endresp(result, resp)
		return
	}

	userid, err := service.CreateUser(user)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
		Endresp(result, resp)
		return
	}
	result.Result = struct {
		UserID string
	}{UserID: userid}
	Endresp(result, resp)
}

func user_put(resp http.ResponseWriter, req *http.Request) {
	result := model.CommonResp{}

	id := req.URL.Query().Get("id")
	if id == "" {
		result.Code = -1
		result.Error = errors.New("id不能为空")
		Endresp(result, resp)
		return
	}
	updater := make(map[string]interface{})
	err := httpserver.Unmarshalreqbody(req, &updater)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
		Endresp(result, resp)
		return
	}
	err = service.UpdateUser(id, updater)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
	}
	Endresp(result, resp)
}
func user_delete(resp http.ResponseWriter, req *http.Request) {
	result := model.CommonResp{}

	userid := req.URL.Query().Get("id")

	if userid != "" {
		err := service.DeleteUser(userid)
		if err != nil {
			result.Code = -1
			result.Error = err.Error()
		}
		Endresp(result, resp)
		return
	} else {
		result.Error = errors.New("whick one do you want to delete?")
		Endresp(result, resp)
		return
	}

}
