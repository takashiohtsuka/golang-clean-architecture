package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"

	backendregistry "golang-clean-architecture/pkg/backend/registry"
	backendrouter "golang-clean-architecture/pkg/backend/infrastructure/router"
	"golang-clean-architecture/pkg/config"
	frontendregistry "golang-clean-architecture/pkg/frontend/registry"
	frontendrouter "golang-clean-architecture/pkg/frontend/infrastructure/router"
	"golang-clean-architecture/pkg/infrastructure/datastore"
	"golang-clean-architecture/pkg/infrastructure/router"
	"golang-clean-architecture/pkg/infrastructure/validator"
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
	br := backendregistry.NewRegistry(db)
	fr := frontendregistry.NewRegistry(db)

	e := echo.New()
	e.Validator = validator.NewCustomValidator()
	e = router.NewRouter(e, r.NewAppController())
	e = backendrouter.NewRouter(e, br.NewAppController())
	e = frontendrouter.NewRouter(e, fr.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
