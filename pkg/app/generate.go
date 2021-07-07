package app

import (
	"log"
	"os"

	execute "github.com/alexellis/go-execute/pkg/v1"
	"github.com/andiwork/andictl/configs"
	"github.com/andiwork/andictl/utils"
)

func Generate() {

	log.Println("les configs ======", configs.AppConfs)
	// create app folder structure
	// folder configs
	os.MkdirAll(configs.AppDir+"configs", os.ModePerm)
	// folder docs/swagger-ui
	os.MkdirAll(configs.AppDir+"docs/swagger-ui", os.ModePerm)
	// folder pkg/middleware
	os.MkdirAll(configs.AppDir+"pkg/middleware", os.ModePerm)
	// initialiaze go module
	os.Chdir(configs.AppDir)
	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		ExecShellCommand("go", []string{"mod", "init", configs.AppConfs.App.Name})
	}
	defer ExecShellCommand("go", []string{"mod", "tidy"})
	//Generate main.go
	data, _ := mainGoTmpl.ReadFile("templates/main.go.gotmpl")
	utils.ProcessTmplFiles(".", "main.go", data, false)
	// Generate configs
	// package files
	genFolder := "configs"

	data, _ = appTmpl.ReadFile("templates/app.yaml.gotmpl")
	utils.ProcessTmplFiles(genFolder, "app.yaml", data, false)

	data, _ = prodTmpl.ReadFile("templates/prod.yaml.gotmpl")
	utils.ProcessTmplFiles(genFolder, "prod.yaml", data, false)

	data, _ = appGoTmpl.ReadFile("templates/app.go.gotmpl")
	//fmt.Printf("Total bytes: %s\n", data)
	utils.ProcessTmplFiles(genFolder, "app.go", data, false)

	data, _ = gormGoTmpl.ReadFile("templates/gorm.go.gotmpl")
	utils.ProcessTmplFiles(genFolder, "gorm.go", data, false)

	data, _ = restfulGoTmpl.ReadFile("templates/restful.go.gotmpl")
	utils.ProcessTmplFiles(genFolder, "restful.go", data, false)

	data, _ = swaggerGoTmpl.ReadFile("templates/swagger.go.gotmpl")
	utils.ProcessTmplFiles(genFolder, "swagger.go", data, false)

	// Download swagger ui files
	if _, err := os.Stat("docs/swagger-ui/dist"); os.IsNotExist(err) {

		swagger, err := utils.DownloadFile("v3.51.1.tar.gz", "https://github.com/swagger-api/swagger-ui/archive/refs/tags/v3.51.1.tar.gz")
		if err == nil {
			log.Println("swagger: ", swagger)
			//path, _ := os.Getwd()
			//untar
			ExecShellCommand("tar -xzf "+swagger+" -C /tmp", []string{})
			ExecShellCommand("mv /tmp/swagger-ui-3.51.1/dist docs/swagger-ui", []string{})
		} else {
			panic(err)
		}
	}
}

func ExecShellCommand(bin string, args []string) {
	log.Println("Executing : ", bin, args)
	cmd := execute.ExecTask{
		Command:     bin,
		Args:        args,
		StreamStdio: false,
	}

	res, err := cmd.Execute()
	if err != nil {
		panic(err)
	}

	if res.ExitCode != 0 {
		panic("Non-zero exit code: " + res.Stderr)
	}
	//fmt.Printf("stdout: %s, stderr: %s, exit-code: %d\n", res.Stdout, res.Stderr, res.ExitCode)

}
