package main

import (
	"pionirx/config"
	"pionirx/db"
	"pionirx/routes"
)

func main() {
	cfg := config.Init()
	db.Init()

	e := routes.Init()
	e.Logger.Fatal(e.Start(cfg.Hostname + ":" + cfg.HostnamePort))

}
