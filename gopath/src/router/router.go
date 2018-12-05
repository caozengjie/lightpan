package router

import (
	"github.com/light4d/object4d/common/server"
)

func Init() {

	server.M.ServerMux.HandleFunc("/user", checktoken)
	server.M.ServerMux.HandleFunc("/group", checktoken)
	server.M.ServerMux.HandleFunc("/group/owner", checktoken)
	server.M.ServerMux.HandleFunc("/group/user", checktoken)
	server.M.ServerMux.HandleFunc("/login", login)

	server.F.ServerMux.HandleFunc("/", file)
	server.O.ServerMux.HandleFunc("/", object4d)
}
