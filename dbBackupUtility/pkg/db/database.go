package database

import (
	"database/sql"
	"os"

	"github.com/awalshy/db-backup-utility/pkg/config"
)

type Database interface {
	Connect() error
	ValidateCredentials() error
	LoadConfig(config *config.Config) error
	Backup() error
	Disconnect()
}

type database struct {
	dbtype	 string
	host     string
	port     int
	user     string
	password string
	dbname   string
	db		 *sql.DB
	bkpFile	 *os.File
}

var (
	backupFilePath = "/tmp/dbbackup"
)

func GetDatabase(dbms string) Database {
	if dbms == "postgres" {
		return getPostgresDB()
	}
	if dbms == "mysql" {
		return getMySQLDB()
	}
	return nil
}