package load

import "download-go/entity"

func ArgCommand(arg []string) entity.Conf {
	command := entity.LoadConfigController()
	inMap := CommandMap(command)
	command.SetConf(entity.LoadConfig())
	command.SetMaps(inMap)
	for x := 0; x < len(arg); x += 2 {
		command.SetMap(arg[x], arg[x+1])
	}
	return command.GetConf()
}
