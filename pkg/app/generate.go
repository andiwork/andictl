package app

import (
	"log"
	"os"
	"strings"

	"html/template"

	execute "github.com/alexellis/go-execute/pkg/v1"
	"github.com/andiwork/andictl/configs"
	"github.com/andiwork/andictl/utils"
	"github.com/google/uuid"
)

var appDir = "test/"

func Generate() {

	log.Println("les configs ======", configs.AppConfs)
	// create app folder structure
	// folder configs
	os.MkdirAll(appDir+"configs", os.ModePerm)
	// folder docs/swagger-ui
	os.MkdirAll(appDir+"docs/swagger-ui", os.ModePerm)
	// folder pkg
	os.MkdirAll(appDir+"pkg", os.ModePerm)
	// initialiaze go module
	os.Chdir(appDir)
	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		ExecShellCommand("go", []string{"mod", "init", configs.AppConfs.App.Name})
	}
	defer ExecShellCommand("go", []string{"mod", "tidy"})
	//Generate main.go
	data, _ := mainGoTmpl.ReadFile("templates/main.go.gotmpl")
	processTmplFiles(".", "main.go", data, false)
	// Generate configs
	// package files
	genFolder := "configs"

	data, _ = appTmpl.ReadFile("templates/app.yaml.gotmpl")
	processTmplFiles(genFolder, "app.yaml", data, false)

	data, _ = prodTmpl.ReadFile("templates/prod.yaml.gotmpl")
	processTmplFiles(genFolder, "prod.yaml", data, false)

	data, _ = appGoTmpl.ReadFile("templates/app.go.gotmpl")
	//fmt.Printf("Total bytes: %s\n", data)
	processTmplFiles(genFolder, "app.go", data, false)

	data, _ = gormGoTmpl.ReadFile("templates/gorm.go.gotmpl")
	processTmplFiles(genFolder, "gorm.go", data, false)

	data, _ = restfulGoTmpl.ReadFile("templates/restful.go.gotmpl")
	processTmplFiles(genFolder, "restful.go", data, false)

	data, _ = swaggerGoTmpl.ReadFile("templates/swagger.go.gotmpl")
	processTmplFiles(genFolder, "swagger.go", data, false)

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

func processTmplFiles(folder string, dstFileName string, tmplData []byte, debug bool) {
	tmpl := template.New("app-conf").Funcs(template.FuncMap{
		"uuidWithOutHyphen": uuidWithOutHyphen,
	})
	tmpl, err := tmpl.Parse(string(tmplData))
	if err != nil {
		log.Fatal("Error Parsing template: ", err)
		return
	}
	filePath := folder + "/" + dstFileName
	if debug {
		err = tmpl.Execute(os.Stderr, configs.AppConfs.App)
	} else {
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatal("Error creating file. ", err)
			return
		}

		err = tmpl.Execute(file, configs.AppConfs.App)

	}
	if err != nil {
		log.Fatal("Error executing template. ", filePath, err)
	}

}

func uuidWithOutHyphen() (value string) {
	uuidWithHyphen := uuid.New()
	value = strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	return
}
