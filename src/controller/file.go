package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/goonode/mogo"
	"go.mongodb.org/mongo-driver/bson"
	"goapi/model"
	"net/http"
)

func Posting(httpContext *gin.Context) {
	model := mogo.NewDoc(model.File{}).(*model.File)
	httpContext.BindJSON(&model)
	model.Save()
	httpContext.JSON(http.StatusCreated, model)
}

func Getting(context *gin.Context) {
	model := mogo.NewDoc(model.File{}).(*model.File)
	filename := context.Param("filename")
	err := model.FindOne(bson.M{"filename": filename}, model)
	if err != nil {
		return
	}
	context.JSON(http.StatusOK, model)
}
