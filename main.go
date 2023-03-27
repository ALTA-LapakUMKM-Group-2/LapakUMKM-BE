package main

import (
	"lapakUmkm/app/configs"
	"lapakUmkm/app/database"
	"lapakUmkm/app/router"
	"lapakUmkm/utils/helpers"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := configs.InitConfig()
	db := database.InitDBMysql(*cfg)
	database.InitMigration(db)

	//clientSSO
	helpers.OauthConfig.ClientID = cfg.CLIENTIDGOOGLE
	helpers.OauthConfig.ClientSecret = cfg.CLIENTSECRETGOOGLE
	helpers.ServerKey = cfg.SERVER_KEY_MIDTRANS

	e := echo.New()
	router.InitRouter(db, e)
	e.Logger.Fatal(e.Start(":8080"))
}
