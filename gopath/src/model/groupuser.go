package model

import (
	"time"

	"github.com/gobestsdk/gobase/utils"
)

type Groupuser struct {
	ID       string
	User     string
	Jointime interface{}
}

func (u *Groupuser) FixShow() *Groupuser {
	if u.Jointime != nil {
		u.Jointime = (u.Jointime.(time.Time)).Format(utils.DateTimeFormart)
	}
	return u
}
