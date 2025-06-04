package main

import (
	"ASD/factory"
	"ASD/process"
	"flag"
	"log"
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
		*appEnv = `./`
	}

	//if *appMode == "" {
	//	*appMode = "LIVE"
	//}
	fac := factory.Factory{JSONConfigPath: *appEnv, AppMode: *appMode}
	fac.Initialize()
	defer destroy(&fac)

	log.Printf("Tartget Service : %d", fac.Propertys().ServiceNames)

	var proc process.ASDProcess
	proc.Initialize(&fac)
	proc.Processing()

}
