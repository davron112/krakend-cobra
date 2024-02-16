package main

import (
	"os"

	cmd "github.com/davron112/krakend-cobra/v2"
	viper "github.com/davron112/krakend-viper/v2"
	"github.com/davron112/lura/v2/config"
	"github.com/davron112/lura/v2/logging"
	"github.com/davron112/lura/v2/proxy"
	"github.com/davron112/lura/v2/router/gin"
)

func main() {
	cmd.Execute(viper.New(), func(serviceConfig config.ServiceConfig) {
		logger, _ := logging.NewLogger("DEBUG", os.Stdout, "")
		gin.DefaultFactory(proxy.DefaultFactory(logger), logger).New().Run(serviceConfig)
	})
}
