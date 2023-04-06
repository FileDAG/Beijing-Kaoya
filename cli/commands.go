package cli

import (
	/*"dyaic/config"
	"dyaic/diff"
	"dyaic/ipfs"
	"dyaic/monitor"
	"dyaic/utils"*/
	/*"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"*/

	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	//"io/ioutil"
)

/*func (cli *CLI) commit(loc string, bs bool) {
	if loc == "" {
		loc = config.TempLocation
	}
	locLen := len(loc)
	err := filepath.Walk(loc, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rLoc := path[locLen:]
		repoLoc := config.RepoLocation + rLoc
		repoInfo, err := os.Stat(repoLoc)

		if utils.Exist(err) {
			if info.IsDir() {
				return nil
			}
			if info.ModTime().After(repoInfo.ModTime()) { // file has been modified, sync needed
				fmt.Println("File has been modified:", rLoc)
				patchName := repoLoc + ".patch"
				if bs {
					diff.GenBSPatch(repoLoc, path, patchName)
					diff.BSPatch(repoLoc, repoLoc, patchName, true)
				} else {
					diff.GenPatch(repoLoc, path, patchName)
					diff.Patch(repoLoc, repoLoc, patchName, true)
				}
				fmt.Println("Updated.")
				// TODO: send changes tx
				// TODO: sync changes with other nodes
			}
		} else { // new file (or folder), creation needed
			if info.IsDir() {
				fmt.Println("Creating folder:", repoLoc)
				err = os.Mkdir(repoLoc, 0755)
				if err != nil {
					return err
				}
			} else {
				fmt.Println("New file:", rLoc)
				utils.Copy(path, repoLoc)
				fmt.Println("Copied.")
				// TODO: send file tx
				// TODO: sync file with other nodes
			}
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

func (cli *CLI) gitwalker() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Panic(err)
	}
	gitwalkerDir := homedir + "/.gitwalker/"
	for d := 1; ; d++ {
		newDir := gitwalkerDir + fmt.Sprintf("%04d", d)
		oldDir := gitwalkerDir + fmt.Sprintf("%04d", d+1)
		fmt.Printf("Start Patching %04d~%04d\n", d+1, d)
		diff.GenPatchForDirectory(oldDir, newDir)
	}
}*/

type uploadmsg struct {
	Perv_ref string `json:"perv_reference"`
	Curr_ref string `json:"curr_reference"`
}

func (cli *CLI) Run_Swarm() {
	fmt.Println("in func run-swarm")
	cmd := exec.Command("bash", "dev_start.sh")
	out, _ := cmd.Output()
	fmt.Println(string(out))
}

func (cli *CLI) Upload(file, address string) {
	fmt.Println("in func upload")
	param1 := "@" + file
	param2 := "\"" + "Swarm-Postage-Batch-Id: " + address + "\"" + " " + "\"" + "http://localhost:1633/bzz?name=" + file + "\""
	fmt.Println("curl" + " " + "--data-binary" + " " + param1 + " " + "-H" + " " + param2)
	cmd := exec.Command("bash", "-c", "curl"+" "+"--data-binary"+" "+param1+" "+"-H"+" "+param2)
	//_ = cmd.Run()
	out, _ := cmd.Output()
	fmt.Println(string(out))

	t := []byte(out)
	t = t[14 : len(t)-4]
	msg := uploadmsg{Perv_ref: string(t), Curr_ref: string(t)}
	f, _ := os.Create("index_" + "patch1" + "msg" + ".json")
	encoder := json.NewEncoder(f)
	_ = encoder.Encode(msg)
	f.Close()

	param3 := "@" + f.Name()
	param4 := "\"" + "Swarm-Postage-Batch-Id: " + address + "\"" + " " + "\"" + "http://localhost:1633/bzz?name=" + f.Name() + "\""
	cmd2 := exec.Command("bash", "-c", "curl"+" "+"--data"+" "+param3+" "+"-H"+" "+param4)
	//_ = cmd2.Run()
	out2, _ := cmd2.Output()
	fmt.Println("the file reference is:", string(out2))
}

func (cli *CLI) Upload_Patch(patchName, old_index string, address string) {
	param1 := "@" + patchName
	param2 := "\"" + "Swarm-Postage-Batch-Id: " + address + "\"" + " " + "\"" + "http://localhost:1633/bzz?name=" + patchName + "\""
	fmt.Println("curl" + " " + "--data-binary" + " " + param1 + " " + "-H" + " " + param2)
	cmd := exec.Command("bash", "-c", "curl"+" "+"--data-binary"+" "+param1+" "+"-H"+" "+param2)
	//_ = cmd.Run()
	out, _ := cmd.Output()
	//fmt.Println(string(out))

	t := []byte(out)
	t = t[14 : len(t)-4]
	msg := uploadmsg{Perv_ref: old_index, Curr_ref: string(t)}
	f, _ := os.Create("index_" + patchName + "msg" + ".json")
	encoder := json.NewEncoder(f)
	_ = encoder.Encode(msg)
	f.Close()

	param3 := "@" + f.Name()
	param4 := "\"" + "Swarm-Postage-Batch-Id: " + address + "\"" + " " + "\"" + "http://localhost:1633/bzz?name=" + f.Name() + "\""
	cmd2 := exec.Command("bash", "-c", "curl"+" "+"--data"+" "+param3+" "+"-H"+" "+param4)
	//_ = cmd2.Run()
	out2, _ := cmd2.Output()
	fmt.Println("the newest file reference is:", string(out2))
}

func (cli *CLI) Gen_Patch(old_file, new_file string, version_num int) {
	patchName := "patch" + strconv.Itoa(version_num)
	cmd := exec.Command("bsdiff", old_file, new_file, patchName)
	_ = cmd.Run()
}

func (cli *CLI) Download(ref string, version int) {
	println("curl" + " " + "-OJ" + " " + "http://localhost:1633/bzz/" + ref + "/")
	cmd := exec.Command("bash", "-c", "curl"+" "+"-OJ"+" "+"http://localhost:1633/bzz/"+ref+"/")
	cmd.Run()

	for i := version; i >= 1; i-- {
		file := "index_" + "patch" + strconv.Itoa(i) + "msg" + ".json"
		f, _ := os.Open(file)

		data, _ := ioutil.ReadAll(f)
		f.Close()
		var msg uploadmsg
		_ = json.Unmarshal(data, &msg)
		if i != 1 {
			println("curl" + " " + "-OJ" + " " + "http://localhost:1633/bzz/" + msg.Curr_ref + "/")
			cmd2 := exec.Command("bash", "-c", "curl"+" "+"-OJ"+" "+"http://localhost:1633/bzz/"+msg.Curr_ref+"/")
			cmd2.Run()
			println("curl" + " " + "-OJ" + " " + "http://localhost:1633/bzz/" + msg.Perv_ref + "/")
			cmd3 := exec.Command("bash", "-c", "curl"+" "+"-OJ"+" "+"http://localhost:1633/bzz/"+msg.Perv_ref+"/")
			cmd3.Run()
		} else {
			println("curl" + " " + "-OJ" + " " + "http://localhost:1633/bzz/" + msg.Curr_ref + "/")
			cmd2 := exec.Command("bash", "-c", "curl"+" "+"-OJ"+" "+"http://localhost:1633/bzz/"+msg.Curr_ref+"/")
			cmd2.Run()
		}
	}
	if version > 1 {
		cmd4 := exec.Command("bspatch", "bee.jpg", "bee2.jpg", "patch2")
		_ = cmd4.Run()
	}
	for i := 2; i < version; i++ {
		cmd4 := exec.Command("bspatch", "bee"+strconv.Itoa(i)+".jpg", "bee"+strconv.Itoa(i+1)+".jpg", "patch"+strconv.Itoa(i+1))
		_ = cmd4.Run()
	}

}

/*func (cli *CLI) hashFile(loc string) {
	hashBegin := time.Now()
	if loc == "" {
		loc = config.TempLocation
	}
	fmt.Println(utils.Md5File(loc))
	hashEnd := time.Now()
	fmt.Println(hashEnd.Sub(hashBegin))
}

func (cli *CLI) patch(loc string, bs bool) {
	if loc == "" {
		loc = config.TempLocation
	}
	locLen := len(loc)
	err := filepath.Walk(loc, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rLoc := path[locLen:]
		repoLoc := config.RepoLocation + rLoc
		repoInfo, err := os.Stat(repoLoc)

		if utils.Exist(err) {
			if info.IsDir() {
				return nil
			}
			if !utils.SameFile(path, repoLoc) { // file has been modified, sync needed
				fmt.Println("File has been modified:", rLoc, ", new file size: ", info.Size(), ", old file size: ", repoInfo.Size())
				patchName := repoLoc + ".patch"
				if bs {
					diff.GenBSPatch(repoLoc, path, patchName)
				} else {
					diff.GenPatch(repoLoc, path, patchName)
				}
				fmt.Println("Generated patch file ", repoLoc, ".patch")

				// upload patch file to ipfs
				err = ipfs.Upload(patchName)
				if err != nil {
					return err
				}
				fmt.Println("Uploaded it to IPFS")
			}
		} else { // new file (or folder), creation needed
			if info.IsDir() {
				fmt.Println("New folder:", repoLoc)
			} else {
				fmt.Println("New file:", rLoc)
			}
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

func (cli *CLI) printDiff(loc string) {
	if loc == "" {
		loc = config.TempLocation
	}
	locLen := len(loc)
	err := filepath.Walk(loc, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rLoc := path[locLen:]
		repoLoc := config.RepoLocation + rLoc
		repoInfo, err := os.Stat(repoLoc)

		if utils.Exist(err) {
			if info.IsDir() {
				return nil
			}
			if info.ModTime().After(repoInfo.ModTime()) { // file has been modified, sync needed
				chs := diff.GenerateChanges(repoLoc, path)
				if len(chs.Item) == 0 {
					return nil
				}
				fmt.Println("File has been modified:", rLoc)
				diff.ShowDiff(repoLoc, path)
			}
		} else { // new file (or folder), creation needed
			if info.IsDir() {
				fmt.Println("New folder:", repoLoc)
			} else {
				fmt.Println("New file:", rLoc)
			}
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

func (cli *CLI) saveDiff(loc string) {
	if loc == "" {
		loc = config.TempLocation
	}
	locLen := len(loc)
	err := filepath.Walk(loc, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rLoc := path[locLen:]
		repoLoc := config.RepoLocation + rLoc
		repoInfo, err := os.Stat(repoLoc)

		if utils.Exist(err) {
			if info.IsDir() {
				return nil
			}
			if info.ModTime().After(repoInfo.ModTime()) { // file has been modified, sync needed
				chs := diff.GenerateChanges(repoLoc, path)
				if len(chs.Item) == 0 {
					return nil
				}
				fmt.Println("File has been modified:", rLoc)
				diff.SaveDyaicDiff(repoLoc, path)
			}
		} else { // new file (or folder), creation needed
			if info.IsDir() {
				fmt.Println("New folder:", repoLoc)
			} else {
				fmt.Println("New file:", rLoc)
			}
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

func (cli *CLI) printFolder(loc string) {
	if loc == "" {
		loc = config.TempLocation
	}
	err := filepath.Walk(loc, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path, info.ModTime(), info.Size())
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

func (cli *CLI) watch(loc string) {
	watcher := monitor.Watch(loc)
	defer watcher.Close()
	select {}
}*/
