package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"

	backendregistry "golang-clean-architecture/pkg/backend/registry"
	backendrouter "golang-clean-architecture/pkg/backend/infrastructure/router"
	"golang-clean-architecture/pkg/adapter/controller"
	"golang-clean-architecture/pkg/config"
	frontendregistry "golang-clean-architecture/pkg/frontend/registry"
	frontendrouter "golang-clean-architecture/pkg/frontend/infrastructure/router"
	"golang-clean-architecture/pkg/infrastructure/datastore"
	"golang-clean-architecture/pkg/infrastructure/router"
	"golang-clean-architecture/pkg/infrastructure/storage"
	"golang-clean-architecture/pkg/infrastructure/validator"
	"golang-clean-architecture/pkg/registry"
)

func main() {
	config.ReadConfig()

	e := echo.New()
	e.Validator = validator.NewCustomValidator()

	if config.C.SkipDB {
		// DBなしで起動（ヘルスチェックのみ有効）
		log.Println("SkipDB=true: starting without DB connection")
		e = router.NewRouter(e, controller.AppController{})
	} else {
		db := datastore.NewDB()
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalln(err)
		}
		defer sqlDB.Close()

		s3Client, err := storage.NewS3Client()
		if err != nil {
			log.Fatalln(err)
		}
		if err := storage.CreateBuckets(s3Client); err != nil {
			log.Fatalln(err)
		}
		storageRepo := storage.NewStorageRepository(s3Client)

		r := registry.NewRegistry(db)
		br := backendregistry.NewRegistry(db, storageRepo)
		fr := frontendregistry.NewRegistry(db)

		e = router.NewRouter(e, r.NewAppController())
		e = backendrouter.NewRouter(e, br.NewAppController())
		e = frontendrouter.NewRouter(e, fr.NewAppController())
	}

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
