package service

import (
	"crypto/md5"
	"errors"
	"fmt"

	"github.com/light4d/object4d/dao"
	"github.com/light4d/object4d/model"
	"io"
	"strconv"
	"time"
)

func Login(userid, password string) (string, error) {
	existuser, ex := CheckUserExist(userid)
	if !ex {
		return "", errors.New("user not found")
	}

	if (existuser.Password) != model.DBPassword(password) {
		return "", errors.New("用户名密码错误")
	}

	t := Token()

	dao.Redis().Set(t, userid, time.Hour*48)
	return t, nil
}

func Token() string {
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token
}

func Checktoken(token string) string {
	uid, err := dao.Redis().Get(token).Result()
	if err != nil {
		return ""
	}
	return uid
}
