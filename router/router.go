package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m/controller"
)

func NewRouter(tagsController *controller.TagsController, profileController *controller.ProfileController, portofolioController *controller.PortofolioController, certificateController *controller.CertificateController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome To API")
	})
	baseRouter := router.Group("/api")

	tagsRouter := baseRouter.Group("/tags")
	tagsRouter.GET("", tagsController.FindAll)
	tagsRouter.GET("/:tagId", tagsController.FindById)
	tagsRouter.POST("", tagsController.Create)
	tagsRouter.PATCH("/:tagId", tagsController.Update)
	tagsRouter.DELETE("/:tagId", tagsController.Delete)

	profileRouter := baseRouter.Group("/profile")
	profileRouter.GET("", profileController.FindAll)
	profileRouter.GET("/:profileId", profileController.FindById)
	profileRouter.POST("", profileController.Create)
	profileRouter.PATCH("/:profileId", profileController.Update)
	profileRouter.DELETE("/:profileId", profileController.Delete)

	portofolioRouter := baseRouter.Group("/portofolio")
	portofolioRouter.GET("", portofolioController.FindAll)
	portofolioRouter.GET("/:portofolioId", portofolioController.FindById)
	portofolioRouter.GET("/data", portofolioController.FindData)
	portofolioRouter.GET("/back", portofolioController.FindBack)
	portofolioRouter.GET("/front", portofolioController.FindFront)
	portofolioRouter.GET("/design", portofolioController.FindDesign)
	portofolioRouter.GET("/industrial", portofolioController.FindIndustrial)
	portofolioRouter.GET("/category", portofolioController.FindCategory)
	portofolioRouter.POST("", portofolioController.Create)
	portofolioRouter.PATCH("/:portofolioId", portofolioController.Update)
	portofolioRouter.DELETE("/:portofolioId", portofolioController.Delete)

	certificateRouter := baseRouter.Group("/certificate")
	certificateRouter.GET("", certificateController.FindAll)
	certificateRouter.GET("/:certificateId", certificateController.FindById)
	certificateRouter.GET("/soft", certificateController.FindSoftSkill)
	certificateRouter.GET("/industrial", certificateController.FindIndustrial)
	certificateRouter.GET("/programming", certificateController.FindProgramming)
	certificateRouter.GET("/language", certificateController.FindLanguage)
	certificateRouter.GET("/event", certificateController.FindEvent)
	certificateRouter.GET("/category", certificateController.FindCategory)
	certificateRouter.POST("", certificateController.Create)
	certificateRouter.PATCH("/:certificateId", certificateController.Update)
	certificateRouter.DELETE("/:certificateId", certificateController.Delete)

	return router
}
