# Backup Utility

A Database Backup Golang tool

## Config

The cli loads a config yaml file of the following format

```yaml
dbms: postgres # mysql, mongo
dbConfig:
  host: localhost
  port: 5432
  user: root
  password: password
  dbname: defaultdb
backupStorage: local #s3, r2, cloud storage
storage:
  accessKey: akey
  privateAccessKey: pakey
  path: /path/to/backup # backupname_date_time <- will be added to the end
```

## Usage

Creating a backup

```bash
./backup-cli backup --config-file ./config.yaml
```
There are other options that can be passed to the cli

```
--compress    Compress backup to a gzip archive
--rm-local    If a cloud storage is specified and successfully uploaded, the local backup file(s) will be deleted
--strategy  <strategy>  Overwrite the strategy in the config file : full, incremental, differential
--log-file    Overwrite the log file path of the config file
```

## Features

| Feature | Status |
| ------- | ------ |
| MySQL | Nope |
| Postgres | WIP |
| MongoDB | Nope |
| Multiple Parallel dbs | Nope |
| Full bkp | WIP |
| Incremental bkp | Nope |
| Differential bkp | Nope |
| Compression | Nope |
| Localstorage | WIP |
| Cloud s3 | Nope |
| Cloud r2 | Nope |
| Cloud gcp | Nope |
| Cloud Az blob | Nope |
| File Logging | Nope |
| Notification Slack | Nope |
| Webhook call | Nope |
