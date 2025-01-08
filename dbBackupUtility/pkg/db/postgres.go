package database

import "github.com/awalshy/db-backup-utility/pkg/config"

func getPostgresDB() Database {
	return &database{
		dbtype: "postgres",
	}
}

func (d *database) LoadConfig(cfg *config.Config) error {
	d.host = cfg.DBConfig.Host
	d.port = cfg.DBConfig.Port
	d.user = cfg.DBConfig.User
	d.password = cfg.DBConfig.Password
	d.dbname = cfg.DBConfig.DBName
	return nil
}

func (d *database) Connect() error {
	// Implementation here
	return nil
}

func (d *database) Backup() error {
	// Implementation here
	return nil
}

func (d *database) ValidateCredentials() error {
	return nil
}