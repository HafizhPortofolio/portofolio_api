package controller

import (
	// "fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/m/data/request"
	"github.com/m/data/response"
	"github.com/m/helper"
	"github.com/m/service"
)

type CategoriesController struct {
	categoriesService service.CategoriesService
}

func NewCategoriesController(service service.CategoriesService) *CategoriesController {
	return &CategoriesController{
		categoriesService: service,
	}
}

// create controller
func (controller *CategoriesController) Create(ctx *gin.Context) {
	createCategoriesRequest := request.CreateCategoriesRequest{}
	err := ctx.ShouldBindJSON(&createCategoriesRequest)
	helper.ErrorPanic(err)
	controller.categoriesService.Create(createCategoriesRequest)
	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	//ctx.Header("Access-Control-Allow-Origin", "*")

	ctx.Header("Access-Control-Allow-Origin", os.Getenv("ALLOWED_HOST"))
	ctx.JSON(http.StatusOK, webResponse)

}

// update controller
func (controller *CategoriesController) Update(ctx *gin.Context) {
	updateCategoriesRequest := request.UpdateCategoriesRequest{}
	err := ctx.ShouldBindJSON(&updateCategoriesRequest)
	helper.ErrorPanic(err)

	categoryId := ctx.Param("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.ErrorPanic(err)
	updateCategoriesRequest.Id = id

	controller.categoriesService.Update(updateCategoriesRequest)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	//ctx.Header("Access-Control-Allow-Origin", "*")

	ctx.Header("Access-Control-Allow-Origin", os.Getenv("ALLOWED_HOST"))
	ctx.JSON(http.StatusOK, webResponse)

}

// delete controller
func (controller *CategoriesController) Delete(ctx *gin.Context) {
	categoryId := ctx.Param("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.ErrorPanic(err)

	controller.categoriesService.Delete(id)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	//ctx.Header("Access-Control-Allow-Origin", "*")

	ctx.Header("Access-Control-Allow-Origin", os.Getenv("ALLOWED_HOST"))
	ctx.JSON(http.StatusOK, webResponse)
}

// FindById controller
func (controller *CategoriesController) FindById(ctx *gin.Context) {
	categoryId := ctx.Param("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.ErrorPanic(err)

	categoryResponse := controller.categoriesService.FindById(id)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   categoryResponse,
	}
	ctx.Header("Content-Type", "application/json")
	//ctx.Header("Access-Control-Allow-Origin", "*")

	ctx.Header("Access-Control-Allow-Origin", os.Getenv("ALLOWED_HOST"))
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAll controller
func (controller *CategoriesController) FindAll(ctx *gin.Context) {

	categoryResponse := controller.categoriesService.FindAll()

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   categoryResponse,
	}
	ctx.Header("Content-Type", "application/json")
	// allowedOrigins := [2]string{"http://127.0.0.1:3000", "http://localhost:3000"}
	// origin := ctx
	// fmt.Println()
	// log.Info().Msg(origin)
	// if allowedOrigins.includes(origin) {
	// 	ctx.setHeader("Access-Control-Allow-Origin", origin)
	// }
	//ctx.Header("Access-Control-Allow-Origin", "*")

	ctx.Header("Access-Control-Allow-Origin", os.Getenv("ALLOWED_HOST"))
	ctx.JSON(http.StatusOK, webResponse)
}

// FindCategory controller
func (controller *CategoriesController) FindCategory(ctx *gin.Context) {

	categoryResponse := controller.categoriesService.FindCategory()

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   categoryResponse,
	}
	ctx.Header("Content-Type", "application/json")
	// allowedOrigins := [2]string{"http://127.0.0.1:3000", "http://localhost:3000"}
	// origin := ctx
	// fmt.Println()
	// log.Info().Msg(origin)
	// if allowedOrigins.includes(origin) {
	// 	ctx.setHeader("Access-Control-Allow-Origin", origin)
	// }
	//ctx.Header("Access-Control-Allow-Origin", "*")

	ctx.Header("Access-Control-Allow-Origin", os.Getenv("ALLOWED_HOST"))
	ctx.JSON(http.StatusOK, webResponse)
}

// FindPortofolio controller
func (controller *CategoriesController) FindPortofolio(ctx *gin.Context) {

	categoryResponse := controller.categoriesService.FindPortofolio("Portofolio")

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   categoryResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")

	// ctx.Header("Access-Control-Allow-Origin", os.Getenv("ALLOWED_HOST"))
	ctx.JSON(http.StatusOK, webResponse)
}

// FindCertificate controller
func (controller *CategoriesController) FindCertificate(ctx *gin.Context) {

	categoryResponse := controller.categoriesService.FindCertificate("Certificate")

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   categoryResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")

	// ctx.Header("Access-Control-Allow-Origin", os.Getenv("ALLOWED_HOST"))
	ctx.JSON(http.StatusOK, webResponse)
}
