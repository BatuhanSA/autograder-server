package util

import (
	"path/filepath"
	"strings"

	"github.com/edulinq/autograder/internal/log"
)

const (
	UNKNOWN_VERSION  string = "???"
	UNKNOWN_HASH     string = "????????"
	VERSION_FILENAME string = "VERSION.txt"
	VERSION_JSON     string = "VERSION.json"
	DIRTY_SUFFIX     string = "dirty"
	HASH_LENGTH      int    = 8
)

type Version struct{
	Full string `json:"full-version"`
}

func GetAutograderVersion() string {
	versionPath := ShouldAbs(filepath.Join(ShouldGetThisDir(), "..", "..", VERSION_FILENAME))
	if !IsFile(versionPath) {
		log.Error("Version file does not exist.", log.NewAttr("path", versionPath))
		return UNKNOWN_VERSION
	}

	version, err := ReadFile(versionPath)
	if err != nil {
		log.Error("Failed to read the version file.", err, log.NewAttr("path", versionPath))
		return UNKNOWN_VERSION
	}

	return strings.TrimSpace(version)
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
		dirtySuffix = "-" + UNKNOWN_VERSION
	}

	if isDirty {
		dirtySuffix = "-" + DIRTY_SUFFIX
	}

	return version + "-" + hash[0:HASH_LENGTH] + dirtySuffix
}

func GetAutograderFullVersion() string {
	versionJSONPath := ShouldAbs(filepath.Join(ShouldGetThisDir(), "..", "..", VERSION_JSON))
	if !PathExists(versionJSONPath) {
		return ComputeAutograderFullVersion()
	}

	var readVersion Version

	err := JSONFromFile(versionJSONPath,&readVersion)
	if err != nil {
		log.Error("Failed to read the version JSON file.", err, log.NewAttr("path", versionJSONPath))
		return UNKNOWN_VERSION
	}

	return readVersion.Full 
}