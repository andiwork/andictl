package model

import (
	"fmt"
	"os"
	"strings"

	"github.com/andiwork/andictl/configs"
	"github.com/andiwork/andictl/utils"
	"github.com/spf13/viper"
)

func Generate(model configs.AndiModel) {

	// create package folder
	//path, _ := os.Getwd()
	//pack := path[strings.LastIndex(path, "/")+1:]
	modelSlug := utils.AndictlSlugify(model.Name)

	//fetch model in config file
	exist, models := IsKeyInConfFile("models", "name", modelSlug)

	if !exist {
		if model.Package == "new package" {
			model.Package = modelSlug
		}
		packPath := configs.AppDir + "pkg/" + model.Package
		fmt.Println("create ", packPath)
		os.MkdirAll(packPath, os.ModePerm)

		// Generate model files
		data, _ := modelGoTmpl.ReadFile("templates/model.go.gotmpl")
		utils.ProcessTmplFiles(packPath, modelSlug+"_model.go", data, model, false)

		data, _ = modelResourceGoTmpl.ReadFile("templates/model_resource.go.gotmpl")
		utils.ProcessTmplFiles(packPath, modelSlug+"_resource.go", data, model, false)

		fmt.Println("======= TODO ======")

		initGo := packPath + "/init.go"
		if _, err := os.Stat(initGo); os.IsNotExist(err) {
			data, _ = initGoTmpl.ReadFile("templates/init.go.gotmpl")
			utils.ProcessTmplFiles(packPath, "init.go", data, model, false)
			//fmt.Println("create ", packPath+"/init.go")
		} else {
			// register new models
			register := fmt.Sprintf("models = append(models, new(%s))", strings.Title(model.Name))
			fmt.Println("Add: ", register, " before //andi-add-model-to-migrate in file: ", initGo)
			fmt.Println("==")
			//utils.InsertInfile(register, "//andi-generate-do-not-remove", initGo)
		}

		// import new package in gorm.go
		gormFile := configs.AppDir + "configs/gorm.go"
		importPackage := fmt.Sprintf("\"%s/pkg/%s\"", model.Module, model.Package)
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
		fmt.Println("==")
		fmt.Println("Execute: go mod tidy")
		fmt.Println("===================")
		//Update andictl.yaml with new model
		updateAndictlConfFile(modelSlug, model.Package, models)
	} else {
		fmt.Println("model", modelSlug, "already exist")
	}

}

func updateAndictlConfFile(modelName string, modelPackage string, models []interface{}) {
	model := make(map[string]string, 1)
	model["name"] = modelName
	model["package"] = modelPackage
	models = append(models, model)
	viper.Set("models", models)
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println("Eror while writing config file:", err)
	}
}

func IsKeyInConfFile(getKey string, searchKey string, searchValue string) (exist bool, entries []interface{}) {
	if err := viper.ReadInConfig(); err == nil {
		fromFile := viper.Get(getKey)
		if fromFile != nil {
			entries = fromFile.([]interface{})
			//fmt.Println("get model 0 ", models[0].(map[interface{}]interface{})["package"])
			for _, v := range entries {
				if v.(map[interface{}]interface{})[searchKey] == searchValue {
					exist = true
					break
				}
			}
		}

	}
	return
}
