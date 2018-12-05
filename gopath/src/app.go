package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/light4d/object4d/common/server"

	"github.com/gobestsdk/gobase/log"
	"github.com/light4d/object4d/common/config"
	"github.com/light4d/object4d/router"
)

func main() {

	log.Setlogfile("opengw.log")
	log.Info(log.Fields{"app": "exec args", "args": os.Args})

	defer func() {
		if error := recover(); error != nil {
			log.Fatal(log.Fields{"panic": error})
			exit(-1)
		}
	}()

	go func() {
		singals := make(chan os.Signal)
		signal.Notify(singals,
			syscall.SIGTERM,
			syscall.SIGINT,
			syscall.SIGKILL,
			syscall.SIGHUP,
			syscall.SIGQUIT,
		)
		<-singals
		exit(0)
	}()
	if len(os.Args) > 1 {
		config.ParseConfig(os.Args[1])
	}

	router.Init()
	server.Run()
}

func exit(status int) {
	server.Stop()
	log.Info(log.Fields{"app": status})
	os.Exit(status)
}
