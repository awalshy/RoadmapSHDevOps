package database

import "github.com/awalshy/db-backup-utility/pkg/config"

type Database interface {
	Connect() error
	ValidateCredentials() error
	LoadConfig(config *config.Config) error
	Backup() error
}

type database struct {
	dbtype	 string
	host     string
	port     int
	user     string
	password string
	dbname   string
}

var (
	backupFilePath = "/tmp/dbbackup"
)

func GetDatabase(dbms string) Database {
	if dbms == "postgres" {
		return getPostgresDB()
	}
	return nil
}