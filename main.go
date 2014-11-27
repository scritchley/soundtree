package main

import (
	"os"
	"path"
	"path/filepath"
	"time"

	"gopkg.in/alecthomas/kingpin.v1"
)

var (
	pwd, _     = os.Getwd()
	initialise = kingpin.Command("init", "initialise a new soundtree directory")

	add      = kingpin.Command("add", "add files to be tracked by soundtree")
	addFiles = add.Arg("files", "specify which files should be added").Default("*").String()

	branch     = kingpin.Command("branch", "create a new branch of this soundtree")
	branchName = branch.Arg("name", "name of the new branch").String()

	switchTo     = kingpin.Command("switch", "switch to a branch of this soundtree")
	switchToName = switchTo.Arg("name", "name of the branch").String()
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func walker(metadata Metadata, matchTerm string) filepath.WalkFunc {

	return func(pathName string, info os.FileInfo, err error) error {

		_, fileName := path.Split(pathName)

		matchFile, err := path.Match(matchTerm, fileName)
		if err != nil {
			panic(err)
		}

		if matchFile != true {
			return nil
		}

		metadata.AddFile(pathName, TrackedFile{
			Added: time.Now(),
		})

		return nil

	}

}

func main() {

	var pwd, err = os.Getwd()
	if err != nil {
		panic(err)
	}

	kingpin.Version("0.0.1")
	kingpin.Parse()

	switch kingpin.Parse() {
	case "init":
		CreateSoundTreeMetadata()
	case "add":
		metadata := LoadSoundTreeMetadata()
		filepath.Walk(pwd, walker(metadata, *addFiles))
		metadata.Save()
	}

}
