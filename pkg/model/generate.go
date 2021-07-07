package model

import (
	"log"
	"os"
	"strings"

	"github.com/andiwork/andictl/configs"
	"github.com/metal3d/go-slugify"
)

func Generate(model configs.AndiModel) {

	// create package folder
	//path, _ := os.Getwd()
	//pack := path[strings.LastIndex(path, "/")+1:]
	modelSlug := AndictlSlugify(model.Name)
	log.Println("Creating model: ", modelSlug, "in package :", model.Package)
	os.MkdirAll(configs.AppDir+"pkg/"+modelSlug, os.ModePerm)
}

func AndictlSlugify(text string) string {
	return strings.ToLower(strings.ReplaceAll(slugify.Marshal(text), "-", ""))
}
