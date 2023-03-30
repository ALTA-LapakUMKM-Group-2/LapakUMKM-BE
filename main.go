package main

import (
	"lapakUmkm/app/configs"
	"lapakUmkm/app/database"
	"lapakUmkm/app/router"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := configs.InitConfig()
	db := database.InitDBMysql(*cfg)
	database.InitMigration(db)
	configs.LoadConfig(cfg)

	e := echo.New()
	router.InitRouter(db, e)
	e.Logger.Fatal(e.Start(":8080"))
}
