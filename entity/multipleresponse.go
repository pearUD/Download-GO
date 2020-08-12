package entity

import (
	"net/http"
	"os"
)

type Item struct {
	Request *http.Request
	Start   int64
	End     int64
	File    *os.File
}

type WItem struct {
	Start  int64
	Length int
	Buffer [1024 * 1024]byte
}
