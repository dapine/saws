package fs

import "io/ioutil"

func ReadResource(path string) ([]byte, error) {
	// TODO: Convert path from HTTP URI format to Filesystem format

	// SUPERHACK. Delete dis!!!
	if path == "/" {
		path = "index.html"
	}

	return ioutil.ReadFile(path)
}
