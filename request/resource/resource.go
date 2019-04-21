package resource

import "time"

type Resource struct {
	data         []byte
	contentType  string
	lastModified time.Time
	length       int64
	encoding     string
}

var Empty Resource = Resource{data: []byte{}, contentType: "",
	lastModified: time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
	length:       0, encoding: ""}

func New(data []byte, contentType string, lastModified time.Time,
	length int64, encoding string) Resource {
	return Resource{data: data, contentType: contentType,
		lastModified: lastModified, length: length, encoding: encoding}
}

func (r Resource) Data() []byte {
	return r.data
}

func (r Resource) ContentType() string {
	return r.contentType
}

func (r Resource) LastModified() time.Time {
	return r.lastModified
}

func (r Resource) Length() int64 {
	return r.length
}

func (r Resource) Encoding() string {
	return r.encoding
}
