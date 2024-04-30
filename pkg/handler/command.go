package handler

import "env-tools/pkg/repository/command"

func EnvToolsInstall() (err error) {

	pwd, err := command.GetSystemPassword()
	if err != nil {
		return
	}

	service := command.NewEnvService(pwd)

	// // err = service.XcodeInstall()
	// // if err != nil {
	// // 	return
	// // }

	// // if !service.WaitXcodeInstall() {
	// // 	return
	// // }

	// err = service.HomebrewInstall()
	// if err != nil {
	// 	return
	// }

	// err = service.BrewInstallFromKit()
	// if err != nil {
	// 	return
	// }

	// err = service.BrewInstallFromCask()
	// if err != nil {
	// 	return
	// }

	err = service.VSCodeExtensionsInstall()
	if err != nil {
		return
	}

	// err = service.OhMyZshInstall()
	// if err != nil {
	// 	return
	// }

	// err = service.SwitchToZshShell()
	// if err != nil {
	// 	return
	// }

	return
}
