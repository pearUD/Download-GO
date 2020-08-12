package load

import (
	"bufio"
	"download-go/entity"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ConfigFile(file string) entity.Conf {
	var l []byte
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("打开配置文件时出现错误 Error：%s", err.Error())
	}
	r := bufio.NewReader(f)
	command := entity.LoadConfigController()
	inMap := CommandMap(command)
	command.SetConf(entity.LoadConfig())
	command.SetMaps(inMap)
	for {
		l, _, err = r.ReadLine()
		if err != nil {
			if err == io.EOF {
				fmt.Println("配置文件加载完成")
				break
			}
			log.Fatalf("配置文件加载时出现错误 Error：%s", err.Error())
		}
		command.SetMap(spliceLine(string(l)))
	}
	return command.GetConf()
}

func spliceLine(line string) (k, v string) {
	for i := 0; i < len(line); i++ {
		if line[i:i+1] == ":" {
			k = line[:i]
			v = line[i+1:]
			break
		}
	}
	return strings.ToLower(strings.Trim(k, " ")), strings.Trim(v, " ")
}
