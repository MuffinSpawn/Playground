package environment

import (
	"log"
	"os/exec"
)


func CreateVirtualEnvironment(path string) (err error) {
	command := exec.Command("python", "-m", "venv", path)

	if err = command.Run(); err != nil {
		log.Fatal(err)
	}

	return
}