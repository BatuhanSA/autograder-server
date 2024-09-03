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

	version, err := ReadFile(versionPath)
	if err != nil {
		log.Error("Failed to read the version file.", err, log.NewAttr("path", versionPath))
		return UNKNOWN_VERSION
	}

	return strings.TrimSpace(version)
}

func ComputeAutograderFullVersion() (string,string,string) {
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

	return version, hash[0:HASH_LENGTH] , dirtySuffix
}

func ReadAutograderFullVersionJSON() (version string, hash string, state string) {
	versionPath := ShouldAbs(filepath.Join(ShouldGetThisDir(), "..", "..", VERSION_JSON))
	if !IsFile(versionPath) {
		log.Error("Version file does not exist.", log.NewAttr("path", versionPath))
		return UNKNOWN_VERSION, UNKNOWN_HASH, ""
	}

	var readVersion Version

	err := JSONFromFile(versionPath,&readVersion)
	if err != nil {
		log.Error("Failed to read the version JSON file.", err, log.NewAttr("path", versionPath))
		return UNKNOWN_VERSION, UNKNOWN_HASH, ""
	}

	var shortVersion = GetAutograderVersion()

	if !(readVersion.Hash == ""){
		if !(readVersion.State == ""){
			return shortVersion , readVersion.Hash, readVersion.State
		}

		return shortVersion , readVersion.Hash, ""
	}
	
	return shortVersion , UNKNOWN_HASH, ""
}

func WriteAutograderFullVersionJSON(version string, hash string, state string) {
	versionJSONPath := ShouldAbs(filepath.Join(ShouldGetThisDir(), "..", "..", VERSION_JSON))

	versionFull := Version{
		Short:    version,
		Hash:     hash[0:HASH_LENGTH],
		State:     state,
	}

	err := ToJSONFile(&versionFull,versionJSONPath)
	if err != nil {
		log.Error("Failed to write to the JSON file", err, log.NewAttr("path", versionJSONPath))

	}
}

func GetAutograderFullVersion() string {
	repoPath := ShouldAbs(filepath.Join(ShouldGetThisDir(), "..", "..",".git"))
	if PathExists(repoPath){
		shortVersion, hash, stateVersion := ComputeAutograderFullVersion()
		if (stateVersion == ""){
			WriteAutograderFullVersionJSON(shortVersion,hash,"")
			return shortVersion + "-" + hash
		}

		WriteAutograderFullVersionJSON(shortVersion,hash,stateVersion)
		return shortVersion + "-" + hash + "-" + stateVersion
	}

	shortVersion, hash, stateVersion := ReadAutograderFullVersionJSON()

	if (stateVersion == ""){
		return shortVersion + "-" + hash
	}

	return shortVersion + "-" + hash + "-" + stateVersion
}
