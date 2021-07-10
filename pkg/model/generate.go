package model

import (
	"fmt"
	"log"
	"os"

	"github.com/andiwork/andictl/configs"
	"github.com/andiwork/andictl/utils"
	"github.com/spf13/viper"
)

type TemplateData struct {
	First map[interface{}]interface{}
	Data  []map[interface{}]interface{}
}

var (
	templateData TemplateData
	appModule    string
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

		// register  models in package init.go
		initGo := packPath + "/init.go"
		registerModels := make([]map[interface{}]interface{}, 0)
		newModel := make(map[interface{}]interface{}, 1)
		newModel["name"] = model.Name
		newModel["package"] = model.Package
		if _, err := os.Stat(initGo); !os.IsNotExist(err) {
			// collect existing models
			registerModels, _ = IsKeyInConfFile("models", "package", model.Package)
		}
		//==> add new model to register
		registerModels = append(registerModels, newModel)
		templateData.Data = registerModels
		templateData.First = registerModels[0]
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
		data, _ = restfulWebserviceGoTmpl.ReadFile("templates/restful.go.gotmpl")
		utils.ProcessTmplFiles(confDir, "restful.go", data, packages, false)

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

//IsKeyInConfFile search an element identified by searchKey=searchValue in config file
func IsKeyInConfFile(getKey string, searchKey string, searchValue string) (exist []map[interface{}]interface{}, entries []interface{}) {
	fromFile := viper.Get(getKey)
	if fromFile != nil {
		entries = fromFile.([]interface{})
		//fmt.Println("get model 0 ", models[0].(map[interface{}]interface{})["package"])
		for _, v := range entries {
			value := v.(map[interface{}]interface{})
			if value[searchKey] == searchValue {
				exist = append(exist, value)
			}
		}
	}

	return
}

//GetElementInConf get an element identified by searchKey in config file
func GetDistinctElementInConf(getKey string, searchKey string) (exist map[string]string) {
	fromFile := viper.Get(getKey)
	exist = make(map[string]string)
	if fromFile != nil {
		entries := fromFile.([]interface{})
		//fmt.Println("get model 0 ", models[0].(map[interface{}]interface{})["package"])
		for _, v := range entries {
			value := v.(map[interface{}]interface{})
			log.Println(" ==== ", value)
			//find the key in map from conf file
			if element, ok := value[searchKey]; ok {
				//add the value found into a map to avoid duplication
				exist[element.(string)] = appModule
			}
		}
	}

	return
}
