package command

import (
	"env-tools/setting"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"

	"golang.org/x/term"
)

// GetSystemPassword gets the system password from the user.
func GetSystemPassword() (password string, err error) {

	fmt.Print("Please enter your password: ")

	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println("\nError reading password.")
		return
	}

	password = string(bytePassword)
	return
}

type EnvService struct {
	password string
}

func NewEnvService(pwd string) *EnvService {
	return &EnvService{
		password: pwd,
	}
}

// XcodeInstall installs Xcode.
func (e *EnvService) XcodeInstall() (err error) {

	cmd := exec.Command("sh", "-c", fmt.Sprintf("echo %s | sudo -S xcode-select --install", e.password))
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error starting command.")
		return
	}

	return
}

// WaitXcodeInstall waits for Xcode to be installed.
func (e *EnvService) WaitXcodeInstall() bool {

	fmt.Println("wait xcode install...")

	timeout := time.After(10 * time.Minute)

	for {
		select {
		case <-timeout:
			fmt.Println("Timeout after 10 minutes.")
			return false

		default:
			cmd := exec.Command("xcode-select", "-p")
			cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Run()
			if err == nil {
				fmt.Println("Xcode installed.")
				return true
			}
		}
	}
}

// HomebrewInstall installs Homebrew.
func (e *EnvService) HomebrewInstall() (err error) {

	cmd := exec.Command("sh", "-c", fmt.Sprintf(`echo %s | sudo -S /bin/bash -c $(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)`, e.password))
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error waiting for command to finish.")
		return
	}

	return
}

// BrewInstallFromKit installs the kit packages using Homebrew.
func (e *EnvService) BrewInstallFromKit() (err error) {

	// Get the list of kit packages from the settings.
	kits := setting.GetBrewKit()

	// Install each kit package.
	for _, kit := range kits {

		cmd := exec.Command("sh", "-c", fmt.Sprintf(`echo %s | sudo -S brew install %s`, e.password, kit))
		cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			fmt.Println("Error waiting for command to finish.")
			return
		}
	}

	return
}

// BrewInstallFromCask installs the cask packages using Homebrew.
func (e *EnvService) BrewInstallFromCask() (err error) {

	// Get the list of cask packages from the settings.
	casks := setting.GetBrewCask()

	// Install each cask package.
	for _, cask := range casks {

		cmd := exec.Command("sh", "-c", fmt.Sprintf(`echo %s | sudo -S brew install --cask %s`, e.password, cask))
		cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			fmt.Println("Error waiting for command to finish.")
			return
		}
	}

	return
}

// OhMyZshInstall installs Oh My Zsh.
func (e *EnvService) OhMyZshInstall() (err error) {

	cmd := exec.Command("sh", "-c", fmt.Sprintf(`echo %s | sudo -S sh -c "RUNZSH=no $(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"`, e.password))
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error waiting for command to finish.")
		return
	}

	return
}

// SwitchToZshShell switches the shell to Zsh.
func (e *EnvService) SwitchToZshShell() (err error) {

	cmd := exec.Command("sh", "-c", fmt.Sprintf(`echo %s | sudo -S chsh -s /bin/zsh`, e.password))
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error waiting for command to finish.")
		return
	}

	return
}

// VSCodeExtensionsInstall installs the extensions for Visual Studio Code.
func (e *EnvService) VSCodeExtensionsInstall() (err error) {

	// Get the list of extensions from the settings.
	extList := setting.GetVSCodeExtensions()

	// Install each extension.
	for _, extension := range extList {

		cmd := exec.Command("sh", "-c", fmt.Sprintf(`echo %s | sudo -S code --install-extension %s`, e.password, extension))
		cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			fmt.Println("Error waiting for command to finish.")
			return
		}
	}

	return
}
