package ia

import (
	"nptw/config"
	"nptw/utils/log"
	"os/exec"
)

func Configure() bool {
	// Setup ia
	cmd := exec.Command(
		"./bin/ia", "configure",
		"--username", config.Get().IAEmail,
		"--password", config.Get().IAPassword,
	)
	output, err := cmd.Output()

	// Log
	str := "Configuring `ia` (InternetArchive) connection "
	if err == nil {
		log.I(str + "successful")
		log.I(string(output))
	} else {
		log.E(str + "failed")
		log.E(err.Error())
		log.E(string(output))
	}

	return err == nil
}
