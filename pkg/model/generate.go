package model

import (
	"embed"
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
	//go:embed templates/*
	content embed.FS
)

func Generate(model configs.AndiModel) {

	// create package folder
	//path, _ := os.Getwd()
	//pack := path[strings.LastIndex(path, "/")+1:]
	modelSlug := utils.AndictlSlugify(model.Name)
	appModule = model.Module
	//fetch model in config file
	exist, models := IsKeyInConfFile("models", "name", model.Name)
	//fmt.Println(modelSlug, "IsKeyInConfFile exist:", len(exist))
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
		data, _ := content.ReadFile("templates/model.go.gotmpl")
		utils.ProcessTmplFiles(packPath, modelSlug+"_model.go", data, model, false)

		data, _ = content.ReadFile("templates/model_resource.go.gotmpl")
		utils.ProcessTmplFiles(packPath, modelSlug+"_resource.go", data, model, false)

		data, _ = content.ReadFile("templates/model_service.go.gotmpl")
		utils.ProcessTmplFiles(packPath, modelSlug+"_service.go", data, model, false)

		data, _ = content.ReadFile("templates/model_repository.go.gotmpl")
		utils.ProcessTmplFiles(packPath, modelSlug+"_repository.go", data, model, false)

		data, _ = content.ReadFile("templates/model_cache_service.go.gotmpl")
		utils.ProcessTmplFiles(packPath, modelSlug+"_cache_service.go", data, model, false)

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
		data, _ = content.ReadFile("templates/init.go.gotmpl")
		utils.ProcessTmplFiles(packPath, "init.go", data, templateData, false)

		data, _ = content.ReadFile("templates/type.go.gotmpl")
		utils.ProcessTmplFiles(packPath, "type.go", data, templateData, false)

		//==> import new package in gorm.go
		confDir := configs.AppDir + "configs"
		packages := GetDistinctElementInConf("models", "package")
		//add current package to the existing
		packages[model.Package] = appModule
		data, _ = content.ReadFile("templates/gorm.go.gotmpl")
		utils.ProcessTmplFiles(confDir, "gorm.go", data, packages, false)

		//=> import new package and create service in restful.go
		templateData = TemplateData{First: packages, Data: allModels}
		data, _ = content.ReadFile("templates/restful.go.gotmpl")
		utils.ProcessTmplFiles(confDir, "restful.go", data, templateData, false)

		// create model test case
		testsPath := configs.AppDir + "pkg/tests"
		os.MkdirAll(testsPath, os.ModePerm)
		data, _ = content.ReadFile("templates/model_service_test.go.gotmpl")
		utils.ProcessTmplFiles(testsPath, modelSlug+"_service_test.go", data, model, false)

		data, _ = content.ReadFile("templates/mock_model_repository.go.gotmpl")
		utils.ProcessTmplFiles(testsPath, "mock_"+modelSlug+"_repository.go", data, model, false)

		data, _ = content.ReadFile("templates/README.md.gotmpl")
		utils.ProcessTmplFiles(testsPath, "README.md", data, model, false)

		//Update andi.yaml with new model
		updateAndictlConfFile(model.Name, model.Package, models)

		defer utils.ExecShellCommand("go", []string{"mod", "tidy"}, false)

	} else {
		conf := viper.ConfigFileUsed()
		fmt.Println("model", model.Name, "already exist. Take a look at the conf file:", conf)
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
	//fmt.Println(" === seach key ===", getKey, searchKey, searchValue)
	fromFile := viper.Get(getKey)
	if fromFile != nil && reflect.TypeOf(fromFile) != reflect.TypeOf("string") {
		elements := fromFile.([]interface{})
		//fmt.Println("get elements 0 ", elements[0].(map[interface{}]interface{})["package"])
		for _, v := range elements {
			value := v.(map[interface{}]interface{})
			entries = append(entries, value)
			//fmt.Println("append ", value, searchValue)
			if value[searchKey] == searchValue {
				exist = append(exist, value)
				//fmt.Println("find:", searchValue, "in", viper.ConfigFileUsed())
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
