package ia

import (
	"nptw/config"
	"os/exec"
)

func Configure() bool {
	cmd := exec.Command(
		"./bin/ia", "configure",
		"--username", config.Get().IAEmail,
		"--password", config.Get().IAEmail,
	)
	return cmd.Run() == nil
}
