package main

import (
	"flag"
	"net"

	log "github.com/sirupsen/logrus"
	"github.com/tfcloud-go/config"

	"weekly-report/handler"
	"weekly-report/model"
)

var (
	cfgFile string
	CONF    = config.CONF
)

func init() {
	flag.StringVar(&cfgFile, "config", "weekly.conf", "the configuration file for weekly-server")
}

func main() {
	flag.Parse()
	_ = CONF.ParseFile(cfgFile)

	model.InitDB()

	// parse configurations
	if err := CONF.ParseFile(cfgFile); err != nil {
		log.Errorf("failed to parse %s, %v", cfgFile, err)
		return
	}

	// setup API
	r := handler.NewHandler()

	host := CONF.GetString("server", "host")
	port := CONF.GetString("server", "port")
	addr := net.JoinHostPort(host, port)
	log.Printf("Staring server on %s\n", addr)
	log.Fatal(r.Run(addr))
}
