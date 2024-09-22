package ia

import (
	"nptw/utils"
	"os"
	"os/exec"
)

// YtDlp check and auto-download
func Check() bool {
	exists := Exists()
	if !exists {
		Install()
		exists = Exists()
	}
	return exists
}

// Check if ia is available
func Exists() bool {
	utils.Check("Checking Interet Archive (ia)")

	info, err := os.Stat("./bin/ia")
	ok := !os.IsNotExist(err)
	if ok {
		ok = !info.IsDir()
	}

	utils.OkFail(ok)
	return ok
}

// Download ia locally
func Install() bool {
	utils.Check("Installing Interet Archive (ia)")
	cmd := exec.Command("bash", "-c", `
		wget -O ./bin/ia https://archive.org/download/ia-pex/ia
		chmod +x ./bin/ia
	`)

	ok := cmd.Run() == nil
	utils.OkFail(ok)

	return ok
}
