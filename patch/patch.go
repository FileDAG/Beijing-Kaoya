package patch

import (
	"math/rand"
	"os/exec"
	"strconv"
)

func Patch(old_file, new_file string) string {
	version_num := rand.Int()
	patchName := "patch" + strconv.Itoa(version_num)
	cmd := exec.Command("bsdiff", old_file, new_file, patchName)
	_ = cmd.Run()

	return patchName
}
