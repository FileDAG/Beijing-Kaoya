package swarm

import (
	"fmt"
	"os/exec"
	"strings"
)

var script = `
nohup bee dev > bee.log 2>&1 &  # this will directly start a bee node in dev mode(the uploaded file will be saved in dev)

sleep 5

curl http://localhost:1633 # testing whether the bee service has started

curl -s -XPOST http://localhost:1635/stamps/10000000/20  # fund this node, to make this node able to upload files

sleep 5
`

func Run() error {
	cmd := exec.Command("bash")
	cmd.Stdin = strings.NewReader(script)
	out, err := cmd.Output()
	fmt.Println(string(out))
	return err
}
