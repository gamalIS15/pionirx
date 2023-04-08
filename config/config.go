package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Hostname     string
	HostnamePort string
	DbHost       string
	DbUsername   string
	DbPassword   string
	DbPort       string
	DbName       string
}

func NewConfig(hostname string, hostnamePort string, dbHost string, dbUsername string, dbPassword string, dbPort string, dbName string) *Config {
	return &Config{Hostname: hostname, HostnamePort: hostnamePort, DbHost: dbHost, DbUsername: dbUsername, DbPassword: dbPassword, DbPort: dbPort, DbName: dbName}
}

func Init() *Config {

	var connMap map[string]string
	connMap, err := godotenv.Read()

	//Load to struct
	conn := NewConfig(connMap["HOSTNAME"], connMap["HOSTNAME_PORT"], connMap["DB_HOST"], connMap["DB_USERNAME"], connMap["DB_PASSWORD"], connMap["DB_PORT"], connMap["DB_NAME"])

	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	if os.Getenv("ENVIRONMENT") == "Dev" {
		fmt.Printf("%s uses %s\n", os.Getenv("HOSTNAME"), os.Getenv("PORT"))
	}

	return conn
}
