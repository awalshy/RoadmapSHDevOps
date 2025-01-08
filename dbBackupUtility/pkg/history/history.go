package history

import (
	"os"

	"github.com/awalshy/db-backup-utility/pkg/utils"
)

func openHistoryFile() *os.File {
	home, err := os.UserHomeDir()
	if err != nil {
		utils.GetLogger().Error("failed to get home directory")
		return nil
	}

	file, err := os.OpenFile(home + ".local/share/backup/history", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		utils.GetLogger().Error("failed to open history file")
		return nil
	}

	return file
}

type History struct {
	DBName		string
	Timestamp	int64
	Filename	string
	Filesize	int64
	Location	string
	BackupMode	string
}

// Logs to the history file a new backup in the format
// databasename timestamp filename filesize(in b) filelocation(local|cloud) backupmode(full,diff,incr)
func WriteBackupDetailsToHistoryFile() {
	file := openHistoryFile()

}

// read and parse file

// delete from history

// purge backups irl and file