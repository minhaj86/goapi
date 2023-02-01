package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goonode/mogo"
	appconfig "goapi/config"
	"goapi/controller"
	"goapi/model"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	//	var data = `
	//db: localhost
	//`
	filename, _ := filepath.Abs("./config/config.yml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}
	//log.Printf("%s", yamlFile)

	config := appconfig.Database{}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Value: %v \n", config.DB.ConnectionString)

	dbconfig := &mogo.Config{
		ConnectionString: config.DB.ConnectionString,
		Database:         config.DB.Name,
	}
	_, err = mogo.Connect(dbconfig)
	if err != nil {
		log.Fatal(err)
	}
	mogo.ModelRegistry.Register(model.File{})
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/some/:filename", controller.Getting)
	router.POST("/some", controller.Posting)
	//router.PUT("/somePut", putting)
	//router.DELETE("/someDelete", deleting)
	//router.PATCH("/somePatch", patching)
	//router.HEAD("/someHead", head)
	//router.OPTIONS("/someOptions", options)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}
