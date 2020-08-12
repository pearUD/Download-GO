package entity

import (
	"runtime"
	"time"
)

type config struct {
	cpu      int
	timeOut  int
	fileName string
	url      string
	path     string
}

var c *config

func LoadConfig() *config {
	onceConf.Do(func() {
		c = new(config)
		c.SetTimeOut(-1)
		c.SetCpu(runtime.NumCPU())
		c.SetFileName("main-" + time.Now().Format("2006-01-02"))
		c.SetPath("./")
	})
	return c
}

func (c *config) SetCpu(i interface{}) {
	c.cpu = i.(int)
}

func (c *config) GetCpu() int {
	return c.cpu
}

func (c *config) SetTimeOut(i interface{}) {
	c.timeOut = i.(int)
}

func (c *config) GetTimeOut() int {
	return c.timeOut
}

func (c *config) SetFileName(i interface{}) {
	c.fileName = i.(string)
}

func (c *config) GetFileName() string {
	return c.fileName
}

func (c *config) SetUrl(i interface{}) {
	c.url = i.(string)
}

func (c *config) GetUrl() string {
	return c.url
}

func (c *config) SetPath(i interface{}) {
	c.path = i.(string)
}

func (c *config) GetPath() string {
	return c.path
}
