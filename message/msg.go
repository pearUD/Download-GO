package message

import "log"

var msgMap = make(map[string]string)

func init() {
	msgMap["-h"] = "--help\n" +
		"\t-n int \t设置线程数（默认为逻辑cpu数量）\n" +
		"\t-t int \t设置超时时间，单位秒，-1为不限制超时（默认-1）\n" +
		"\t-u string \t设置链接（必须）\n" +
		"\t-p string \t设置保存路径（默认为当前目录）\n" +
		"\t-f string \t设置保存文件名\n" +
		"\t-tx \t以当前目录下的Config.txt文件开始下载\n" +
		"\t-h \t输出命令列表\n" +
		"\t-v \t输出程序信息\n"
	msgMap["-v"] = "--version\n" +
		"\t该程序由pearud编写\n" +
		"\t  gitee：https://gitee.com/pearud\n" +
		"\t  github：https://github.com/pearud\n" +
		""
}

func Help() {
	log.Fatalf(msgMap["-h"])
}

func Version() {
	log.Fatalf(msgMap["-v"])
}
