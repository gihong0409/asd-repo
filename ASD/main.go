package main

import (
	"ASD/factory"
	"flag"
	"os"
)

func destroy(fac *factory.Factory) {

}

func main() {

	print("app started")

	appEnv := flag.String("app-env", os.Getenv("APP_HOME"), "app env")
	appMode := flag.String("app-mode", os.Getenv("APP_MODE"), "app mode")
	flag.Parse()

	if *appEnv == "" {
		*appEnv = `../`
	}

	if *appMode == "" {
		*appMode = "LIVE"
	}

	fac := factory.Factory{JSONConfigPath: *appEnv, AppMode: *appMode}
	fac.Initialize()
	defer destroy(&fac)

	println("최종: ", fac.TargetService)
}
