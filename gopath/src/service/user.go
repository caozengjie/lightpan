package service

import (
	"github.com/gobestsdk/gobase/log"

	"github.com/light4d/object4d/dao"
	"github.com/light4d/object4d/model"

	"errors"
	"regexp"
	"time"
)

var allowupdateUser = map[string]interface{}{
	"name":     "",
	"password": "",
	"face":     "",
}

func checkandfixCreateUser(user *model.User) (err error) {
	//TODO your code

	m, err := regexp.Match("^[a-zA-Z0-9_-]{4,16}$", []byte(user.ID))
	if user.ID == "" || err != nil || !m {
		return model.NewErrData("user id error format", user.ID)
	}
	user.Password = model.DBPassword(user.Password)
	user.Registetime = time.Now()

	return
}

func CheckUserExist(uid string) (*model.User, bool) {
	us, err := SearchUser(map[string]interface{}{
		"id": uid,
	})
	if err == nil && len(us) == 1 {
		return &(us[0]), true
	}
	return nil, false
}
func SearchUser(filter map[string]interface{}) (result []model.User, err error) {
	log.Info(log.Fields{
		"func":   "SearchUsers",
		"filter": filter,
	})
	db := dao.DB()

	err = db.Table("user").Find(&result, filter).Error
	if err != nil {
		log.Warn(log.Fields{
			"Find":   "users",
			"Result": result,
			"Err":    err.Error(),
		})
	}
	for i, _ := range result {
		result[i] = *(result[i].FixShow())
	}
	log.Info(log.Fields{
		"func":   "GetUsers",
		"result": result,
	})
	return

}

func DeleteUser(userid string) (err error) {
	log.Info(log.Fields{
		"func":   "DeleteUser",
		"userid": userid,
	})
	filter := map[string]interface{}{
		"id":   userid,
		"type": "",
	}

	u, ex := CheckUserExist(userid)
	if !ex {
		return errors.New("user not found")
	}
	err = dao.DB().Where(filter).Table("user").Delete(&model.User{}).Error
	if err != nil {
		log.Warn(log.Fields{
			"Model.Delete": u,
			"Where":        filter,
			"Err":          err.Error(),
		})
	}
	return err
}

func CreateUser(user model.User) (userid string, err error) {
	log.Info(log.Fields{
		"func": "CreateUser",
		"user": user,
	})

	_, ex := CheckUserExist(userid)
	if ex {
		return "", errors.New("user already exist")
	}

	err = checkandfixCreateUser(&user)
	if err != nil {
		return "", err
	}

	db := dao.DB()
	err = db.Table("user").Create(&user).Error
	if err != nil {
		log.Warn(log.Fields{
			"user":       user,
			"CreateUser": "DB",
			"Err":        err.Error(),
		})
		return
	}

	users, err := SearchUser(map[string]interface{}{
		"id":   user.ID,
		"type": "",
	})
	if err != nil {
		return
	}
	if len(users) == 0 {
		return "", model.ErrLenNotEqual1
	}
	if len(users) > 1 {
		return "", model.ErrLenBigThan1
	}
	userid = (users[0].ID)
	log.Info(log.Fields{
		"func":   "CreateUser",
		"userid": userid,
	})
	return
}

func UpdateUser(id string, updater map[string]interface{}) (err error) {
	log.Info(log.Fields{
		"func":    "UpdateUser",
		"id":      id,
		"updater": updater,
	})
	errfields := checkupdate(updater, allowupdateUser)
	if len(errfields) > 0 {
		return model.NewErrData(model.FieldCannotupdate, errfields)
	}

	_, ex := CheckUserExist(id)
	if !ex {
		return errors.New("user not found")
	}

	db := dao.DB()
	err = db.Table("user").Where("id = ?", id).Where("type = ''").Updates(updater).Error
	if err != nil {
		log.Warn(log.Fields{
			"func":    "UpdateUser Updates",
			"id":      id,
			"updater": updater,
			"Err":     err.Error(),
		})
	}
	return
}
