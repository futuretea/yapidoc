package conf

//YapiConfig Yapi信息
type YapiConfig struct {
	URL string
}

func readYapi() *YapiConfig {
	yapiConfig := YapiConfig{
		URL: readString("YAPI_URL", "yapi.url"),
	}
	return &yapiConfig
}

//GetYapiURL 获取Yapi信息
func GetYapiURL() string {
	yapiConfig := readYapi()
	return yapiConfig.URL
}
