package fs

import (
	"bytes"
	"errors"
	"html/template"
	"io/ioutil"
	"log"
	"mime"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/dapine/saws/http/resource"
)

func ReadResource(rpath string) (resource.Resource, error) {
	linkUri := rpath
	basePath, _ := filepath.Abs(".")

	if len(rpath) <= 1 && rpath[0] == 47 {
		rpath = ""
		linkUri = ""
	} else {
		rpath, _ = filepath.Abs(rpath)
	}

	rpath = path.Join(basePath, rpath)

	var data []byte
	var fn string

	isdir, err := IsDir(rpath)
	if err != nil {
		return resource.Empty, err
	}

	if isdir {
		fn, data, err = findIndex(rpath)
		if err != nil {
			// if index is not found, return a directory listing
			files, _ := ioutil.ReadDir(rpath)

			s := struct {
				Files    []os.FileInfo
				Basepath string
			}{Files: files, Basepath: linkUri}

			const tmpl = `<html>
	<body>
		{{range $i, $f := .Files}}
		<div>
			<a href="{{$.Basepath}}/{{$f.Name}}">{{$f.Name}}</a>
		</div>
		{{end}}
	</body>
</html>`

			var buf bytes.Buffer

			t := template.Must(template.New("tmpl").Parse(tmpl))
			t.Execute(&buf, s)

			r := resource.New(buf.Bytes(), "text/html", time.Now(), int64(buf.Len()), "")

			return r, nil
		}
	} else {
		data, err = ioutil.ReadFile(rpath)
		if err != nil {
			log.Println("Could not read file: ", err)
		}
	}

	f, err := os.Open(path.Join(rpath, fn))
	if err != nil {
		return resource.Empty, err
	}
	defer f.Close()

	fstat, err := f.Stat()
	if err != nil {
		return resource.Empty, err
	}

	mime := mime.TypeByExtension(path.Ext(path.Join(rpath, fn)))

	if mime == "" {
		mime = "text/plain"
	}

	r := resource.New(data, mime, fstat.ModTime(), fstat.Size(), "")

	return r, nil
}

func findIndex(fspath string) (string, []byte, error) {
	indexFiles := []string{"index.html", "index.htm", "default.html", "default.htm", "home.html", "home.htm"}

	files, err := ioutil.ReadDir(fspath)
	if err != nil {
		return "", []byte{}, err
	}

	for _, f := range files {
		if inStringSlice(f.Name(), indexFiles) {
			d, err := ioutil.ReadFile(path.Join(fspath, f.Name()))
			return f.Name(), d, err
		}
	}

	return "", []byte{}, errors.New("Default index file not found")
}

func inStringSlice(str string, list []string) bool {
	for _, s := range list {
		if str == s {
			return true
		}
	}

	return false
}
