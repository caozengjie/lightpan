package server

import (
	"github.com/gobestsdk/gobase/httpserver"

	"github.com/light4d/object4d/common/config"
)

var (
	M = httpserver.New("yourfs")
	F = httpserver.New("fs")
	O = httpserver.New("ojbect4d")
)

func Run() {

	M.SetPort(config.APPConfig.HttpPort)
	F.SetPort(config.APPConfig.FsPort)
	O.SetPort(config.APPConfig.Object4dPort)
	go M.Run()
	go O.Run()
	F.Run()
}

func Stop() {
	M.Stop()
	O.Stop()
	F.Stop()
}
