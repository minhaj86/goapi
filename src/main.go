package main

import (
	"github.com/gin-gonic/gin"
	appconfig "github.com/goapi/config"
	"github.com/goapi/controller"
	"github.com/goapi/model"
	"github.com/goonode/mogo"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	filename, _ := filepath.Abs("./config/config.yml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	config := appconfig.Database{}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	dbconfig := &mogo.Config{
		ConnectionString: config.DB.ConnectionString,
		Database:         config.DB.Name,
	}
	_, err = mogo.Connect(dbconfig)
	if err != nil {
		log.Fatal(err)
	}
	mogo.ModelRegistry.Register(model.File{})
	router := gin.Default()
	router.GET("/some/:filename", controller.Get)
	router.POST("/some", controller.Posting)
	router.Run()
}
