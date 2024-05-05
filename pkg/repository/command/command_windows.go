//go:build windows
// +build windows

package command

import "os/exec"

func setSysProcAttr(cmd *exec.Cmd) {
	// Windows-specific code or no-op
}
