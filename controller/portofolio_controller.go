package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/m/data/request"
	"github.com/m/data/response"
	"github.com/m/helper"
	"github.com/m/service"
)

type PortofolioController struct {
	portofolioService service.PortofoliosService
}

func NewPortofolioController(service service.PortofoliosService) *PortofolioController {
	return &PortofolioController{
		portofolioService: service,
	}
}

// create controller
func (controller *PortofolioController) Create(ctx *gin.Context) {
	CreatePortofoliosRequest := request.CreatePortofoliosRequest{}
	err := ctx.ShouldBindJSON(&CreatePortofoliosRequest)
	helper.ErrorPanic(err)
	controller.portofolioService.Create(CreatePortofoliosRequest)
	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")

	// ctx.Header("Access-Control-Allow-Origin", os.Getenv("ALLOWED_HOST"))
	ctx.JSON(http.StatusOK, webResponse)

}

// update controller
func (controller *PortofolioController) Update(ctx *gin.Context) {
	updatePortofolioRequest := request.UpdatePortofoliosRequest{}
	err := ctx.ShouldBindJSON(&updatePortofolioRequest)
	helper.ErrorPanic(err)

	portofolioId := ctx.Param("portofolioId")
	id, err := strconv.Atoi(portofolioId)
	helper.ErrorPanic(err)
	updatePortofolioRequest.Id = id

	controller.portofolioService.Update(updatePortofolioRequest)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")

	// ctx.Header("Access-Control-Allow-Origin", os.Getenv("ALLOWED_HOST"))
	ctx.JSON(http.StatusOK, webResponse)

}

// delete controller
func (controller *PortofolioController) Delete(ctx *gin.Context) {
	portofolioId := ctx.Param("portofolioId")
	id, err := strconv.Atoi(portofolioId)
	helper.ErrorPanic(err)

	controller.portofolioService.Delete(id)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")

	// ctx.Header("Access-Control-Allow-Origin", os.Getenv("ALLOWED_HOST"))
	ctx.JSON(http.StatusOK, webResponse)
}

// FindById controller
func (controller *PortofolioController) FindById(ctx *gin.Context) {
	portofolioId := ctx.Param("portofolioId")
	id, err := strconv.Atoi(portofolioId)
	helper.ErrorPanic(err)

	portofolioResponse := controller.portofolioService.FindById(id)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   portofolioResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")

	// ctx.Header("Access-Control-Allow-Origin", os.Getenv("ALLOWED_HOST"))
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAll controller
func (controller *PortofolioController) FindAll(ctx *gin.Context) {

	portofolioResponse := controller.portofolioService.FindAll()

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   portofolioResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")

	// ctx.Header("Access-Control-Allow-Origin", os.Getenv("ALLOWED_HOST"))
	ctx.JSON(http.StatusOK, webResponse)
}

// FindData controller
func (controller *PortofolioController) FindData(ctx *gin.Context) {

	portofolioResponse := controller.portofolioService.FindData("Data Analyst")

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   portofolioResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")

	// ctx.Header("Access-Control-Allow-Origin", os.Getenv("ALLOWED_HOST"))
	ctx.JSON(http.StatusOK, webResponse)
}

// FindBack controller
func (controller *PortofolioController) FindBack(ctx *gin.Context) {

	portofolioResponse := controller.portofolioService.FindBack("Back End Development")

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   portofolioResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")

	// ctx.Header("Access-Control-Allow-Origin", os.Getenv("ALLOWED_HOST"))
	ctx.JSON(http.StatusOK, webResponse)
}

// FindFront controller
func (controller *PortofolioController) FindFront(ctx *gin.Context) {

	portofolioResponse := controller.portofolioService.FindFront("Front End Development")

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   portofolioResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")

	// ctx.Header("Access-Control-Allow-Origin", os.Getenv("ALLOWED_HOST"))
	ctx.JSON(http.StatusOK, webResponse)
}

// FindDesign controller
func (controller *PortofolioController) FindDesign(ctx *gin.Context) {

	portofolioResponse := controller.portofolioService.FindDesign("Design")

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   portofolioResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")

	// ctx.Header("Access-Control-Allow-Origin", os.Getenv("ALLOWED_HOST"))
	ctx.JSON(http.StatusOK, webResponse)
}

// FindIndustrial controller
func (controller *PortofolioController) FindIndustrial(ctx *gin.Context) {

	portofolioResponse := controller.portofolioService.FindIndustrial("Industrial Engineering")

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   portofolioResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")

	// ctx.Header("Access-Control-Allow-Origin", os.Getenv("ALLOWED_HOST"))
	ctx.JSON(http.StatusOK, webResponse)
}

// FindCategory controller
func (controller *PortofolioController) FindCategory(ctx *gin.Context) {

	portofolioResponse := controller.portofolioService.FindCategory()

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   portofolioResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")

	// ctx.Header("Access-Control-Allow-Origin", os.Getenv("ALLOWED_HOST"))
	ctx.JSON(http.StatusOK, webResponse)
}
