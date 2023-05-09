package helper

import (
	"fundamental-payroll-go/helper/apperrors"
	"os"
	"os/exec"
	"runtime"
)

func ClearTerminal() error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd = exec.Command("clear") // Linux and MacOS example, its tested
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls") //Windows example, its tested
	default:
		return apperrors.New(apperrors.ErrPlatformNotSupported)
	}
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
