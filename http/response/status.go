package response

import (
	"bytes"
	"html/template"
)

// TODO: Finish this
var StatusName map[int]string = map[int]string{
	200: "OK",
	404: "Not Found",
}

const htmlStatusTemplate = `<!DOCTYPE html>
<html>
	<head>
		<title>{{.Code}} {{.Name}}</title>
	</head>
	<body>
		<h1>{{.Code}} {{.Name}}</h1>
	</body>
</html>
`

func genHtmlStatusPage(code int) []byte {
	type httpCode struct {
		Code int
		Name string
	}

	t := template.Must(template.New("htmlStatusTemplate").Parse(htmlStatusTemplate))

	var buf bytes.Buffer

	t.Execute(&buf, httpCode{Code: code, Name: StatusName[code]})

	return buf.Bytes()
}
