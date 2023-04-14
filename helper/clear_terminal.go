package helper

import (
	"fundamental-payroll-go/helper/apperrors"
	"os"
	"os/exec"
	"runtime"
)

func ClearTerminal() error {
	var err error
	switch runtime.GOOS {
	case "linux":
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		err = cmd.Run()
	case "darwin":
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		err = cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		err = cmd.Run()
	default:
		err = apperrors.New(apperrors.ErrPlatformNotSupported)
	}
	return err
}
