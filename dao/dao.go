package dao

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/aspire/models"
	"github.com/aspire/utils/database"
	"gopkg.in/yaml.v3"
)

type DAO struct {
	DB *sql.DB
}

func Init() {
	var config models.DbConf

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	// Open YAML file
	file, err := os.Open(path + "\\resources\\config\\database.yaml")
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()

	// Decode YAML file to struct
	if file != nil {
		decoder := yaml.NewDecoder(file)
		if err := decoder.Decode(&config); err != nil {
			log.Fatal("error parsing yaml. " + err.Error())
		}
	}

	err = database.InitDatabase(database.DBConfig{
		DriverName:            config.Driver,
		URL:                   config.Url,
		MaxOpenConnections:    int(config.MaxOpenConn),
		MaxIdleConnections:    int(config.MaxIdleConn),
		ConnectionMaxLifetime: time.Second * time.Duration(config.ConnMaxTime),
		ConnectionMaxIdleTime: time.Second * time.Duration(config.ConnMaxTime),
	})

	if err != nil {
		log.Fatal("error initialising database. " + err.Error())
	}
}
