package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kaoya/patch"
	"kaoya/swarm"
	"kaoya/utils"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
)

const BZZURL = "http://localhost:1633/bzz/"

type uploadmsg struct {
	PrevRef string `json:"perv_reference"`
	CurRef  string `json:"curr_reference"`
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
	msg := uploadmsg{PrevRef: string(t), CurRef: string(t)}
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

func (cli *CLI) Upload_Patch(patchName, old_index string, address string) {
	param1 := "@" + patchName
	param2 := "\"" + "Swarm-Postage-Batch-Id: " + address + "\"" + " " + "\"" + "http://localhost:1633/bzz?name=" + patchName + "\""
	//fmt.Println("curl" + " " + "--data-binary" + " " + param1 + " " + "-H" + " " + param2)
	cmd := exec.Command("bash", "-c", "curl"+" "+"--data-binary"+" "+param1+" "+"-H"+" "+param2)
	//_ = cmd.Run()
	out, _ := cmd.Output()
	//fmt.Println(string(out))

	t := []byte(out)
	t = t[14 : len(t)-4]
	msg := uploadmsg{PrevRef: old_index, CurRef: string(t)}
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

func (cli *CLI) Gen_Patch(old_file, new_file string) {
	patch.Patch(old_file, new_file)
}

func (cli *CLI) Write(oldIndex, newFile, addr string) {
	oldName := strconv.Itoa(rand.Int())
	cli.Download(oldIndex, oldName)
	patchName := patch.Patch(oldName, newFile)
	os.Remove(oldName)
	cli.Upload_Patch(patchName, oldIndex, addr)
	os.Remove(patchName)
}

func (cli *CLI) Download(ref, filename string) {
	//println("curl" + " " + "-OJ" + " " + "http://localhost:1633/bzz/" + ref + "/")
	cmd := exec.Command("bash", "-c", "curl"+" "+"-OJ"+" "+"http://localhost:1633/bzz/"+ref+"/")
	cmd.Run()

	var increments []string

	for {
		f, _ := os.Open("msg.json")
		data, _ := ioutil.ReadAll(f)
		f.Close()
		var msg uploadmsg
		_ = json.Unmarshal(data, &msg)
		os.Remove("msg.json")
		if msg.CurRef == msg.PrevRef {
			//println("curl" + " " + "-OJ" + " " + "http://localhost:1633/bzz/" + msg.CurRef + "/")
			increment, _ := utils.DownloadFromUrl(BZZURL + msg.CurRef + "/")
			increments = append(increments, increment)
			break
		} else {
			//println("curl" + " " + "-OJ" + " " + "http://localhost:1633/bzz/" + msg.CurRef + "/")
			increment, _ := utils.DownloadFromUrl(BZZURL + msg.CurRef + "/")
			increments = append(increments, increment)
			//println("curl" + " " + "-OJ" + " " + "http://localhost:1633/bzz/" + msg.PrevRef + "/")
			cmd3 := exec.Command("bash", "-c", "curl"+" "+"-OJ"+" "+"http://localhost:1633/bzz/"+msg.PrevRef+"/")
			cmd3.Run()
		}
	}

	n := len(increments)

	for i := n - 2; i >= 0; i-- {
		NewName := strconv.Itoa(rand.Int())
		cmd = exec.Command("bspatch", increments[i+1], NewName, increments[i])
		fmt.Println("bspatch", increments[i+1], NewName, increments[i])
		cmd.Run()
		os.Remove(increments[i])
		os.Remove(increments[i+1])
		increments[i] = NewName
	}

	cmd = exec.Command("mv", increments[0], filename)
	cmd.Run()
	

	//if version > 1 {
	//	for i := 2; i <= version; i++ {
	//		cmd4 := exec.Command("bspatch", "bee"+strconv.Itoa(i-1)+".jpg", "bee"+strconv.Itoa(i)+".jpg", "patch"+strconv.Itoa(i))
	//		_ = cmd4.Run()
	//		os.Remove("bee" + strconv.Itoa(i-1) + ".jpg")
	//		os.Remove("patch" + strconv.Itoa(i))
	//	}
	//}
	//out := "the file bee" + strconv.Itoa(version) + ".jpg has been restored!"
	//fmt.Println(out)

}
