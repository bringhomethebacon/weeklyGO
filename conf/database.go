package conf

import config "github.com/tfcloud-go/config"

func init() {
	registerDatabaseConf()
}

func registerDatabaseConf() {
	group := config.NewOptGroup("database")
	CONF.RegisterGroup(group)

	url := config.NewStrOpt("url").WithDefault("mysql://root:root@127.0.0.1:3306/weekly")
	CONF.RegisterOpts(group, url)
}
