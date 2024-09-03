package util

import (
	// "fmt"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/edulinq/autograder/internal/log"
)

const (
	UNKNOWN_VERSION  string = "???"
	UNKNOWN_HASH     string = "????????"
	VERSION_FILENAME string = "VERSION.json"
	DIRTY_SUFFIX     string = "dirty"
	HASH_LENGTH      int    = 8
)

type Version struct{
	Short string `json:"short-version"`
	Hash string `json:"hash-version"`
	State string `json:"state-version"`
}

func GetAutograderVersion() string {
	versionPath := ShouldAbs(filepath.Join(ShouldGetThisDir(), "..", "..", VERSION_FILENAME))
	if !IsFile(versionPath) {
		log.Error("Version file does not exist.", log.NewAttr("path", versionPath))
		return UNKNOWN_VERSION
	}

	var version Version

	err := JSONFromFile(versionPath,&version)
	if err != nil {
		log.Error("Failed to read the version JSON file.", err, log.NewAttr("path", versionPath))
		return UNKNOWN_VERSION
	}

	if version.Short == "" {
		log.Error("Version file does not have a short version",log.NewAttr("path", versionPath))
		return UNKNOWN_VERSION
	}

	return strings.TrimSpace(version.Short)
}

func ComputeAutograderFullVersion() string {
	repoPath := ShouldAbs(filepath.Join(ShouldGetThisDir(), "..", ".."))

	version := GetAutograderVersion()

	hash, err := GitGetCommitHash(repoPath)
	if err != nil {
		log.Error("Failed to get commit hash.", err, log.NewAttr("path", repoPath))
		hash = UNKNOWN_HASH
	}

	dirtySuffix := ""

	isDirty, err := GitRepoIsDirtyHack(repoPath)
	if err != nil {
		dirtySuffix = UNKNOWN_VERSION
	}

	if isDirty {
		dirtySuffix = DIRTY_SUFFIX
	}

	return version + "-" + hash[0:HASH_LENGTH] + "-" + dirtySuffix
}

// func ReadAutograderFullVersionJSON() string {
// 	versionPath := ShouldAbs(filepath.Join(ShouldGetThisDir(), "..", "..", VERSION_FILENAME))
// 	if !IsFile(versionPath) {
// 		log.Error("Version file does not exist.", log.NewAttr("path", versionPath))
// 		return UNKNOWN_VERSION
// 	}

// 	var version Version

// 	err := JSONFromFile(versionPath,&version)
// 	if err != nil {
// 		log.Error("Failed to read the version JSON file.", err, log.NewAttr("path", versionPath))
// 		return UNKNOWN_VERSION
// 	}

// 	var shortVersion = GetAutograderVersion()

// 	if !(version.Hash == ""){
// 		if !(version.State == ""){
// 			return shortVersion + "-" + version.Hash + "-" + version.State
// 		}

// 		return shortVersion + "-" + version.Hash + version.State
// 	}

// }

// func WriteAutograderFullVersionJSON() string {
	// err = ToJSONFile(&versionFull,versionPath)
	// if err != nil {
	// 	log.Error("Failed to write to the JSON file", err, log.NewAttr("path", repoPath))

	// }

// }
func GetAutograderFullVersion() string {
	fullVersion := ComputeAutograderFullVersion()
	
	fmt.Println(fullVersion)
	return fullVersion
}
