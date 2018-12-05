package model

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gobestsdk/gobase/utils"
	"time"
)

type User struct {
	ID          string
	Name        string
	Type        string `json:"-"`
	Face        string
	Registetime interface{}
	Password    string `json:"-"`
	Parent      string `json:"-"`
}

func (u *User) FixShow() *User {
	if u.Registetime != nil {
		u.Registetime = (u.Registetime.(time.Time)).Format(utils.DateTimeFormart)
	}

	return u
}
func DBPassword(password string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(password))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
