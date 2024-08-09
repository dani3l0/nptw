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
		"--password", config.Get().IAPassword,
	)
	log, err := cmd.Output()

	// Log
	str := "Configuring `ia` (InternetArchive) connection "
	if err == nil {
		utils.Log(str + "successful")
		utils.Log(string(log))
	} else {
		utils.Err(str + "failed")
		utils.Err(err.Error())
		utils.Err(string(log))
	}

	return err == nil
}
