package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"

	"golang-clean-architecture/pkg/config"
	"golang-clean-architecture/pkg/infrastructure/datastore"
	"golang-clean-architecture/pkg/infrastructure/router"
	"golang-clean-architecture/pkg/registry"
)

func main() {
	config.ReadConfig()

	db := datastore.NewDB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close()

	r := registry.NewRegistry(db)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
