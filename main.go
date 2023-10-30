package main

import (
	"fmt"

	"glamour_reserve/app/config"
	"glamour_reserve/app/database"
	"glamour_reserve/app/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	appCfg, dbCfg := config.InitConfig()
	database.InitDBMysql(dbCfg)

	app := echo.New()
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
	}))

	app.Use(middleware.CORS())
	app.Pre(middleware.RemoveTrailingSlash())

	routes.UserRoutes(app, database.DB)
	routes.ServicesRoutes(app, database.DB)
	routes.BookingRoutes(app, database.DB)
	
	app.Logger.Fatal(app.Start(fmt.Sprintf(":%d", appCfg.APPPORT)))
}
