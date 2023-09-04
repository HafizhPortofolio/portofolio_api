package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/m/config"
	"github.com/m/controller"
	"github.com/m/helper"
	"github.com/m/model"
	"github.com/m/repository"
	"github.com/m/router"
	"github.com/m/service"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Started Server!")
	//db
	db := config.DatabaseConnection()
	validate := validator.New()
	db.Table("tags").AutoMigrate(&model.Tags{})
	db.Table("profile").AutoMigrate(&model.Profile{})
	db.Table("portofolio").AutoMigrate(&model.Portofolio{})
	db.Table("certificate").AutoMigrate(&model.Certificate{})

	//repo
	tagRepository := repository.NewTagsRepositoryImpl(db)
	profileRepository := repository.NewProfilesRepositoryImpl(db)
	portofolioRepository := repository.NewPortofoliosRepositoryImpl(db)
	certificateRepository := repository.NewCertificatesRepositoryImpl(db)

	//service
	tagsService := service.NewTagsServiceImpl(tagRepository, validate)
	profileService := service.NewProfilesServiceImpl(profileRepository, validate)
	portofolioService := service.NewPortofoliosServiceImpl(portofolioRepository, validate)
	certificateService := service.NewCertificatesServiceImpl(certificateRepository, validate)

	//controller
	tagsController := controller.NewTagsController(tagsService)
	profileController := controller.NewProfileController(profileService)
	portofolioController := controller.NewPortofolioController(portofolioService)
	certificateController := controller.NewCertificateController(certificateService)

	//routes

	routes := router.NewRouter(tagsController, profileController, portofolioController, certificateController)
	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}
	err := server.ListenAndServe()
	helper.ErrorPanic(err)
	// fmt.Println("Hello")
}
