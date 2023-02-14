package controller

import (
	"github.com/dranikpg/dto-mapper"
	"github.com/gin-gonic/gin"
	"github.com/goapi/dtos"
	"github.com/goapi/model"
	"github.com/goonode/mogo"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func Posting(httpContext *gin.Context) {
	fileModel := mogo.NewDoc(model.File{}).(*model.File)
	if httpContext.BindJSON(&fileModel) != nil {
		httpContext.JSON(http.StatusBadRequest, model.Error{
			Title:       "Bad json",
			Description: "Bad json",
		})
		return
	}
	if fileModel.Save() != nil {
		httpContext.JSON(http.StatusInternalServerError, model.Error{
			Title:       "Database error",
			Description: "Database error",
		})
		return
	}
	httpContext.JSON(http.StatusCreated, fileModel)
}

func Get(httpContext *gin.Context) {
	fileModel := mogo.NewDoc(model.File{}).(*model.File)
	filename := httpContext.Param("filename")
	err := fileModel.FindOne(bson.M{"filename": filename}, fileModel)
	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, model.Error{
			Title:       "Database error",
			Description: "Database error",
		})
		return
	}

	mapper := dto.Mapper{}
	fileDto := dtos.File{}
	mapper.Map(&fileDto, fileModel)
	httpContext.JSON(http.StatusOK, fileDto)
}
