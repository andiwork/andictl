package app

import (
	"embed"
	"fmt"
	"os"
	"sync"

	"github.com/andiwork/andictl/configs"
	"github.com/andiwork/andictl/utils"
)

var wg sync.WaitGroup

//go:embed templates/*
var content embed.FS

func Generate() {

	// create app folder structure
	// folder configs
	os.MkdirAll(configs.AppDir+"configs", os.ModePerm)
	fmt.Println("create ", "configs")

	// folder docs/swagger-ui
	os.MkdirAll(configs.AppDir+"docs/swagger-ui", os.ModePerm)
	fmt.Println("create ", "docs/swagger-ui")
	// Download swagger ui files
	if _, err := os.Stat(configs.AppDir + "docs/swagger-ui/dist"); os.IsNotExist(err) {
		wg.Add(1)
		go func() {
			swagger, err := utils.DownloadFile("v3.51.1.tar.gz", "https://github.com/swagger-api/swagger-ui/archive/refs/tags/v3.51.1.tar.gz")
			if err == nil {
				//untar
				utils.ExecShellCommand("tar -xzf "+swagger+" -C /tmp", []string{}, false)
				utils.ExecShellCommand(fmt.Sprintf("mv /tmp/swagger-ui-3.51.1/dist %s/docs/swagger-ui", configs.AppDir), []string{}, false)
			} else {
				fmt.Printf("Error ", err)
				os.Exit(0)
			}
			wg.Done()
		}()

	}

	// folder pkg/middleware
	os.MkdirAll(configs.AppDir+"pkg/middleware", os.ModePerm)
	fmt.Println("create ", "pkg/middleware")

	//defer ExecShellCommand("go", []string{"mod", "tidy"})
	//Generate main.go
	wg.Add(1)
	go func() {
		data, _ := content.ReadFile("templates/main.go.gotmpl")
		utils.ProcessTmplFiles(configs.AppDir, "main.go", data, configs.AppConfs, false)

		data, _ = content.ReadFile("templates/gitignore.gotmpl")
		utils.ProcessTmplFiles(configs.AppDir, ".gitignore", data, configs.AppConfs, false)
		wg.Done()
	}()

	// Generate configs
	// package files
	confDir := configs.AppDir + "configs"
	go func() {
		wg.Add(1)
		data, _ := content.ReadFile("templates/app.yaml.gotmpl")
		utils.ProcessTmplFiles(confDir, "app.yaml", data, configs.AppConfs, false)

		data, _ = content.ReadFile("templates/prod.yaml.gotmpl")
		utils.ProcessTmplFiles(confDir, "prod.yaml", data, configs.AppConfs, false)
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		data, _ := content.ReadFile("templates/app.go.gotmpl")
		utils.ProcessTmplFiles(confDir, "app.go", data, configs.AppConfs, false)

		data, _ = content.ReadFile("templates/gorm.go.gotmpl")
		utils.ProcessTmplFiles(confDir, "gorm.go", data, configs.AppConfs, false)

		data, _ = content.ReadFile("templates/custom_gorm.go.gotmpl")
		utils.ProcessTmplFiles(confDir, "custom_gorm.go", data, configs.AppConfs, false)

		data, _ = content.ReadFile("templates/custom_restful.go.gotmpl")
		utils.ProcessTmplFiles(confDir, "custom_restful.go", data, configs.AppConfs, false)

		wg.Done()
	}()
	go func() {
		wg.Add(1)
		data, _ := content.ReadFile("templates/restful.go.gotmpl")
		utils.ProcessTmplFiles(confDir, "restful.go", data, configs.AppConfs, false)

		data, _ = content.ReadFile("templates/swagger.go.gotmpl")
		utils.ProcessTmplFiles(confDir, "swagger.go", data, configs.AppConfs, false)
		data, _ = content.ReadFile("templates/authz.go.gotmpl")
		utils.ProcessTmplFiles(configs.AppDir+"pkg/middleware", "authz.go", data, nil, false)
		wg.Done()
	}()

	os.MkdirAll(configs.AppDir+"utils", os.ModePerm)
	data, _ := content.ReadFile("templates/db_singleton.go.gotmpl")
	utils.ProcessTmplFiles(configs.AppDir+"utils", "db_singleton.go", data, nil, false)

	if configs.AppConfs.App.AuthType == "jwt" {
		go func() {
			wg.Add(1)
			data, _ := content.ReadFile("templates/jwt.go.gotmpl")
			utils.ProcessTmplFiles(configs.AppDir+"pkg/middleware", "jwt.go", data, nil, false)
			data, _ = content.ReadFile("templates/swagger_helper.go.gotmpl")
			utils.ProcessTmplFiles(configs.AppDir+"utils", "swagger_helper.go", data, nil, false)
			wg.Done()
		}()

	}
	wg.Wait()
	// initialiaze go module
	os.Chdir(configs.AppDir)
	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		utils.ExecShellCommand("go", []string{"mod", "init", configs.AppConfs.App.Name}, false)
	}
	utils.ExecShellCommand("go", []string{"mod", "tidy"}, false)
}
