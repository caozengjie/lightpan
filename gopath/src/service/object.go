package service

import (
	"github.com/gobestsdk/gobase/log"

	"github.com/light4d/object4d/dao"
	"github.com/light4d/object4d/model"
)

var allowupdateObject = map[string]interface{}{
	"name": "",
}

func checkandfixCreateObject(object *model.Object) (err error) {
	//TODO your code

	return
}
func get(who, bucket, object string) {
}
func SearchObject(u string, filter map[string]interface{}) (result []model.Object, err error) {
	log.Info(log.Fields{
		"func":   "SearchObjects",
		"filter": filter,
	})
	db := dao.DB()

	err = db.Table("object").Find(&result, filter).Error
	if err != nil {
		log.Warn(log.Fields{
			"Find":   "objects",
			"Result": result,
			"Err":    err.Error(),
		})
	}
	//for i, _ := range result {
	//	result[i] = *(result[i].FixShow())
	//}
	log.Info(log.Fields{
		"func":   "GetObjects",
		"result": result,
	})
	return

}

//
//func DeleteObject(objectid string) (err error) {
//	log.Info(log.Fields{
//		"func":   "DeleteObject",
//		"objectid": objectid,
//	})
//	filter := map[string]interface{}{
//		"id":   objectid,
//		"type": "",
//	}
//	objects, err := SearchObject(nil)
//
//	if err != nil {
//		log.Warn(log.Fields{
//			"GetObjects": objects,
//			"Err":      err.Error(),
//		})
//		return err
//	}
//
//	if len(objects) == 0 {
//		return model.ErrLenEqual0
//	}
//	if len(objects) > 1 {
//		return model.ErrLenBigThan1
//	}
//	err = dao.DB().Where(filter).Table("object").Delete(&model.Object{}).Error
//	if err != nil {
//		log.Warn(log.Fields{
//			"Model.Delete": objects,
//			"Where":        filter,
//			"Err":          err.Error(),
//		})
//	}
//	return err
//}
//
//func CreateObject(object model.Object) (objectid string, err error) {
//	log.Info(log.Fields{
//		"func": "CreateObject",
//		"object": object,
//	})
//
//	err = checkandfixCreateObject(&object)
//	if err != nil {
//		return
//	}
//
//	db := dao.DB()
//	err = db.Table("object").Create(&object).Error
//	if err != nil {
//		log.Warn(log.Fields{
//			"object":       object,
//			"CreateObject": "DB",
//			"Err":        err.Error(),
//		})
//		return
//	}
//
//	objects, err := SearchObject(map[string]interface{}{
//		"name": object.Name,
//		"type": "group",
//	})
//	if err != nil {
//		return
//	}
//	if len(objects) == 0 {
//		return "", model.ErrLenNotEqual1
//	}
//	if len(objects) > 1 {
//		return "", model.ErrLenBigThan1
//	}
//	objectid = (objects[0].ID)
//	log.Info(log.Fields{
//		"func":   "CreateObject",
//		"objectid": objectid,
//	})
//	return
//}

func UpdateObject(id string, updater map[string]interface{}) (err error) {
	log.Info(log.Fields{
		"func":    "UpdateObject",
		"id":      id,
		"updater": updater,
	})

	errfields := checkupdate(updater, allowupdateObject)
	if len(errfields) > 0 {
		return model.NewErrData(model.FieldCannotupdate, errfields)
	}
	db := dao.DB()
	err = db.Table("object").Where("id = ?", id).Where("type = ''").Updates(updater).Error
	if err != nil {
		log.Warn(log.Fields{
			"func":    "UpdateObject Updates",
			"id":      id,
			"updater": updater,
			"Err":     err.Error(),
		})
	}
	return
}

func GetObject(who, bucket, objectid string) (o *model.Object, err error) {

	if who != bucket {

	}

	if err = dao.DB().Table("object").Where("bucket =? and id = ?", bucket, objectid).Find(o).Error; err != nil {
		return nil, err
	}
	return
}
