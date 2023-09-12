package controller

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/m/data/request"
	"github.com/m/data/response"
	"github.com/m/helper"
	"github.com/m/service"
)

type ExperienceController struct {
	experienceService service.ExperiencesService
}

func NewExperienceController(service service.ExperiencesService) *ExperienceController {
	return &ExperienceController{
		experienceService: service,
	}
}

// create controller
func (controller *ExperienceController) Create(ctx *gin.Context) {
	CreateExperiencesRequest := request.CreateExperienceRequest{}
	err := ctx.ShouldBindJSON(&CreateExperiencesRequest)
	helper.ErrorPanic(err)
	controller.experienceService.Create(CreateExperiencesRequest)
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
func (controller *ExperienceController) Update(ctx *gin.Context) {
	updateExperiencesRequest := request.UpdateExperienceRequest{}
	err := ctx.ShouldBindJSON(&updateExperiencesRequest)
	helper.ErrorPanic(err)

	experienceId := ctx.Param("experienceId")
	id, err := strconv.Atoi(experienceId)
	helper.ErrorPanic(err)
	updateExperiencesRequest.Id = id

	controller.experienceService.Update(updateExperiencesRequest)

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
func (controller *ExperienceController) Delete(ctx *gin.Context) {
	experienceId := ctx.Param("experienceId")
	id, err := strconv.Atoi(experienceId)
	helper.ErrorPanic(err)

	controller.experienceService.Delete(id)

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
func (controller *ExperienceController) FindById(ctx *gin.Context) {
	experienceId := ctx.Param("experienceId")
	id, err := strconv.Atoi(experienceId)
	helper.ErrorPanic(err)

	experienceResponse := controller.experienceService.FindById(id)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   experienceResponse,
	}
	ctx.Header("Content-Type", "application/json")
	//ctx.Header("Access-Control-Allow-Origin", "*")

	ctx.Header("Access-Control-Allow-Origin", os.Getenv("ALLOWED_HOST"))
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAll controller
func (controller *ExperienceController) FindAll(ctx *gin.Context) {

	experienceResponse := controller.experienceService.FindAll()

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   experienceResponse,
	}
	ctx.Header("Content-Type", "application/json")
	//ctx.Header("Access-Control-Allow-Origin", "*")

	ctx.Header("Access-Control-Allow-Origin", os.Getenv("ALLOWED_HOST"))
	ctx.JSON(http.StatusOK, webResponse)
}
