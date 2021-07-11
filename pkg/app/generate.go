package app

import (
	"fmt"
	"os"

	execute "github.com/alexellis/go-execute/pkg/v1"
	"github.com/andiwork/andictl/configs"
	"github.com/andiwork/andictl/utils"
)

func Generate() {

	// create app folder structure
	// folder configs
	os.MkdirAll(configs.AppDir+"configs", os.ModePerm)
	fmt.Println("create ", "configs")

	// folder docs/swagger-ui
	os.MkdirAll(configs.AppDir+"docs/swagger-ui", os.ModePerm)
	fmt.Println("create ", "docs/swagger-ui")

	// folder pkg/middleware
	os.MkdirAll(configs.AppDir+"pkg/middleware", os.ModePerm)
	fmt.Println("create ", "pkg/middleware")

	//defer ExecShellCommand("go", []string{"mod", "tidy"})
	//Generate main.go
	data, _ := mainGoTmpl.ReadFile("templates/main.go.gotmpl")
	utils.ProcessTmplFiles(configs.AppDir, "main.go", data, configs.AppConfs, false)

	// Generate configs
	// package files
	confDir := configs.AppDir + "configs"

	data, _ = appTmpl.ReadFile("templates/app.yaml.gotmpl")
	utils.ProcessTmplFiles(confDir, "app.yaml", data, configs.AppConfs, false)

	data, _ = prodTmpl.ReadFile("templates/prod.yaml.gotmpl")
	utils.ProcessTmplFiles(confDir, "prod.yaml", data, configs.AppConfs, false)

	data, _ = appGoTmpl.ReadFile("templates/app.go.gotmpl")
	utils.ProcessTmplFiles(confDir, "app.go", data, configs.AppConfs, false)

	data, _ = gormGoTmpl.ReadFile("templates/gorm.go.gotmpl")
	utils.ProcessTmplFiles(confDir, "gorm.go", data, configs.AppConfs, false)

	data, _ = restfulGoTmpl.ReadFile("templates/restful.go.gotmpl")
	utils.ProcessTmplFiles(confDir, "restful.go", data, configs.AppConfs, false)

	data, _ = swaggerGoTmpl.ReadFile("templates/swagger.go.gotmpl")
	utils.ProcessTmplFiles(confDir, "swagger.go", data, configs.AppConfs, false)

	data, _ = authzGoTmpl.ReadFile("templates/authz.go.gotmpl")
	utils.ProcessTmplFiles(configs.AppDir+"pkg/middleware", "authz.go", data, nil, false)

	if configs.AppConfs.App.AuthType == "jwt" {
		data, _ = jwtGoTmpl.ReadFile("templates/jwt.go.gotmpl")
		utils.ProcessTmplFiles(configs.AppDir+"pkg/middleware", "jwt.go", data, nil, false)
	}

	// Download swagger ui files
	if _, err := os.Stat(configs.AppDir + "docs/swagger-ui/dist"); os.IsNotExist(err) {

		swagger, err := utils.DownloadFile("v3.51.1.tar.gz", "https://github.com/swagger-api/swagger-ui/archive/refs/tags/v3.51.1.tar.gz")
		if err == nil {
			//untar
			ExecShellCommand("tar -xzf "+swagger+" -C /tmp", []string{})
			ExecShellCommand(fmt.Sprintf("mv /tmp/swagger-ui-3.51.1/dist %s/docs/swagger-ui", configs.AppDir), []string{})
		} else {
			fmt.Printf("Error ", err)
			os.Exit(0)
		}
	}
	// initialiaze go module
	os.Chdir(configs.AppDir)
	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		ExecShellCommand("go", []string{"mod", "init", configs.AppConfs.App.Name})
	}
	fmt.Println("======= TODO ======")
	fmt.Println("Execute: go mod tidy")
	fmt.Println("===================")
}

func ExecShellCommand(bin string, args []string) {
	cmd := execute.ExecTask{
		Command:     bin,
		Args:        args,
		StreamStdio: false,
	}

	res, err := cmd.Execute()
	if err != nil {
		fmt.Printf("Error ", err)
		os.Exit(0)
	}

	if res.ExitCode != 0 {
		fmt.Printf("Non-zero exit code: ", res.Stderr)
		os.Exit(0)
	}
	//fmt.Printf("stdout: %s, stderr: %s, exit-code: %d\n", res.Stdout, res.Stderr, res.ExitCode)

}
