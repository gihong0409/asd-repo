package main

import (
	"flag"
	"os"

	"git.datau.co.kr/ferrari/ferrari-cba/factory"
	"git.datau.co.kr/ferrari/ferrari-cba/process"
)

func destroy(fac *factory.Factory) {
	fac.RedisClient.Close()
}

func main() {

	appEnv := flag.String("app-env", os.Getenv("APP_HOME"), "app env")
	appMode := flag.String("app-mode", os.Getenv("APP_MODE"), "app mode")
	flag.Parse()

	if *appEnv == "" {
		*appEnv = `./`
	}

	if *appMode == "" {
		*appEnv = "DEV"
	}

	fac := factory.Factory{JSONConfigPath: *appEnv, AppMode: *appMode}
	fac.Initialize()
	defer destroy(&fac)

	var proc process.CBAProcess
	proc.Initialize(&fac)
	proc.Processing()
}
