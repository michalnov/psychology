package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

//ReadConfig return structure of configuration
func ReadConfig() (Config, error) {
	out := Config{}
	absPath, _ := filepath.Abs("config.json")
	fmt.Println("Database Opening configuration File")
	temFile, err := ioutil.ReadFile(absPath)
	if err != nil {
		fmt.Println("Reading Db config failed")
		return out, err
	}
	err = json.Unmarshal(temFile, &out)
	if err != nil {
		fmt.Println("Unmarshal config failed")
		return out, err
	}
	return out, nil
}

//InitializeDb produce strings for acces to database
func InitializeDb() (string, string, error) {
	conf, err := ReadConfig()
	if err != nil {
		return "", "", err
	}
	master := ""
	slave := ""
	master = buildDbString(conf.Database.Master)
	if conf.Database.Master.Active {
		fmt.Println("Database Master status: active")
	} else {
		fmt.Println("Database Master status: INactive")
	}
	if conf.Database.Slave.Active {
		fmt.Println("LOG SLAVE status: active")
		slave = buildDbString(conf.Database.Slave)
	}
	fmt.Println("Database config load: succes")
	return master, slave, nil
}

func buildDbString(dbIn db) string {
	return dbIn.User + ":" + dbIn.Password + "@tcp(" + dbIn.Address + ":" + dbIn.Port + ")/" + dbIn.Name
}

func isCorrectDB(dbIn db) bool {
	if dbIn.Address == "" || dbIn.User == "" || dbIn.Password == "" || dbIn.Port == "" || dbIn.Name == "" {
		return false
	}
	return true
}

//Config hold all configuration data
type Config struct {
	Server   server   `json:"server,omitempty"`
	Database database `json:"database,omitempty"`
	Log      log      `json:"log,omitempty"`
	Mailer   mail     `json:"mailer,omitempty"`
}

type mail struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

//loadConfig load configuration file and return new config structure
func loadConfig() {
	fmt.Println("loading config file")
}

type server struct {
	Folder      string `json:"folder,omitempty"`
	Port        string `json:"port,omitempty"`
	Host        string `json:"host,omitempty"`
	ShutdownKey string `json:"shutdownkey,omitempty"`
}

type database struct {
	Master db `json:"master,omitempty"`
	Slave  db `json:"slave,omitempty"`
	Root   db `json:"root,omitempty"`
}

type db struct {
	Active   bool   `json:"active,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	Port     string `json:"port,omitempty"`
	Name     string `json:"name,omitempty"`
	Address  string `json:"address,omitempty"`
}

type log struct {
	Logging  bool     `json:"logging,omitempty"`
	External external `json:"external,omitempty"`
	Database database `json:"database,omitempty"`
}

type external struct {
	Active  bool   `json:"active,omitempty"`
	Address string `json:"address,omitempty"`
}
