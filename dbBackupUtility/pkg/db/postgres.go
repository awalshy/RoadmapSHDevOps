package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/awalshy/db-backup-utility/pkg/config"
	"github.com/awalshy/db-backup-utility/pkg/utils"
)

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
	connStr := "user=" + d.user + " password=" + d.password + " host=" + d.host + " port=" + strconv.Itoa(d.port) + " dbname=" + d.dbname
	db, err := sql.Open(d.dbtype, connStr)
	if err != nil {
		return fmt.Errorf("failed to connect to postgres db: %v", err)
	}
	d.db = db

	return nil
}

func (d *database) Disconnect() {
	d.db.Close()
}

func (d *database) ValidateCredentials() error {
	return nil
}

func (d *database) Backup() error {
	logger := utils.GetLogger()
	unixTimestamp := time.Now().Unix()
	backupFile, err := os.Create(backupFilePath + "/backup_" + strconv.FormatInt(unixTimestamp, 10) + ".sql")
	if err != nil {
		logger.Error("Failed to create backup file: %v", err)
		return err
	}
	defer backupFile.Close()

	tables, err := d.getTables()
	if err != nil {
		logger.Error("Failed to get tables: %v", err)
	}

	err = d.backupSchema(tables)
	if err != nil {
		logger.Error("Failed to backup schema: %v", err)
	}

	err = d.backupData(tables)
	if err != nil {
		logger.Error("Failed to backup data: %v", err)
	}

	return nil
}

func (d *database) getTables() ([]string, error) {
	rows, err := d.db.Query(`SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, err
		}
		tables = append(tables, tableName)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tables, nil
}

func (d *database) backupSchema(tables []string) error {
	for _, table := range tables {
		fmt.Fprintf(d.bkpFile, "\n-- Schema for table %s\n", table)

		// Get the CREATE TABLE statement for the table
		createTableQuery := fmt.Sprintf(`SELECT pg_catalog.pg_get_tabledef('%s')`, table)
		row := d.db.QueryRow(createTableQuery)

		var createTableStmt string
		if err := row.Scan(&createTableStmt); err != nil {
			return fmt.Errorf("failed to get schema for table %s: %v", table, err)
		}

		// Write the CREATE TABLE statement to the backup file
		fmt.Fprintf(d.bkpFile, "%s;\n", createTableStmt)
	}
	return nil
}

func (d *database) backupData(tables []string) error {
	for _, table := range tables {
		fmt.Fprintf(d.bkpFile, "\n-- Data for table %s\n", table)

		// Query all rows from the table
		rows, err := d.db.Query(fmt.Sprintf("SELECT * FROM %s", table))
		if err != nil {
			return fmt.Errorf("failed to query data for table %s: %v", table, err)
		}
		defer rows.Close()

		// Get the column names
		columns, err := rows.Columns()
		if err != nil {
			return fmt.Errorf("failed to get columns for table %s: %v", table, err)
		}

		// Prepare a slice to hold the row data
		values := make([]interface{}, len(columns))
		for rows.Next() {
			// Scan the row values into the slice
			for i := range values {
				values[i] = new(string) // Assuming the data is string for simplicity, adjust based on column types
			}
			if err := rows.Scan(values...); err != nil {
				return fmt.Errorf("failed to scan row data for table %s: %v", table, err)
			}

			// Build the INSERT INTO SQL statement
			var valueStrings []string
			for _, value := range values {
				val := *(value.(*string)) // Type assertion
				val = escapeSQL(val)       // Escape single quotes and other special chars
				valueStrings = append(valueStrings, fmt.Sprintf("'%s'", val))
			}

			// Write the INSERT INTO statement to the backup file
			insertStmt := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", table, strings.Join(columns, ","), strings.Join(valueStrings, ","))
			fmt.Fprintf(d.bkpFile, "%s\n", insertStmt)
		}
	}
	return nil
}