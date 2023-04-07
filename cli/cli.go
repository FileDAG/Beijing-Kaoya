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
	fmt.Println("  upload -file filepath -address addr - upload a non-multi-version file, the address is the batch ID of the previous fund tx")
	fmt.Println("  gen-patch -old-file filepath -new-file filepath -version-number num - generate a patch file for a mulit-version file")
	fmt.Println("  upload-patch -patch-file filepath -old-file-index id -address addr -upload a multi-version file")
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
	/*bscommtCmd := flag.NewFlagSet("bscommit", flag.ExitOnError)
	bspatchCmd := flag.NewFlagSet("bspatch", flag.ExitOnError)
	commitCmd := flag.NewFlagSet("commit", flag.ExitOnError)
	diffCmd := flag.NewFlagSet("diff", flag.ExitOnError)
	gwCmd := flag.NewFlagSet("gw", flag.ExitOnError)
	patchCmd := flag.NewFlagSet("patch", flag.ExitOnError)
	printCmd := flag.NewFlagSet("print", flag.ExitOnError)
	watchCmd := flag.NewFlagSet("watch", flag.ExitOnError)*/
	runSwarmCmd := flag.NewFlagSet("run-swarm", flag.ExitOnError)
	uploadCmd := flag.NewFlagSet("upload", flag.PanicOnError)
	genPatchCmd := flag.NewFlagSet("gen-patch", flag.ExitOnError)
	uploadPatchCmd := flag.NewFlagSet("upload-patch", flag.ExitOnError)
	downloadCmd := flag.NewFlagSet("download", flag.ExitOnError)

	uploadFile := uploadCmd.String("file", "", "the file path to be uploaded")
	uploadAddress := uploadCmd.String("address", "", "the payment address for upload a file")

	oldFile := genPatchCmd.String("old-file", "", "the old version file path")
	newFile := genPatchCmd.String("new-file", "", "the new version file path to be uploaded")
	versionNuber := genPatchCmd.Int("version-number", -1, "the number of the new file vesion")
	patchFile := uploadPatchCmd.String("patch-file", "", "the file path of the patch file to be uploaded")
	oldIndex := uploadPatchCmd.String("old-file-index", "", "the reference of the old file")
	uploadpatchAddress := uploadPatchCmd.String("address", "", "the payment address for upload a file")

	targetReference := downloadCmd.String("reference", "", "the reference of the file to be downloaded")
	//targetNumber := downloadCmd.Int("version-number", 1, "the number of the file vesion")

	/*bscommtLocation := bscommtCmd.String("loc", "", "location to be committed")
	bspatchLocation := bspatchCmd.String("loc", "", "location of files we calc bspatch for")
	commitLocation := commitCmd.String("loc", "", "location to be committed")
	diffLocation := diffCmd.String("loc", "", "location where changes should be showed")
	patchLocation := patchCmd.String("loc", "", "location of files we calc patch for")
	printLocation := printCmd.String("loc", "", "location to be showed")
	watchLocation := watchCmd.String("loc", "", "location to be watched")*/

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
		/*case "bscommit":
			err := bscommtCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
		case "bspatch":
			err := bspatchCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
		case "commit":
			err := commitCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
		case "diff":
			err := diffCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
		case "gw":
			err := gwCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
		case "patch":
			err := patchCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
		case "print":
			err := printCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
		case "watch":
			err := watchCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
		}*/
	}
	if uploadCmd.Parsed() {
		//fmt.Println("parsed")
		cli.Upload(*uploadFile, *uploadAddress)
	}
	if runSwarmCmd.Parsed() {
		cli.Run_Swarm()
	}
	if genPatchCmd.Parsed() {
		cli.Gen_Patch(*oldFile, *newFile, *versionNuber)
	}
	if uploadPatchCmd.Parsed() {
		cli.Upload_Patch(*patchFile, *oldIndex, *uploadpatchAddress)
	}
	if downloadCmd.Parsed() {
		cli.Download(*targetReference)
	}

	/*if bscommtCmd.Parsed() {
		cli.commit(*bscommtLocation, true)
	}

	if bspatchCmd.Parsed() {
		cli.patch(*bspatchLocation, true)
	}

	if commitCmd.Parsed() {
		cli.commit(*commitLocation, false)
	}

	if diffCmd.Parsed() {
		cli.printDiff(*diffLocation)
	}

	if gwCmd.Parsed() {
		cli.gitwalker()
	}

	if patchCmd.Parsed() {
		cli.patch(*patchLocation, false)
	}

	if printCmd.Parsed() {
		cli.printFolder(*printLocation)
	}

	if watchCmd.Parsed() {
		cli.watch(*watchLocation)
	}*/
}
