package model

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"weekly-report/conf"
)

var (
	CONF = conf.CONF
	DB   *gorm.DB
)

type MySQLOptions struct {
	Addr     string
	Username string
	Password string
	Name     string
}

func InitDB() *gorm.DB {
	dbURL := CONF.GetString("database", "url")

	opts, err := parseURL(dbURL)
	if err != nil {
		log.Errorf("mysql: parse mysql url failed: %s", err.Error())
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
		opts.Username, opts.Password, opts.Addr, opts.Name)

	DB, err = gorm.Open(mysql.Open(dataSourceName), nil)
	if err != nil {
		log.Errorf("Failed to connect to database, %s", err.Error())
	}

	if err = DB.AutoMigrate(&Student{}, &Teacher{}); err != nil {
		log.Errorf("Failed to migrate database, %s", err.Error())
	}
	return DB
}

func parseURL(dbURL string) (*MySQLOptions, error) {
	u, err := url.Parse(dbURL)
	if err != nil {
		return nil, err
	}

	if u.Scheme != "mysql" {
		err := fmt.Errorf("invalid URL scheme: %s, supports mysql only", u.Scheme)
		return nil, err
	}

	o := &MySQLOptions{}

	if u.Host == "" {
		u.Host = "127.0.0.1:3306"
	}
	o.Addr = u.Host

	if u.Path == "" {
		err := errors.New("database name is required")
		return nil, err
	}
	o.Name = strings.ReplaceAll(u.Path, "/", "")

	if p, ok := u.User.Password(); ok {
		o.Username = u.User.Username()
		o.Password = p
	} else {
		o.Password = u.User.Username()
	}

	return o, nil
}
