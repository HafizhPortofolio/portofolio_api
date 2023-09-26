package controller

import (
	"net/http"
	"os"
	"strconv"

	"github.com/MafuSora/portofolio_db/data/request"
	"github.com/MafuSora/portofolio_db/data/response"
	"github.com/MafuSora/portofolio_db/helper"
	"github.com/MafuSora/portofolio_db/service"
	"github.com/gin-gonic/gin"
)

type FooterController struct {
	footerService service.FooterService
}

func NewFooterController(service service.FooterService) *FooterController {
	return &FooterController{
		footerService: service,
	}
}

// create controller
func (controller *FooterController) Create(ctx *gin.Context) {
	CreateFooterRequest := request.CreateFootersRequest{}
	err := ctx.ShouldBindJSON(&CreateFooterRequest)
	helper.ErrorPanic(err)
	controller.footerService.Create(CreateFooterRequest)
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
func (controller *FooterController) Update(ctx *gin.Context) {
	updateFooterRequest := request.UpdateFooterRequest{}
	err := ctx.ShouldBindJSON(&updateFooterRequest)
	helper.ErrorPanic(err)

	footerId := ctx.Param("footerId")
	id, err := strconv.Atoi(footerId)
	helper.ErrorPanic(err)
	updateFooterRequest.Id = id

	controller.footerService.Update(updateFooterRequest)

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
func (controller *FooterController) Delete(ctx *gin.Context) {
	footerId := ctx.Param("footerId")
	id, err := strconv.Atoi(footerId)
	helper.ErrorPanic(err)

	controller.footerService.Delete(id)

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
func (controller *FooterController) FindById(ctx *gin.Context) {
	footerId := ctx.Param("footerId")
	id, err := strconv.Atoi(footerId)
	helper.ErrorPanic(err)

	footerResponse := controller.footerService.FindById(id)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   footerResponse,
	}
	ctx.Header("Content-Type", "application/json")
	//ctx.Header("Access-Control-Allow-Origin", "*")

	ctx.Header("Access-Control-Allow-Origin", os.Getenv("ALLOWED_HOST"))
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAll controller
func (controller *FooterController) FindAll(ctx *gin.Context) {

	footerResponse := controller.footerService.FindAll()

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   footerResponse,
	}
	ctx.Header("Content-Type", "application/json")
	//ctx.Header("Access-Control-Allow-Origin", "*")

	ctx.Header("Access-Control-Allow-Origin", os.Getenv("ALLOWED_HOST"))
	ctx.JSON(http.StatusOK, webResponse)
}

// FindFirst controller
func (controller *FooterController) FindFirst(ctx *gin.Context) {

	footerResponse := controller.footerService.FindFirst()

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   footerResponse,
	}
	ctx.Header("Content-Type", "application/json")
	//ctx.Header("Access-Control-Allow-Origin", "*")

	ctx.Header("Access-Control-Allow-Origin", os.Getenv("ALLOWED_HOST"))
	ctx.JSON(http.StatusOK, webResponse)
}
