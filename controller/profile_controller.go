package controller

import (
	"net/http"
	"strconv"

	"github.com/MafuSora/portofolio_db/data/request"
	"github.com/MafuSora/portofolio_db/data/response"
	"github.com/MafuSora/portofolio_db/helper"
	"github.com/MafuSora/portofolio_db/service"
	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	profilesService service.ProfilesService
}

func NewProfileController(service service.ProfilesService) *ProfileController {
	return &ProfileController{
		profilesService: service,
	}
}

// create controller
func (controller *ProfileController) Create(ctx *gin.Context) {
	// enablecors
	createProfilesRequest := request.CreateProfilesRequest{}
	err := ctx.ShouldBindJSON(&createProfilesRequest)
	helper.ErrorPanic(err)
	controller.profilesService.Create(createProfilesRequest)
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
func (controller *ProfileController) Update(ctx *gin.Context) {
	updateProfileRequest := request.UpdateProfilesRequest{}
	err := ctx.ShouldBindJSON(&updateProfileRequest)
	helper.ErrorPanic(err)

	profileId := ctx.Param("profileId")
	id, err := strconv.Atoi(profileId)
	helper.ErrorPanic(err)
	updateProfileRequest.Id = id

	controller.profilesService.Update(updateProfileRequest)

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
func (controller *ProfileController) Delete(ctx *gin.Context) {
	profileId := ctx.Param("profileId")
	id, err := strconv.Atoi(profileId)
	helper.ErrorPanic(err)

	controller.profilesService.Delete(id)

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
func (controller *ProfileController) FindById(ctx *gin.Context) {
	profileId := ctx.Param("profileId")
	id, err := strconv.Atoi(profileId)
	helper.ErrorPanic(err)

	profileResponse := controller.profilesService.FindById(id)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   profileResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")

	// ctx.Header("Access-Control-Allow-Origin", os.Getenv("ALLOWED_HOST"))
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAll controller
func (controller *ProfileController) FindAll(ctx *gin.Context) {

	profileResponse := controller.profilesService.FindAll()

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   profileResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")

	// ctx.Header("Access-Control-Allow-Origin", os.Getenv("ALLOWED_HOST"))
	ctx.JSON(http.StatusOK, webResponse)
}

// FindFirst controller
func (controller *ProfileController) FindFirst(ctx *gin.Context) {

	profileResponse := controller.profilesService.FindFirst()

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   profileResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Access-Control-Allow-Origin", "*")

	// ctx.Header("Access-Control-Allow-Origin", os.Getenv("ALLOWED_HOST"))
	ctx.JSON(http.StatusOK, webResponse)
}
