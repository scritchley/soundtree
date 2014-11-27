package main

import (
	"encoding/json"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"path"
	"time"
)

const (
	METAFILENAME string      = "metadata.json"
	DIRECTORY    string      = ".soundtree"
	PERMS        os.FileMode = 0755
)

type Metadata struct {
	LastModified time.Time              `json:"lastModified"`
	CreatedAt    time.Time              `json:"createdAt"`
	TrackedFiles map[string]TrackedFile `json:"trackedFiles"`
	Branches     []Branch               `json:"branches"`
}

type TrackedFile struct {
	Added time.Time `json:"added"`
}

type Branch struct {
	Commits []Commit `json:"commits"`
}

type Commit struct {
}

func LoadSoundTreeMetadata() Metadata {

	var soundTreeMetadata Metadata

	dat, err := ioutil.ReadFile(path.Join(pwd, DIRECTORY, METAFILENAME))
	if err != nil {
		color.Red("not a soundtree directory, use 'soundtree -init' to start a new project")
	}

	json.Unmarshal(dat, &soundTreeMetadata)
	check(err)

	return soundTreeMetadata

}

func CreateSoundTreeMetadata() {

	dat, err := ioutil.ReadFile(path.Join(pwd, DIRECTORY, METAFILENAME))
	if err == nil {
		color.Red("soundtree already initialised")
		return
	}

	soundTreeMetadata := Metadata{
		LastModified: time.Now(),
		CreatedAt:    time.Now(),
		TrackedFiles: make(map[string]TrackedFile),
	}
	// Create soundtree metadata json
	dat, err = json.Marshal(soundTreeMetadata)
	check(err)

	// Create the soundtree hidden folder
	os.Mkdir(path.Join(pwd, DIRECTORY), PERMS)

	err = ioutil.WriteFile(path.Join(pwd, DIRECTORY, METAFILENAME), dat, PERMS)
	check(err)

}

func (s Metadata) Save() {

	dat, err := json.Marshal(s)
	check(err)
	err = ioutil.WriteFile(path.Join(pwd, DIRECTORY, METAFILENAME), dat, PERMS)
	check(err)

}

func (s Metadata) AddFile(pathName string, newFile TrackedFile) {

	if _, ok := s.TrackedFiles[pathName]; ok {
	} else {
		color.Blue("added: " + pathName)
		s.TrackedFiles[pathName] = newFile
	}

}
