# Export mssql database tables to CSV

dbexport is a small utility written in GO while I am learning golang. dbexport
connects to a mssql database or azure sqldb, get the list of tables and export
all into CSV files by simply running ```dbexport``` command.

## Pre-req
Setup golang dev environment on your linux or mac. See [Install the Go
Tools](https://golang.org/doc/install#install). Then clone this repository.

## Build dbexport

```bash
git clone https://github.com/erickang7/dbexport
cd dbexport
go build

```

dbexport executable should be created.

## Configure dbconfig.yaml
dbconfig.yaml defines configuration settings such as server, database,
username, the file path to create and save CSV files. Check
```dbconfig.yaml.template``` file for sample. Note that dbexport app looks for
dbconfig.yaml file in the same local directory.

```bash
cp dbconfig.yaml.template dbconfig.yaml
vim dbconfig.yaml
```

```yaml
servername: localhost 
port: 1433 
databasename: mydb 
user: myuser 
password: IKnowThatThisIsNotABestPractice!! 
csvpath: ~/tmp/csvfiles 
```
