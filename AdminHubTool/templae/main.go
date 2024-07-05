package templae

const TemplateMainStr = `package main

import (
	"flag"
	"github.com/xiaka53/AdminHub/exec/router"
	"github.com/xiaka53/AdminHub/public"
	"github.com/xiaka53/DeployAndLog/lib"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	flag.Parse()
	path := flag.Arg(0)
	if path == "" {
		path = "./conf/local/"
	}
	if err := lib.InitModule(path, []string{"base", "mysql", "redis"}); err != nil {
		log.Fatal(err)
	}
	_ = public.InitMysql()
	defer lib.Destroy()
	public.InitValidate()

	router.HttpServerRun(writeRouter())
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	router.HttpServerStop()
}
`
