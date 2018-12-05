package service

import (
	"errors"
	"fmt"
	"github.com/gobestsdk/gobase/log"
	"github.com/light4d/object4d/dao"
	"github.com/light4d/object4d/model"
	"strings"
	"time"
)

func SearchGroupuser(filter map[string]interface{}) (result []model.Groupuser, err error) {
	log.Info(log.Fields{
		"func":   "SearchGroupuser",
		"filter": filter,
	})

	db := dao.DB()
	err = db.Table("groupuser").Find(&result, filter).Error
	if err != nil {
		log.Warn(log.Fields{
			"Find":   "groupuser",
			"Result": result,
			"Err":    err.Error(),
		})
	}
	for i, _ := range result {
		result[i] = *(result[i].FixShow())
	}
	log.Info(log.Fields{
		"func":   "SearchGroup",
		"result": result,
	})
	return

}
func GetGroupuser(group string) (result []model.User, err error) {
	log.Info(log.Fields{
		"func":   "GetGroupuser",
		"filter": group,
	})

	db := dao.DB()

	db.Raw("SELECT user.id,user.name,user.face,user  FROM user,groupuser WHERE groupuser.id= ? and groupuser.user=user.id and user.type='' ", group).Scan(&result)
	if err != nil {
		log.Warn(log.Fields{
			"Find":   "groupuser",
			"Result": result,
			"Err":    err.Error(),
		})
	}
	for i, _ := range result {
		result[i] = *(result[i].FixShow())
	}
	log.Info(log.Fields{
		"func":   "GetGroupuser",
		"result": result,
	})
	return
}
func GetUsergroup(user string) (result []model.Group, err error) {
	log.Info(log.Fields{
		"func":   "GetUsergroup",
		"filter": user,
	})

	db := dao.DB()

	db.Raw("SELECT user.id,user.name,user.face,user.parent,user.registetime FROM user,groupuser WHERE groupuser.user= ? and groupuser.user=? and user.type= 'group' ", user, user).Scan(&result)
	if err != nil {
		log.Warn(log.Fields{
			"Find":   "groupuser",
			"Result": result,
			"Err":    err.Error(),
		})
	}
	for i, _ := range result {
		result[i] = *(result[i].FixShow())
	}
	log.Info(log.Fields{
		"func":   "GetGroupuser",
		"result": result,
	})
	return
}
func AddGroupusers(who, group string, us []string) (err error) {

	log.Info(log.Fields{
		"func":   "Addusers",
		"detail": who + " add users: " + fmt.Sprint(us) + " to " + group,
	})

	g, ex := CheckGroupExist(group)
	if !ex {
		return errors.New("group not found")
	}
	if g.Parent != who {
		return errors.New("you no permission add user to this group")
	}

	for _, u := range us {
		_, ex := CheckUserExist(u)
		if !ex {
			return model.NewErrData("user not exist", u)
		}
	}
	db := dao.DB()

	for _, u := range us {
		if err = db.Table("groupuser").Create(&model.Groupuser{
			ID:       group,
			User:     u,
			Jointime: time.Now(),
		}).Error; err != nil {
			log.Warn(log.Fields{
				"groupuser":       group,
				"CreateGroupuser": "DB",
				"Err":             err.Error(),
			})
			return
		}
	}
	return db.Commit().Error
}

func DeleteGroupusers(who, group string, us []string) (err error) {
	log.Info(log.Fields{
		"func":   "Addusers",
		"detail": who + " delete users: " + fmt.Sprint(us) + " from " + group,
	})
	g, ex := CheckGroupExist(group)
	if !ex {
		return errors.New("group not found")
	}
	if g.Parent != who {
		return errors.New("you no permission delete user from this group")
	}
	for _, u := range us {
		_, ex := CheckUserExist(u)
		if !ex {
			return model.NewErrData("user not exist", u)
		}
	}
	db := dao.DB()

	if err = db.Table("groupuser").Delete(&model.Groupuser{}).Where("user in (" + strings.Join(us, ",") + ")").Error; err != nil {
		log.Warn(log.Fields{
			"groupuser":       group,
			"CreateGroupuser": "DB",
			"Err":             err.Error(),
		})

		return err
	}
	return
}
func ResetGroupusers(who, group string, us []string) (err error) {
	log.Info(log.Fields{
		"func":   "Addusers",
		"detail": who + " reset users: " + fmt.Sprint(us) + " to " + group,
	})
	g, ex := CheckGroupExist(group)
	if !ex {
		return errors.New("group not found")
	}
	if g.Parent != who {
		return errors.New("you no permission reset user to this group")
	}
	for _, u := range us {
		_, ex := CheckUserExist(u)
		if !ex {
			return model.NewErrData("user not exist", u)
		}
	}
	db := dao.DB()

	if err = db.Table("groupuser").Delete(model.Groupuser{}, map[string]interface{}{}).Error; err != nil {
		log.Warn(log.Fields{
			"groupuser":       group,
			"CreateGroupuser": "DB",
			"Err":             err.Error(),
		})

		return err
	}
	for _, u := range us {
		if err = db.Table("groupuser").Create(&model.Groupuser{
			ID:       group,
			User:     u,
			Jointime: time.Now(),
		}).Error; err != nil {
			log.Warn(log.Fields{
				"groupuser":       group,
				"CreateGroupuser": "DB",
				"Err":             err.Error(),
			})

			return err
		}
	}
	return
}
