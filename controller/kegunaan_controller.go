package controller

import (
	// "fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/MafuSora/portofolio_db/data/request"
	"github.com/MafuSora/portofolio_db/data/response"
	"github.com/MafuSora/portofolio_db/helper"
	"github.com/MafuSora/portofolio_db/service"
	"github.com/gin-gonic/gin"
)

type KegunaansController struct {
	kegunaansService service.KegunaansService
}

func NewKegunaansController(service service.KegunaansService) *KegunaansController {
	return &KegunaansController{
		kegunaansService: service,
	}
}

// create controller
func (controller *KegunaansController) Create(ctx *gin.Context) {
	createKegunaansRequest := request.CreateKegunaansRequest{}
	err := ctx.ShouldBindJSON(&createKegunaansRequest)
	helper.ErrorPanic(err)
	controller.kegunaansService.Create(createKegunaansRequest)
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
func (controller *KegunaansController) Update(ctx *gin.Context) {
	updateKegunaansRequest := request.UpdateKegunaansRequest{}
	err := ctx.ShouldBindJSON(&updateKegunaansRequest)
	helper.ErrorPanic(err)

	kegunaanId := ctx.Param("kegunaanId")
	id, err := strconv.Atoi(kegunaanId)
	helper.ErrorPanic(err)
	updateKegunaansRequest.Id = id

	controller.kegunaansService.Update(updateKegunaansRequest)

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
func (controller *KegunaansController) Delete(ctx *gin.Context) {
	kegunaanId := ctx.Param("kegunaanId")
	id, err := strconv.Atoi(kegunaanId)
	helper.ErrorPanic(err)

	controller.kegunaansService.Delete(id)

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
func (controller *KegunaansController) FindById(ctx *gin.Context) {
	kegunaanId := ctx.Param("kegunaanId")
	id, err := strconv.Atoi(kegunaanId)
	helper.ErrorPanic(err)

	kegunaanResponse := controller.kegunaansService.FindById(id)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   kegunaanResponse,
	}
	ctx.Header("Content-Type", "application/json")
	//ctx.Header("Access-Control-Allow-Origin", "*")

	ctx.Header("Access-Control-Allow-Origin", os.Getenv("ALLOWED_HOST"))
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAll controller
func (controller *KegunaansController) FindAll(ctx *gin.Context) {

	kegunaanResponse := controller.kegunaansService.FindAll()

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   kegunaanResponse,
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
