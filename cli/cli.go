package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct{}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  run-swarm - start a swarm node and fund some balance for this node")
	fmt.Println("  upload -file filepath -addr addr - upload a non-multi-version file, the address is the batch ID of the previous fund tx")
	fmt.Println("  gen-patch -old filepath -new filepath - generate a patch file for a mulit-version file")
	fmt.Println("  upload-patch -patch-file filepath -old-index id -addr addr - upload a multi-version file")
	fmt.Println("  write -old-index id -new filepath -addr addr - write new file on id")
	fmt.Println("  download -reference - download a file")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) Run() {
	cli.validateArgs()
	runSwarmCmd := flag.NewFlagSet("run-swarm", flag.ExitOnError)
	uploadCmd := flag.NewFlagSet("upload", flag.PanicOnError)
	genPatchCmd := flag.NewFlagSet("gen-patch", flag.ExitOnError)
	uploadPatchCmd := flag.NewFlagSet("upload-patch", flag.ExitOnError)
	writeCmd := flag.NewFlagSet("write", flag.ExitOnError)
	downloadCmd := flag.NewFlagSet("download", flag.ExitOnError)

	uploadFile := uploadCmd.String("file", "", "the file path to be uploaded")
	uploadAddress := uploadCmd.String("addr", "", "the payment address for upload a file")

	oldFile := genPatchCmd.String("old", "", "the old version file path")
	newFile := genPatchCmd.String("new", "", "the new version file path to be uploaded")
	patchFile := uploadPatchCmd.String("patch-file", "", "the file path of the patch file to be uploaded")
	oldIndex := uploadPatchCmd.String("old-index", "", "the reference of the old file")
	uploadpatchAddress := uploadPatchCmd.String("addr", "", "the payment address for upload a file")

	writeOldIndex := writeCmd.String("old", "", "the reference of the previous version")
	writeNewFile := writeCmd.String("new", "", "the new file")
	writeAddr := writeCmd.String("addr", "", "the payment address for writing")

	targetReference := downloadCmd.String("reference", "", "the reference of the file to be downloaded")
	downloadFilename := downloadCmd.String("o", "", "the output filename")

	switch os.Args[1] {
	case "upload":
		err := uploadCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "upload-patch":
		err := uploadPatchCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "gen-patch":
		err := genPatchCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "write":
		err := writeCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "download":
		err := downloadCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "run-swarm":
		err := runSwarmCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	}
	if uploadCmd.Parsed() {
		cli.Upload(*uploadFile, *uploadAddress)
	}
	if runSwarmCmd.Parsed() {
		cli.Run_Swarm()
	}
	if genPatchCmd.Parsed() {
		cli.Gen_Patch(*oldFile, *newFile)
	}
	if uploadPatchCmd.Parsed() {
		cli.Upload_Patch(*patchFile, *oldIndex, *uploadpatchAddress)
	}
	if writeCmd.Parsed() {
		cli.Write(*writeOldIndex, *writeNewFile, *writeAddr)
	}
	if downloadCmd.Parsed() {
		cli.Download(*targetReference, *downloadFilename)
	}

}
