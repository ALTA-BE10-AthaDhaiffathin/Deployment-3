package main

import (
	"fmt"

	"github.com/ALTA-BE10-AthaDhaiffathin/Deployment-3.git/config"
	"github.com/ALTA-BE10-AthaDhaiffathin/Deployment-3.git/factory"
	"github.com/ALTA-BE10-AthaDhaiffathin/Deployment-3.git/infrastructure/database/mysql"
	"github.com/labstack/echo/v4"
)

// TODO : Migrate Database

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	mysql.MigrateData(db)
	e := echo.New()

	factory.InitFactory(e, db)

	fmt.Println("Menjalankan program...")
	dsn := fmt.Sprintf(":%d", config.SERVERPORT)
	e.Logger.Fatal(e.Start(dsn))
}
