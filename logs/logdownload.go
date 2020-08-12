package logs

import (
	"download-go/entity"
	"fmt"
	"log"
	"sync"
	"time"
)

var (
	unit      = [5]string{"byte", "KB", "MB", "GB", "TB"}
	onceL     sync.Once
	resp      entity.Resources
	content   int64
	inLoad    int64
	outTime   int
	timeLimit int
)

func LogDownload(resources entity.Resources, conf entity.Conf) {
	onceL.Do(func() {
		resp = resources
		content = resp.GetContent()
		inLoad = resp.GetInLoad()
		outTime = conf.GetTimeOut()
		timeLimit = conf.GetTimeOut()
		logMessage()
	})
	for {
		time.Sleep(time.Second * 1)
		logMessage()
	}
}

func logMessage() {
	TInLoad := resp.GetInLoad()
	s, su := changeUnit(float64(TInLoad - inLoad))
	if outTime >= 0 {
		if outTime == 0 {
			log.Fatalf("已超时\n")
		}
		if s == 0 {
			outTime--
		} else if s > 0 {
			outTime = timeLimit
		}
	}
	inLoad = TInLoad
	if content == 0 {
		t, tu := changeUnit(float64(TInLoad))
		fmt.Printf("\r 正在下载 已下载%.2f %s,当前速度 %.2f %s /s", t, tu, s, su)
		return
	}
	fmt.Printf("\r 正在下载 当前进度 %.2f%%,当前速度 %.2f %s /s", float64(TInLoad*10000/content)/100.00, s, su)
	return
}

func changeUnit(i float64) (float64, string) {
	u := 0
	for {
		if i <= 1024.00 {
			break
		}
		i = i / 1024.00
		u++
	}
	return i, unit[u]
}
