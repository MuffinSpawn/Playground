package environment

import (
	"fmt"
	_ "io"
	_ "log"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)


func StartVirtualEnvironment(path string) {
	shell_path, command_path, err := os_specific_command_paths(path)
	fmt.Println(shell_path)

	if err != nil {
		panic(err)
	}

	if runtime.GOOS == "windows" {
		command := exec.Command(shell_path, "-NoExit", command_path)
		// command := exec.Command(shell_path, "-NoExit", command_path , ";", "python --version")

		command.Env = os.Environ()
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		command.Stdin = os.Stdin

		if err := command.Run(); err != nil {
			panic(err)
		}
	} else {
		if err := syscall.Exec(shell_path, []string{command_path}, os.Environ()); err != nil {
			panic(err)
		}
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
	// Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
	shell_path, err = exec.LookPath("PowerShell.exe")
	command_path = base_path + string(os.PathSeparator) + "Scripts" + string(os.PathSeparator) + "Activate.ps1"

	return
}


func unix_command_paths(base_path string) (shell_path, command_path string, err error) {
	shell_path, err = exec.LookPath("bash")
	command_path = base_path + string(os.PathSeparator) + "bin" + string(os.PathSeparator) + "activate"

	return
}
