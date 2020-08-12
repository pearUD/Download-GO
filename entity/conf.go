package entity

import "sync"

var (
	onceConf sync.Once
	onceConfC sync.Once
)

type Conf interface {
	GetCpu() int
	GetTimeOut() int
	GetFileName() string
	GetUrl() string
	GetPath() string
	SetCpu(i interface{})
	SetTimeOut(i interface{})
	SetFileName(i interface{})
	SetUrl(i interface{})
	SetPath(i interface{})
}

type ConfC interface {
	GetConf() Conf
	SetMap(k, v string)
	SetMaps(inMaps map[string]func(i interface{}))
}
