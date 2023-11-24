package model

import (
	"fmt"
	"log"
	"server/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDatabase() (err error) {

	// load config
	const source = "%s:%s@tcp(%s:%d)/%s?readTimeout=1500ms&writeTimeout=1500ms&charset=utf8&loc=Local&parseTime=true"
	addr := config.Conf.Database.Address
	port := config.Conf.Database.Port
	user := config.Conf.Database.User
	pwd := config.Conf.Database.Password
	schema := config.Conf.Database.Schema
	dsn := fmt.Sprintf(source, user, pwd, addr, port, schema)

	// connect to mysql
	conn := mysql.Open(dsn)
	db, err = gorm.Open(conn, &gorm.Config{})
	if err == nil {
		log.Printf("Successfully connected to the mysql database at `%s:%d`", addr, port)
	}

	return
}
