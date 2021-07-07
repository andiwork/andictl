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
	log.Println("Creating model: ", modelSlug, "in package :", model.Package)
	os.MkdirAll(configs.AppDir+"pkg/"+modelSlug, os.ModePerm)
}
