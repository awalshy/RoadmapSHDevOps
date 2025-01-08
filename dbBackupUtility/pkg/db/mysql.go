package database

func getMySQLDB() Database {
	return &database{
		dbtype: "mysql",
	}
}