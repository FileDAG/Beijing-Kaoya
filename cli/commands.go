package cli

import (
	"Beijing-Kaoya/swarm"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

type uploadmsg struct {
	PervRef string `json:"perv_reference"`
	CurrRef string `json:"curr_reference"`
}

func (cli *CLI) Run_Swarm() {
	err := swarm.Run()
	if err != nil {

	}
}

func (cli *CLI) Upload(file, address string) {
	//fmt.Println("in func upload")
	param1 := "@" + file
	param2 := "\"" + "Swarm-Postage-Batch-Id: " + address + "\"" + " " + "\"" + "http://localhost:1633/bzz?name=" + file + "\""
	//fmt.Println("curl" + " " + "--data-binary" + " " + param1 + " " + "-H" + " " + param2)
	cmd := exec.Command("bash", "-c", "curl"+" "+"--data-binary"+" "+param1+" "+"-H"+" "+param2)
	//_ = cmd.Run()
	out, _ := cmd.Output()
	//fmt.Println(string(out))

	t := []byte(out)
	t = t[14 : len(t)-4]
	msg := uploadmsg{PervRef: string(t), CurrRef: string(t)}
	f, _ := os.Create("msg" + ".json")
	encoder := json.NewEncoder(f)
	_ = encoder.Encode(msg)
	f.Close()

	param3 := "@" + f.Name()
	param4 := "\"" + "Swarm-Postage-Batch-Id: " + address + "\"" + " " + "\"" + "http://localhost:1633/bzz?name=" + f.Name() + "\""
	cmd2 := exec.Command("bash", "-c", "curl"+" "+"--data"+" "+param3+" "+"-H"+" "+param4)
	//_ = cmd2.Run()
	out2, _ := cmd2.Output()
	fmt.Println("the file reference is:", string(out2))
	os.Remove("msg" + ".json")

}

func (cli *CLI) Upload_Patch(patchName, oldIndex string, address string) {
	param1 := "@" + patchName
	param2 := "\"" + "Swarm-Postage-Batch-Id: " + address + "\"" + " " + "\"" + "http://localhost:1633/bzz?name=" + patchName + "\""
	//fmt.Println("curl" + " " + "--data-binary" + " " + param1 + " " + "-H" + " " + param2)
	cmd := exec.Command("bash", "-c", "curl"+" "+"--data-binary"+" "+param1+" "+"-H"+" "+param2)
	//_ = cmd.Run()
	out, _ := cmd.Output()
	//fmt.Println(string(out))

	t := []byte(out)
	t = t[14 : len(t)-4]
	msg := uploadmsg{PervRef: oldIndex, CurrRef: string(t)}
	f, _ := os.Create("msg" + ".json")
	encoder := json.NewEncoder(f)
	_ = encoder.Encode(msg)
	f.Close()

	param3 := "@" + f.Name()
	param4 := "\"" + "Swarm-Postage-Batch-Id: " + address + "\"" + " " + "\"" + "http://localhost:1633/bzz?name=" + f.Name() + "\""
	cmd2 := exec.Command("bash", "-c", "curl"+" "+"--data"+" "+param3+" "+"-H"+" "+param4)
	//_ = cmd2.Run()
	out2, _ := cmd2.Output()
	fmt.Println("the newest file reference is:", string(out2))
	os.Remove("msg" + ".json")
}

func (cli *CLI) Gen_Patch(old_file, new_file string, version_num int) {
	patchName := "patch" + strconv.Itoa(version_num)
	cmd := exec.Command("bsdiff", old_file, new_file, patchName)
	_ = cmd.Run()

	out := "the patch file patch" + strconv.Itoa(version_num) + " has been generated!"
	fmt.Println(out)
}

func (cli *CLI) Download(ref string) {
	//println("curl" + " " + "-OJ" + " " + "http://localhost:1633/bzz/" + ref + "/")
	cmd := exec.Command("bash", "-c", "curl"+" "+"-OJ"+" "+"http://localhost:1633/bzz/"+ref+"/")
	cmd.Run()

	version := 1

	for {
		file := "msg" + ".json"
		f, _ := os.Open(file)
		data, _ := ioutil.ReadAll(f)
		f.Close()
		var msg uploadmsg
		_ = json.Unmarshal(data, &msg)
		os.Remove("msg" + ".json")
		if msg.CurrRef == msg.PervRef {
			//println("curl" + " " + "-OJ" + " " + "http://localhost:1633/bzz/" + msg.Curr_ref + "/")
			cmd2 := exec.Command("bash", "-c", "curl"+" "+"-OJ"+" "+"http://localhost:1633/bzz/"+msg.CurrRef+"/")
			cmd2.Run()
			break
		} else {
			//println("curl" + " " + "-OJ" + " " + "http://localhost:1633/bzz/" + msg.Curr_ref + "/")
			cmd2 := exec.Command("bash", "-c", "curl"+" "+"-OJ"+" "+"http://localhost:1633/bzz/"+msg.CurrRef+"/")
			cmd2.Run()
			//println("curl" + " " + "-OJ" + " " + "http://localhost:1633/bzz/" + msg.Perv_ref + "/")
			cmd3 := exec.Command("bash", "-c", "curl"+" "+"-OJ"+" "+"http://localhost:1633/bzz/"+msg.PervRef+"/")
			cmd3.Run()
			version = version + 1
		}
	}
	if version > 1 {
		for i := 2; i <= version; i++ {
			cmd4 := exec.Command("bspatch", "bee"+strconv.Itoa(i-1)+".jpg", "bee"+strconv.Itoa(i)+".jpg", "patch"+strconv.Itoa(i))
			_ = cmd4.Run()
			os.Remove("bee" + strconv.Itoa(i-1) + ".jpg")
			os.Remove("patch" + strconv.Itoa(i))
		}
	}
	out := "the file bee" + strconv.Itoa(version) + ".jpg has been restored!"
	fmt.Println(out)

}
