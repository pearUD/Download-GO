package download

import (
	"download-go/entity"
	"io"
	"log"
)

func RunByOne(r entity.Resources) {
	file := r.GetFile()
	_, err := io.CopyBuffer(file, r, buff[:])
	if err != nil {
		log.Fatalf(err.Error())
	}
}
