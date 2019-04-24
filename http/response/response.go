package response

import (
	"fmt"
	"time"

	"github.com/dapine/saws/fs"
	"github.com/dapine/saws/http/request"
	"github.com/dapine/saws/http/resource"
)

// TODO: Finish this
var StatusName map[int]string = map[int]string{
	200: "OK",
	404: "Not Found",
}

type Header struct {
	// TODO: Implement more fields
	httpVersion string
	statusCode  int
	statusName  string
	date        time.Time
	resource    resource.Resource
}

func NewHeader(httpVersion string, req request.Request) Header {
	t := time.Now().UTC()
	r, err := fs.ReadResource(req.RequestLine().RequestUri())
	if err != nil {
		sc := 404
		// TODO: Use templates to create generic codes html
		html := `<!DOCTYPE html>
		<html>
			<head>
				<title>404 Not Found</title>
			</head>
			<body>
				<h1>404 Not Found</h1>
			</body>
		</html>`

		resource404 := resource.New([]byte(html), "text/html", time.Now(), int64(len([]byte(html))), "")
		return Header{httpVersion: httpVersion, statusCode: sc, statusName: StatusName[sc], date: t, resource: resource404}
	}

	sc := 200

	return Header{httpVersion: httpVersion, statusCode: sc, statusName: StatusName[sc], date: t, resource: r}
}

func (h Header) String() string {
	d := h.date.Format(time.RFC1123)
	lm := h.resource.LastModified().UTC().Format(time.RFC1123)
	return fmt.Sprintf("%s %d %s\r\nDate: %s\r\nLast-Modified: %s\r\nContent-Length: %d\r\nContent-Type: %s\r\n\r\n%s", h.httpVersion, h.statusCode,
		h.statusName, d, lm, h.resource.Length(), h.resource.ContentType(),
		h.resource.Data())
}
