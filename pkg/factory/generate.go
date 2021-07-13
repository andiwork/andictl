package factory

import (
	"fmt"
	"os"

	"github.com/andiwork/andictl/configs"
	"github.com/andiwork/andictl/utils"
)

func Generate(factory configs.AndiFactory) {

	factorySlug := utils.AndictlSlugify(factory.Name)
	if factory.Package == "new package" {
		factory.Package = factorySlug
	}

	packPath := configs.AppDir + "pkg/" + factory.Package
	fmt.Println("create ", packPath)
	os.MkdirAll(packPath, os.ModePerm)

	// Generate model files
	data, _ := factoryTmpl.ReadFile("templates/factory.go.gotmpl")
	utils.ProcessTmplFiles(packPath, factorySlug+"_factory.go", data, factory, false)

	data, _ = defaultTmpl.ReadFile("templates/default.go.gotmpl")
	utils.ProcessTmplFiles(packPath, factorySlug+"_default.go", data, factory, false)

	data, _ = helloTmpl.ReadFile("templates/hello.go.gotmpl")
	utils.ProcessTmplFiles(packPath, "dummy_hello.go", data, factory, false)

	data, _ = worldTmpl.ReadFile("templates/world.go.gotmpl")
	utils.ProcessTmplFiles(packPath, "dummy_world.go", data, factory, false)

	data, _ = readmeTmpl.ReadFile("templates/README.md.gotmpl")
	utils.ProcessTmplFiles(packPath, "README.md", data, factory, false)

}
