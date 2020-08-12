package entity

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

type ConfigController struct {
	Conf Conf
	maps map[string]func(i interface{})
}

var (
	configC *ConfigController
)

func LoadConfigController() *ConfigController {
	onceConfC.Do(func() {
		configC = new(ConfigController)
	})
	return configC
}

func (c *ConfigController) SetMaps(inMaps map[string]func(i interface{})) {
	c.maps = inMaps
}

func (c *ConfigController) SetConf(conf Conf) {
	c.Conf = conf
}

func (c *ConfigController) GetConf() Conf {
	return c.Conf
}

func (c *ConfigController) SetCPU(i interface{}) {
	cpu, err := strconv.Atoi(i.(string))
	if err != nil {
		log.Fatalf(err.Error())
	}
	if cpu <= 0 || cpu > 100 {
		log.Println("线程数设值错误")
		return
	}
	c.Conf.SetCpu(cpu)
}

func (c *ConfigController) SetTimeOut(i interface{}) {
	timeOut, err := strconv.Atoi(i.(string))
	if err != nil {
		log.Fatalf(err.Error())
	}
	if timeOut <= 0 && timeOut != -1 {
		log.Println("超时设值错误")
		return
	}
	c.Conf.SetTimeOut(timeOut)
}

func (c *ConfigController) SetFileName(i interface{}) {
	if i == "" {
		return
	}
	c.Conf.SetFileName(i)
}

func (c *ConfigController) SetUrl(i interface{}) {
	if strings.HasPrefix(i.(string), "http://") || strings.HasPrefix(i.(string), "https://") {
		c.Conf.SetUrl(i)
		return
	}
	log.Fatalf("请设置正确的链接 当前链接：%s", i)
}

func (c *ConfigController) SetPath(i interface{}) {
	reg, err := regexp.CompilePOSIX(`[A-Za-z]:[/\\].*`)
	if err == nil && reg.MatchString(i.(string)) {
		c.Conf.SetPath(i)
		return
	}
	return
}

func (c *ConfigController) SetMap(k, v interface{}) {
	if k != nil || k != "" || v != nil || v != "" {
		if _, ok := c.maps[k.(string)]; !ok {
			log.Fatalf("使用未定义参数：%s", k)
		}
		c.maps[strings.ToLower(k.(string))](v)
	}
}
