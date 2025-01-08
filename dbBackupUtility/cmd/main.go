package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func initApp(app *cli.App) {
	app.Name = "Database Backup Utility"
	app.Usage = "Let's you backup a database and upload the backup to a cloud storage"

	app.Commands = []cli.Command{
		{
			Name: "backup",
			HelpName: "backup",
			Action: BackupDB,
			ArgsUsage: "-c",
			Usage: "Backups a database and stores the backup",
			Description: "Backups a database and stores the backup",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name: "config-file",
					Usage: "YAML Config File Path",
				},
				&cli.BoolFlag{
					Name: "rm-local",
					Usage: "Remove Local Backup files",
				},
				&cli.BoolFlag{
					Name: "compress",
					Usage: "Create a Gzip compressed archive",
				},
				&cli.StringFlag{
					Name: "strategy",
					Usage: "Overwrite the backup strategy: full, incremental, differential",
				},
				&cli.StringFlag{
					Name: "log-file",
					Usage: "Overwrite the log file path",
				},
			},
		},
	}
}

func main() {
	app := cli.NewApp()
	initApp(app)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}