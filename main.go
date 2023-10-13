package main

import (
	"fmt"

	"glamour_reserve/app/config"
	"glamour_reserve/app/database"
	"glamour_reserve/app/routes"

	"github.com/labstack/echo/v4"
)


func main() {
	appCfg, dbCfg := config.InitConfig()
	database.InitDBMysql(dbCfg)
	app := echo.New()
	routes.UserRoutes(app, database.DB)
	app.Logger.Fatal(app.Start(fmt.Sprintf(":%d", appCfg.APPPORT)))
}
