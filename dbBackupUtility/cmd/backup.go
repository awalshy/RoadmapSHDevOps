package main

import (
	"errors"
	"log"

	"github.com/awalshy/db-backup-utility/pkg/config"
	database "github.com/awalshy/db-backup-utility/pkg/db"
	"github.com/urfave/cli"
)

func BackupDB(c *cli.Context) error {
	if len(c.Args()) > 0 {
		return errors.New("no arguments expected, use flags")
	}
	if !c.IsSet("config-file") {
		return errors.New("the config file must be specified")
	}
	if err := backupDatabase(c); err != nil {
		return err
	}
	log.Println("Backup Successfull")
	return nil
}

func backupDatabase(c *cli.Context) error {
	configFilePath := c.String("config-file")
	log.Println("Config file " + configFilePath)
	// Config
	cfg := config.GetConfig(configFilePath)
	// Load DB
	db := database.GetDatabase(cfg.DBMS)
	db.LoadConfig(cfg)
	db.Connect()

	// Backup
	db.Backup()

	// Upload to storage if needed

	return nil
}