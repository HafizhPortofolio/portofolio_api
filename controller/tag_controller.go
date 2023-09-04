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

type TagsController struct {
	tagsService service.TagsService
}

func NewTagsController(service service.TagsService) *TagsController {
	return &TagsController{
		tagsService: service,
	}
}

// create controller
func (controller *TagsController) Create(ctx *gin.Context) {
	createTagsRequest := request.CreateTagsRequest{}
	err := ctx.ShouldBindJSON(&createTagsRequest)
	helper.ErrorPanic(err)
	controller.tagsService.Create(createTagsRequest)
	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// update controller
func (controller *TagsController) Update(ctx *gin.Context) {
	updateTagsRequest := request.UpdateTagsRequest{}
	err := ctx.ShouldBindJSON(&updateTagsRequest)
	helper.ErrorPanic(err)

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	updateTagsRequest.Id = id

	controller.tagsService.Update(updateTagsRequest)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// delete controller
func (controller *TagsController) Delete(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	controller.tagsService.Delete(id)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindById controller
func (controller *TagsController) FindById(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	tagResponse := controller.tagsService.FindById(id)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAll controller
func (controller *TagsController) FindAll(ctx *gin.Context) {

	tagResponse := controller.tagsService.FindAll()

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK!",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
