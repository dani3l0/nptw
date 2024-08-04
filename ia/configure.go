package ia

import (
	"nptw/config"
	"nptw/utils"
	"os/exec"
)

func Configure() bool {
	// Setup ia
	cmd := exec.Command(
		"./bin/ia", "configure",
		"--username", config.Get().IAEmail,
		"--password", config.Get().IAEmail,
	)
	result := cmd.Run() == nil

	// Log
	str := "Configuring `ia` (InternetArchive) connection "
	if result {
		utils.Log(str + "successful")
	} else {
		utils.Err(str + "failed")
	}

	return result
}
