package environment

import (
	_ "log"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)


func StartVirtualEnvironment(path string) {
	shell_path, command_path, err := os_specific_command_paths(path)

	if err != nil {
		panic(err)
	}

	if err := syscall.Exec(shell_path, []string{command_path}, os.Environ()); err != nil {
		panic(err)
	}
}


func os_specific_command_paths(base_path string) (shell_path, command_path string, err error) {
	if runtime.GOOS == "windows" {
		shell_path, command_path, err = windows_command_paths(base_path)
	} else {
		shell_path, command_path, err = unix_command_paths(base_path)
	}

	return
}


func windows_command_paths(base_path string) (shell_path, command_path string, err error) {
	shell_path, err = exec.LookPath("cmd.exe")
	command_path = base_path + string(os.PathSeparator) + "Scripts" + string(os.PathSeparator) + "activate.bat"

	return
}


func unix_command_paths(base_path string) (shell_path, command_path string, err error) {
	shell_path, err = exec.LookPath("bash")
	command_path = base_path + string(os.PathSeparator) + "bin" + string(os.PathSeparator) + "activate"

	return
}
