package main

import (
	"awesomeProject/router"
	"awesomeProject/sql"
)

func main() {
	sql.InitMysql()
	router.NewRouter()
}
