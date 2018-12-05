package model

import (
	"github.com/gobestsdk/gobase/utils"
	"time"
)

type Group struct {
	ID          string
	Name        string
	Type        string `json:"-"`
	Face        string
	Parent      string
	Registetime interface{}
}

func (u *Group) FixShow() *Group {
	if u.Registetime != nil {
		u.Registetime = (u.Registetime.(time.Time)).Format(utils.DateTimeFormart)
	}

	return u
}
