package model

import "fmt"

type CommonResp struct {
	Error  interface{}
	Code   int
	Result interface{}
}

type Err struct {
	Detail string
	Data   interface{}
}

func (e Err) Error() string {
	return e.Detail + ":" + fmt.Sprint(e.Data)
}

func NewErr(detail string) error {
	return &Err{Detail: detail}
}

func NewErrData(detail string, data interface{}) error {
	return &Err{Detail: detail, Data: data}
}

var FieldCannotupdate = "there fields can't updated"
var ErrLenBigThan1 = NewErr("len(*)>1,already exist")
var ErrLenNotEqual1 = NewErr("len(*)!=1,create failed")
var ErrLenEqual0 = NewErr("len(*)==0,not found")
