package uri

import "regexp"

// var uri = regexp.MustCompile(`^((?i)http:)//([\w\-\.]+(\:[0-9]*)?)(/(?:[\w\-\.]+(\?[\w\-\&\=]+)?)?)?$`)

var reAbsPath = regexp.MustCompile(`^(/(?:[\w\-\.]+(\?[\w\-\&\=]+)?)?)?$`)

func ValidAbsPath(absPath string) bool {
	return reAbsPath.Match([]byte(absPath))
}

func Query(absPath string) string {
	return reAbsPath.FindAllStringSubmatch(absPath, -1)[0][2]
}
