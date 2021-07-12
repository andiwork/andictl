package model

import (
	"fmt"
	"os"
	"reflect"

	"github.com/andiwork/andictl/configs"
	"github.com/andiwork/andictl/utils"
	"github.com/spf13/viper"
)

type TemplateData struct {
	First map[interface{}]interface{}
	Data  []map[interface{}]interface{}
}

var (
	appModule string
)

func Generate(model configs.AndiModel) {

	// create package folder
	//path, _ := os.Getwd()
	//pack := path[strings.LastIndex(path, "/")+1:]
	modelSlug := utils.AndictlSlugify(model.Name)
	appModule = model.Module
	//fetch model in config file
	exist, models := IsKeyInConfFile("models", "name", modelSlug)

	if len(exist) == 0 {
		// set application AuthType
		model.AuthType = viper.GetString("application.authtype")
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

		data, _ = modelServiceGoTmpl.ReadFile("templates/model_service.go.gotmpl")
		utils.ProcessTmplFiles(packPath, modelSlug+"_service.go", data, model, false)

		// register  models in package init.go
		registerModels := make([]map[interface{}]interface{}, 1)
		allModels := make([]map[interface{}]interface{}, 1)
		newModel := make(map[interface{}]interface{}, 1)
		newModel["name"] = model.Name
		newModel["package"] = model.Package
		// collect existing models
		registerModels, allModels = IsKeyInConfFile("models", "package", model.Package)
		registerModels = append(registerModels, newModel)
		allModels = append(allModels, newModel)
		//==> add new model to register
		templateData := TemplateData{First: registerModels[0], Data: registerModels}
		data, _ = initGoTmpl.ReadFile("templates/init.go.gotmpl")
		utils.ProcessTmplFiles(packPath, "init.go", data, templateData, false)

		//==> import new package in gorm.go
		confDir := configs.AppDir + "configs"
		packages := GetDistinctElementInConf("models", "package")
		//add current package to the existing
		packages[model.Package] = appModule
		data, _ = gormMigrateTmpl.ReadFile("templates/gorm.go.gotmpl")
		utils.ProcessTmplFiles(confDir, "gorm.go", data, packages, false)

		//=> import new package and create service in restful.go
		templateData = TemplateData{First: packages, Data: allModels}
		data, _ = restfulWebserviceGoTmpl.ReadFile("templates/restful.go.gotmpl")
		utils.ProcessTmplFiles(confDir, "restful.go", data, templateData, false)

		fmt.Println("======= TODO ======")
		fmt.Println("Execute: go mod tidy")
		fmt.Println("===================")
		//Update andictl.yaml with new model
		updateAndictlConfFile(modelSlug, model.Package, models)
	} else {
		conf := viper.ConfigFileUsed()
		fmt.Println("model", modelSlug, "already exist. Take a look at the conf file:", conf)
	}

}

func updateAndictlConfFile(modelName string, modelPackage string, models []map[interface{}]interface{}) {
	model := make(map[interface{}]interface{}, 1)
	model["name"] = modelName
	model["package"] = modelPackage
	models = append(models, model)
	viper.Set("models", models)
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println("Eror while writing config file:", err)
	}
}

//IsKeyInConfFile search an element identified by searchKey=searchValue in config file
func IsKeyInConfFile(getKey string, searchKey string, searchValue string) (exist []map[interface{}]interface{}, entries []map[interface{}]interface{}) {
	fromFile := viper.Get(getKey)
	if fromFile != nil && reflect.TypeOf(fromFile) != reflect.TypeOf("string") {
		elements := fromFile.([]interface{})
		//fmt.Println("get model 0 ", models[0].(map[interface{}]interface{})["package"])
		for _, v := range elements {
			value := v.(map[interface{}]interface{})
			entries = append(entries, value)
			if value[searchKey] == searchValue {
				exist = append(exist, value)
			}
		}
	}
	//os.Exit(0)

	return
}

//GetElementInConf get an element identified by searchKey in config file
func GetDistinctElementInConf(getKey string, searchKey string) (exist map[interface{}]interface{}) {
	fromFile := viper.Get(getKey)
	exist = make(map[interface{}]interface{}, 1)
	if fromFile != nil && reflect.TypeOf(fromFile) != reflect.TypeOf("string") {
		entries := fromFile.([]interface{})
		//fmt.Println("get model 0 ", models[0].(map[interface{}]interface{})["package"])
		for _, v := range entries {
			value := v.(map[interface{}]interface{})
			//find the key in map from conf file
			if element, ok := value[searchKey]; ok {
				//add the value found into a map to avoid duplication
				exist[element.(string)] = appModule
			}
		}
	}

	return
}
