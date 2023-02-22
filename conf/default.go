package conf

import config "github.com/tfcloud-go/config"

var CONF = config.CONF

func init() {
	serverConf()
}

func serverConf() {
	serverGroup := config.NewOptGroup("DEFAULT")
	CONF.RegisterGroup(serverGroup)

	hostOpt := config.NewStrOpt("host").WithDefault("127.0.0.1")
	portOpt := config.NewStrOpt("port").WithDefault("1234")
	CONF.RegisterOpts(serverGroup, hostOpt, portOpt)
}
