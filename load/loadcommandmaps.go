package load

import (
	"download-go/entity"
	"download-go/message"
)

func CommandMap(command *entity.ConfigController) map[string]func(i interface{}) {
	inMap := make(map[string]func(i interface{}))
	inMap["-n"] = command.SetCPU
	inMap["-t"] = command.SetTimeOut
	inMap["-u"] = command.SetUrl
	inMap["-p"] = command.SetPath
	inMap["-f"] = command.SetFileName
	inMap["num"] = command.SetCPU
	inMap["timeout"] = command.SetTimeOut
	inMap["url"] = command.SetUrl
	inMap["path"] = command.SetPath
	inMap["filename"] = command.SetFileName
	return inMap
}

func OnceCommandMap() map[string]func() {
	inMap := make(map[string]func())
	inMap["-h"] = message.Help
	inMap["-v"] = message.Version
	return inMap
}
