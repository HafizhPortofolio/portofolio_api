package main

import (
	"net/http"

	"github.com/MafuSora/portofolio_db/config"
	"github.com/MafuSora/portofolio_db/controller"
	"github.com/MafuSora/portofolio_db/helper"
	"github.com/MafuSora/portofolio_db/model"
	"github.com/MafuSora/portofolio_db/repository"
	"github.com/MafuSora/portofolio_db/router"
	"github.com/MafuSora/portofolio_db/service"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Started Server!")
	//db
	db := config.DatabaseConnection()
	validate := validator.New()
	db.Table("tags").AutoMigrate(&model.Tags{})
	db.Table("profiles").AutoMigrate(&model.Profiles{})
	db.Table("portofolios").AutoMigrate(&model.Portofolios{})
	db.Table("certificates").AutoMigrate(&model.Certificate{})
	db.Table("categories").AutoMigrate(&model.Category{})
	db.Table("kegunaans").AutoMigrate(&model.Kegunaan{})
	db.Table("footers").AutoMigrate(&model.Footer{})
	db.Table("experiences").AutoMigrate(&model.Experience{})

	//repo
	tagRepository := repository.NewTagsRepositoryImpl(db)
	profileRepository := repository.NewProfilesRepositoryImpl(db)
	portofolioRepository := repository.NewPortofoliosRepositoryImpl(db)
	certificateRepository := repository.NewCertificatesRepositoryImpl(db)
	categoryRepository := repository.NewCategoriesRepositoryImpl(db)
	kegunaanRepository := repository.NewKegunaansRepositoryImpl(db)
	footerRepository := repository.NewFootersRepositoryImpl(db)
	experienceRepository := repository.NewExperiencesRepositoryImpl(db)

	//service
	tagsService := service.NewTagsServiceImpl(tagRepository, validate)
	profileService := service.NewProfilesServiceImpl(profileRepository, validate)
	portofolioService := service.NewPortofoliosServiceImpl(portofolioRepository, validate)
	certificateService := service.NewCertificatesServiceImpl(certificateRepository, validate)
	categoryService := service.NewCategoriesServiceImpl(categoryRepository, validate)
	kegunaanService := service.NewKegunaansServiceImpl(kegunaanRepository, validate)
	footerService := service.NewFootersServiceImpl(footerRepository, validate)
	experienceService := service.NewExperiencesServiceImpl(experienceRepository, validate)

	//controller
	tagsController := controller.NewTagsController(tagsService)
	profileController := controller.NewProfileController(profileService)
	portofolioController := controller.NewPortofolioController(portofolioService)
	certificateController := controller.NewCertificateController(certificateService)
	categoryController := controller.NewCategoriesController(categoryService)
	kegunaanController := controller.NewKegunaansController(kegunaanService)
	footerController := controller.NewFooterController(footerService)
	experienceController := controller.NewExperienceController(experienceService)

	//routes

	routes := router.NewRouter(tagsController, categoryController, kegunaanController, profileController, footerController, experienceController, portofolioController, certificateController)
	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}
	err := server.ListenAndServe()
	helper.ErrorPanic(err)
	// fmt.Println("Hello")
}
