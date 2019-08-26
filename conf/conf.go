package conf

import (
	"github.com/futuretea/yapidoc/logs"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	config *viper.Viper
)

func readString(keys ...string) string {
	var result = ""
	for _, key := range keys {
		result = config.GetString(key)
		if result != "" {
			return result
		}
	}
	return result
}

func readInt(keys ...string) int {
	var result = 0
	for _, key := range keys {
		result = config.GetInt(key)
		if result != 0 {
			return result
		}
	}
	return result
}

func init() {
	config = viper.New()
	//bind cmdline
	pflag.String("config", "config.yaml", "config file")
	pflag.Parse()
	err := config.BindPFlags(pflag.CommandLine)
	if err != nil {
		logs.Fatal("bind cmdline err", err)
	}

	//read file
	config.SetConfigFile(config.GetString("config"))
	if err := config.ReadInConfig(); err != nil {
		logs.Fatal("read config file err", err)
	}

	//get env
	config.AutomaticEnv()
}
