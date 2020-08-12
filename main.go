package main

import (
	"download-go/download"
	"download-go/load"
	"fmt"
	"log"
	"os"
)

func main() {
	arg := os.Args
	c := load.Load(arg)
	if c.GetUrl() == "" {
		log.Fatalf("请设置链接地址")
	}
	fmt.Println(c)
	download.Runs(c)
}
