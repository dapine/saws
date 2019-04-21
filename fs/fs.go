package fs

import (
	"errors"
	"io/ioutil"
	"log"
	"path"
)

func ReadResource(path string) ([]byte, error) {
	// TODO: Convert path from HTTP URI format to Filesystem format
	// XXX: Better relative path resolve
	// 47: /
	if path[0] == 47 {
		path = "." + path
	} else {
		path = "./" + path
	}

	var data []byte

	isdir, err := IsDir(path)
	if err != nil {
		return []byte{}, err
	}

	if isdir {
		data, err = findIndex(path)
		if err != nil {
			// if index is not found, return a directory listing
			log.Println("There is no index file. Showing dir list")
		}
	}

	return data, nil
}

func findIndex(fspath string) ([]byte, error) {
	indexFiles := []string{"index.html", "index.htm", "default.html", "default.htm", "home.html", "home.htm"}

	files, err := ioutil.ReadDir(fspath)
	if err != nil {
		return []byte{}, err
	}

	for _, f := range files {
		if inStringSlice(f.Name(), indexFiles) {
			return ioutil.ReadFile(path.Join(fspath, f.Name()))
		}
	}

	return []byte{}, errors.New("Default index file not found")
}

func inStringSlice(str string, list []string) bool {
	for _, s := range list {
		if str == s {
			return true
		}
	}

	return false
}
