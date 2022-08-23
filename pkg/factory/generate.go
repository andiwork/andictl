package factory

import (
	"embed"
	"fmt"
	"os"

	"github.com/andiwork/andictl/configs"
	"github.com/andiwork/andictl/utils"
)

//go:embed templates/*
var content embed.FS

func Generate(factory configs.AndiFactory) {

	factorySlug := utils.AndictlSlugify(factory.Name)
	if factory.Package == "new package" {
		factory.Package = factorySlug
	}

	packPath := configs.AppDir + "pkg/" + factory.Package
	fmt.Println("create ", packPath)
	os.MkdirAll(packPath, os.ModePerm)

	// Generate model files
	data, _ := content.ReadFile("templates/factory.go.gotmpl")
	utils.ProcessTmplFiles(packPath, factorySlug+"_factory.go", data, factory, false)

	data, _ = content.ReadFile("templates/hello.go.gotmpl")
	utils.ProcessTmplFiles(packPath, "dummy_hello.go", data, factory, false)

	data, _ = content.ReadFile("templates/world.go.gotmpl")
	utils.ProcessTmplFiles(packPath, "dummy_world.go", data, factory, false)

	data, _ = content.ReadFile("templates/README.md.gotmpl")
	utils.ProcessTmplFiles(packPath, "README.md", data, factory, false)

}
