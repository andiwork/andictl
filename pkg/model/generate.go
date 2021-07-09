package model

import (
	"log"
	"os"

	"github.com/andiwork/andictl/configs"
	"github.com/andiwork/andictl/utils"
)

func Generate(model configs.AndiModel) {

	// create package folder
	//path, _ := os.Getwd()
	//pack := path[strings.LastIndex(path, "/")+1:]
	modelSlug := utils.AndictlSlugify(model.Name)
	if model.Package == "new package" {
		model.Package = modelSlug
	}
	log.Println("model: ", model)
	log.Println("Creating model: ", modelSlug, "in package :", model.Package)

	packPath := configs.AppDir + "pkg/" + model.Package
	os.MkdirAll(packPath, os.ModePerm)

	// Generate model files
	data, _ := modelGoTmpl.ReadFile("templates/model.go.gotmpl")
	utils.ProcessTmplFiles(packPath, "model.go", data, model, false)

	data, _ = modelResourceGoTmpl.ReadFile("templates/model_resource.go.gotmpl")
	utils.ProcessTmplFiles(packPath, modelSlug+"_resource.go", data, model, false)

	if _, err := os.Stat("init.go"); os.IsNotExist(err) {
		data, _ = initGoTmpl.ReadFile("templates/init.go.gotmpl")
		utils.ProcessTmplFiles(packPath, "init.go", data, model, false)
	}

}
