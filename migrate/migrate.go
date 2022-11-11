package main

import (
	"github.com/rzksobhy27/go-server/inits"
	"github.com/rzksobhy27/go-server/tables"
)

func init() {
	inits.LoadEnv()
	inits.ConnectDB()
}

func main() {
	inits.DB.AutoMigrate(&tables.User{})
}
