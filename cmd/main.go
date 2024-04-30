package main

import (
	"env-tools/pkg/handler"
	"env-tools/setting"
)

func main() {

	err := setting.Init()
	if err != nil {
		return
	}

	err = handler.EnvToolsInstall()
	if err != nil {
		return
	}
}
