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

type CertificateController struct {
	certificateService service.CertificatesService
}

func NewCertificateController(service service.CertificatesService) *CertificateController {
	return &CertificateController{
		certificateService: service,
	}
}

// create controller
func (controller *CertificateController) Create(ctx *gin.Context) {
	CreateCertificatesRequest := request.CreateCertificatesRequest{}
	err := ctx.ShouldBindJSON(&CreateCertificatesRequest)
	helper.ErrorPanic(err)
	controller.certificateService.Create(CreateCertificatesRequest)
	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, webResponse)

}

// update controller
func (controller *CertificateController) Update(ctx *gin.Context) {
	updateCertificatesRequest := request.UpdateCertificatesRequest{}
	err := ctx.ShouldBindJSON(&updateCertificatesRequest)
	helper.ErrorPanic(err)

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	updateCertificatesRequest.Id = id

	controller.certificateService.Update(updateCertificatesRequest)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, webResponse)

}

// delete controller
func (controller *CertificateController) Delete(ctx *gin.Context) {
	certificateId := ctx.Param("certificateId")
	id, err := strconv.Atoi(certificateId)
	helper.ErrorPanic(err)

	controller.certificateService.Delete(id)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindById controller
func (controller *CertificateController) FindById(ctx *gin.Context) {
	certificateId := ctx.Param("certificateId")
	id, err := strconv.Atoi(certificateId)
	helper.ErrorPanic(err)

	certificateResponse := controller.certificateService.FindById(id)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   certificateResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAll controller
func (controller *CertificateController) FindAll(ctx *gin.Context) {

	certificateResponse := controller.certificateService.FindAll()

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   certificateResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindSoftSkill controller
func (controller *CertificateController) FindSoftSkill(ctx *gin.Context) {

	certificateResponse := controller.certificateService.FindSoftSkill("Soft Skill")

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   certificateResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindIndustrial controller
func (controller *CertificateController) FindIndustrial(ctx *gin.Context) {

	certificateResponse := controller.certificateService.FindIndustrial("Industrial Engineering")

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   certificateResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindProgramming controller
func (controller *CertificateController) FindProgramming(ctx *gin.Context) {

	certificateResponse := controller.certificateService.FindProgramming("Programming")

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   certificateResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindLanguage controller
func (controller *CertificateController) FindLanguage(ctx *gin.Context) {

	certificateResponse := controller.certificateService.FindLanguage("Language")

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   certificateResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindEvent controller
func (controller *CertificateController) FindEvent(ctx *gin.Context) {

	certificateResponse := controller.certificateService.FindEvent("Event")

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   certificateResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindCategory controller
func (controller *CertificateController) FindCategory(ctx *gin.Context) {

	certificateResponse := controller.certificateService.FindCategory()

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   certificateResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, webResponse)
}
