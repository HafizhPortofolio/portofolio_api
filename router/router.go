package router

import (
	"net/http"

	"github.com/MafuSora/portofolio_db/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PATCH, OPTIONS")
			// fmt.Println(c.Request)
			// c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func NewRouter(tagsController *controller.TagsController, categoriesController *controller.CategoriesController, kegunaansController *controller.KegunaansController, profileController *controller.ProfileController, footerController *controller.FooterController, experienceController *controller.ExperienceController, portofolioController *controller.PortofolioController, certificateController *controller.CertificateController) *gin.Engine {
	router := gin.Default()
	// router.Use(corsMiddleware())
	router.Use(cors.Default())
	router.LoadHTMLGlob("templates/*")
	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"GET", "PATCH", "POST", "PUT", "DELETE"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin == "https://github.com"
	// 	},
	// 	MaxAge: 12 * time.Hour,
	// }))
	router.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	baseRouter := router.Group("/api")
	baseRouter.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	tagsRouter := baseRouter.Group("/tags")
	tagsRouter.GET("", tagsController.FindAll)
	tagsRouter.GET("/:tagId", tagsController.FindById)
	tagsRouter.POST("", tagsController.Create)
	tagsRouter.PATCH("/:tagId", tagsController.Update)
	tagsRouter.DELETE("/:tagId", tagsController.Delete)

	categoryRouter := baseRouter.Group("/category")
	categoryRouter.GET("", categoriesController.FindAll)
	categoryRouter.GET("/:categoryId", categoriesController.FindById)
	categoryRouter.GET("/category", categoriesController.FindCategory)
	categoryRouter.GET("/portofolio", categoriesController.FindPortofolio)
	categoryRouter.GET("/certificate", categoriesController.FindCertificate)
	categoryRouter.POST("", categoriesController.Create)
	categoryRouter.PATCH("/:categoryId", categoriesController.Update)
	categoryRouter.DELETE("/:categoryId", categoriesController.Delete)

	kegunaanRouter := baseRouter.Group("/kegunaan")
	kegunaanRouter.GET("", kegunaansController.FindAll)
	kegunaanRouter.GET("/:kegunaanId", kegunaansController.FindById)
	kegunaanRouter.POST("", kegunaansController.Create)
	kegunaanRouter.PATCH("/:kegunaanId", kegunaansController.Update)
	kegunaanRouter.DELETE("/:kegunaanId", kegunaansController.Delete)

	profileRouter := baseRouter.Group("/profile")
	profileRouter.GET("", profileController.FindAll)
	profileRouter.GET("/first", profileController.FindFirst)
	profileRouter.GET("/:profileId", profileController.FindById)
	profileRouter.POST("", profileController.Create)
	// profileRouter.OPTIONS("", profileController.Create)
	profileRouter.PATCH("/:profileId", profileController.Update)
	profileRouter.DELETE("/:profileId", profileController.Delete)

	footerRouter := baseRouter.Group("/footer")
	footerRouter.GET("", footerController.FindAll)
	footerRouter.GET("/first", footerController.FindFirst)
	footerRouter.GET("/:footerId", footerController.FindById)
	footerRouter.POST("", footerController.Create)
	// profileRouter.OPTIONS("", profileController.Create)
	footerRouter.PATCH("/:footerId", footerController.Update)
	footerRouter.DELETE("/:footerId", footerController.Delete)

	experienceRouter := baseRouter.Group("/experience")
	experienceRouter.GET("", experienceController.FindAll)
	experienceRouter.GET("/:experienceId", experienceController.FindById)
	experienceRouter.POST("", experienceController.Create)
	experienceRouter.PATCH("/:experienceId", experienceController.Update)
	experienceRouter.DELETE("/:experienceId", experienceController.Delete)

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
	router.Run()
	return router
}
