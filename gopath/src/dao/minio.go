package dao

import (
	"github.com/gobestsdk/gobase/log"
	"github.com/light4d/object4d/model"

	"github.com/minio/minio-go"
)

func GetMinioconfig(minioid int) (m *model.Miniocon, err error) {
	err = DB().Table("miniocon").Find(minioid, m).Error
	return
}

func NewMinioclient(minioid int) (c *minio.Client, err error) {
	// 初使化minio client对象。
	cfg, err := GetMinioconfig(minioid)
	if err != nil {
		log.Warn(log.Fields{
			"Func": "NewMinioclient",
			"Err":  err.Error(),
		})
		return nil, err
	}
	c, err = minio.New(cfg.Endpoint, cfg.AccessKeyID, cfg.SecretAccessKey, cfg.Secure)
	if err != nil {
		log.Warn(log.Fields{
			"Func": "NewMinioclient",
			"Err":  err.Error(),
		})
	}
	return
}
