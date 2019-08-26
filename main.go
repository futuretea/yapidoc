package main

import (
	"github.com/futuretea/yapidoc/conf"
	"github.com/futuretea/yapidoc/logs"
	"github.com/futuretea/yapidoc/router"
)

func main() {
	if err := router.GinEngine.Run(conf.GetListenURL()); err != nil {
		logs.Fatal("gin run error", err)
	}
}
