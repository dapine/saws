package fs

import (
	"os"
)

// func IsExt(path string) bool {
// 	if path.Ext(path) != "" {
// 		return true
// 	}

// 	return false
// }

func IsDir(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		return false, err
	}

	return stat.IsDir(), nil
}
