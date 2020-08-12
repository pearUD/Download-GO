package entity

import (
	"download-go/files"
	"net/http"
	"os"
	"sync"
)

type Resources interface {
	ResourcesInit(conf Conf)
	GetUrl() string
	SetResponse(response *http.Response)
	GetResponse() *http.Response
	GetFile() *os.File
	SetContent(i int64)
	GetContent() int64
	SetInLoad(i int64)
	GetInLoad() int64
	Read(p []byte) (n int, err error)
}

type resources struct {
	url      string
	response *http.Response
	file     *os.File
	content  int64
	inLoad   int64
}

var (
	onceR sync.Once
	r     *resources
)

func GteResources() *resources {
	onceR.Do(func() {
		r = new(resources)
		r.SetInLoad(0)
	})
	return r
}

func (r *resources) ResourcesInit(conf Conf) {
	r.url = conf.GetUrl()
	r.file = files.OpenFile(conf.GetPath(), conf.GetFileName())
}

func (r *resources) GetUrl() string {
	return r.url
}

func (r *resources) SetResponse(response *http.Response) {
	r.response = response
}

func (r *resources) GetResponse() *http.Response {
	return r.response
}

func (r *resources) GetFile() *os.File {
	return r.file
}

func (r *resources) SetContent(i int64) {
	r.content = i
}

func (r *resources) GetContent() int64 {
	return r.content
}

func (r *resources) SetInLoad(i int64) {
	r.inLoad = i
}

func (r *resources) GetInLoad() int64 {
	return r.inLoad
}

func (r *resources) Read(p []byte) (n int, err error) {
	n, err = r.response.Body.Read(p)
	r.inLoad += int64(n)
	return
}
