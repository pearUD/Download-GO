package load

import (
	"download-go/entity"
	"log"
	"strings"
)

func Load(arg []string) entity.Conf {
	maps := OnceCommandMap()
	if len(arg) > 2 {
		if !find(arg, maps) && !findTx(arg, "-tx") {
			return ArgCommand(arg[1:])
		} else {
			log.Fatalf("参数错误 当前参数：%s", arg)
		}
	} else {
		if strings.ToLower(arg[1]) == "-tx" {
			return ConfigFile("./Config.txt")
		}
		if v, ok := maps[strings.ToLower(arg[1])]; ok {
			v()
		}
		log.Fatalf("参数错误 当前参数：%s", arg)
	}
	return nil
}

func find(arg []string, maps map[string]func()) (b bool) {
	for _, i := range arg {
		if _, ok := maps[strings.ToLower(i)]; ok {
			return true
		}
	}
	return
}

func findTx(arg []string, c string) (b bool) {
	for _, i := range arg {
		if strings.ToLower(i) == c {
			return true
		}
	}
	return
}
