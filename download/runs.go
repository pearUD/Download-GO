package download

import (
	"download-go/entity"
	"download-go/logs"
	"download-go/response"
	"log"
)

var (
	resources entity.Resources
	buff      = [buffer]byte{}
)

const buffer = 1024

func Runs(conf entity.Conf) {
	resources = entity.GteResources()
	resources.ResourcesInit(conf)
	request := response.CreateRequest(conf.GetUrl(), "GET", nil)
	responses := response.GetResponse(request)
	maps := make(map[string]string)
	maps["Accept-Ranges"] = "bytes"
	c := response.Condition(responses, maps)
	resources.SetResponse(responses)
	resources.SetContent(responses.ContentLength)
	go logs.LogDownload(resources, conf)
	if !c || responses.ContentLength == 0 {
		log.Print("该链接不支持并发下载...")
		RunByOne(resources)
	} else {
		RunByMultiple(resources, conf)
	}
	defer resources.GetResponse().Body.Close()
	defer func() { _ = resources.GetFile().Close() }()
}
