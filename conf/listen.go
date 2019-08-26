package conf

import (
	"fmt"
)

//ListenConfigStruct 监听信息
type ListenConfigStruct struct {
	Host string
	Port int
}

func readListen() *ListenConfigStruct {
	listenConfig := ListenConfigStruct{
		Host: readString("LISTEN_HOST", "listen.host"),
		Port: readInt("LISTEN_PORT", "listen.port"),
	}
	return &listenConfig
}

//GetListenURL 获取监听信息
func GetListenURL() string {
	listenConfig := readListen()
	return fmt.Sprintf(
		"%s:%d",
		listenConfig.Host,
		listenConfig.Port,
	)
}
