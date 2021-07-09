package model

import (
	"fmt"
	"os"
	"strings"

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
	packPath := configs.AppDir + "pkg/" + model.Package
	fmt.Println("create ", packPath)
	os.MkdirAll(packPath, os.ModePerm)

	// Generate model files
	data, _ := modelGoTmpl.ReadFile("templates/model.go.gotmpl")
	utils.ProcessTmplFiles(packPath, "model.go", data, model, false)
	fmt.Println("create ", packPath+"/model.go")

	data, _ = modelResourceGoTmpl.ReadFile("templates/model_resource.go.gotmpl")
	utils.ProcessTmplFiles(packPath, modelSlug+"_resource.go", data, model, false)
	fmt.Println("create ", packPath+"/"+modelSlug+"_resource.go")

	initGo := packPath + "/init.go"
	fmt.Println("======= TODO ======")
	if _, err := os.Stat(initGo); os.IsNotExist(err) {
		data, _ = initGoTmpl.ReadFile("templates/init.go.gotmpl")
		utils.ProcessTmplFiles(packPath, "init.go", data, model, false)
		fmt.Println("create ", packPath+"/init.go")
	} else {
		// register new models
		register := fmt.Sprintf("models = append(models, new(%s))", strings.Title(model.Name))
		fmt.Println("Add: ", register, " before //andi-add-model-to-migrate in file: ", initGo)
		fmt.Println("==")
		//utils.InsertInfile(register, "//andi-generate-do-not-remove", initGo)
	}

	// import new package in gorm.go
	gormFile := configs.AppDir + "configs/gorm.go"
	importPackage := fmt.Sprintf("%s/pkg/%s", model.Module, model.Package)
	fmt.Println("Import: ", importPackage, " before //andi-import-do-not-remove in file: ", gormFile)

	// register migration new model in gorm.go
	migrateModel := fmt.Sprintf("GormDb.AutoMigrate(%s.Migrate()...)", model.Package)
	fmt.Println("Add: ", migrateModel, " before //andi-auto-migrate-do-not-remove in file: ", gormFile)
	fmt.Println("==")
	// import new package and create service in restful.go
	restfulFile := configs.AppDir + "configs/restful.go"
	fmt.Println("Import: ", importPackage, " before //andi-import-do-not-remove in file: ", restfulFile)

	createService := fmt.Sprintf("restful.DefaultContainer.Add(%s.New(GormDb).WebService())", strings.ToLower(model.Name))
	fmt.Println("Add: ", createService, " before //andi-add-restful-webservice in file: ", restfulFile)
	fmt.Println("===================")
}
