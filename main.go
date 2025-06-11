package main

import (
	"flag"
	"os"

	"git.datau.co.kr/earth/earth-asd/factory"
	"git.datau.co.kr/earth/earth-asd/process"
)

func main() {
	println("app started")

	appEnv := flag.String("app-env", os.Getenv("APP_HOME"), "app env")
	appMode := flag.String("app-mode", os.Getenv("APP_MODE"), "app mode")
	flag.Parse()

	if *appEnv == "" {
		*appEnv = `./`
	}

	fac := factory.Factory{JSONConfigPath: *appEnv, AppMode: *appMode}

	fac.Initialize()

	var proc process.ASDProcess
	proc.Initialize(&fac)
	proc.Processing()
}
