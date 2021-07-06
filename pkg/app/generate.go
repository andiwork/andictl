package app

import (
	"log"
	"os"
	"strings"

	"html/template"

	execute "github.com/alexellis/go-execute/pkg/v1"
	"github.com/andiwork/andictl/configs"
	"github.com/google/uuid"
)

var test = "test/"

func Generate() {

	log.Println("les configs ======", configs.AppConfs)
	// create app folder structure
	// folder configs
	os.MkdirAll(test+"configs", os.ModePerm)
	// folder docs/swagger-ui
	os.MkdirAll(test+"docs/swagger-ui", os.ModePerm)
	// folder pkg
	os.MkdirAll(test+"pkg", os.ModePerm)
	// initialiaze go module
	os.Chdir(test)
	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		ExecShellCommand("go", []string{"mod", "init", configs.AppConfs.App.Name})
	}
	ExecShellCommand("go", []string{"mod", "tidy"})

	// Generate configs files
	data, _ := appTmpl.ReadFile("templates/app.yaml.tmpl")
	processTmplFiles("configs", "app.yaml", data)

	data, _ = prodTmpl.ReadFile("templates/prod.yaml.tmpl")
	processTmplFiles("configs", "prod.yaml", data)

}

func ExecShellCommand(bin string, args []string) {
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

func processTmplFiles(folder string, dstFileName string, tmplData []byte) {
	tmpl := template.New("app-conf").Funcs(template.FuncMap{
		"uuidWithOutHyphen": uuidWithOutHyphen,
	})
	tmpl, err := tmpl.Parse(string(tmplData))
	if err != nil {
		log.Fatal("Error Parsing template: ", err)
		return
	}
	filePath := folder + "/" + dstFileName
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Error creating file. ", err)
		return
	}

	err = tmpl.Execute(file, configs.AppConfs.App)

	if err != nil {
		log.Fatal("Error executing template. ", filePath, err)
	}

}

func uuidWithOutHyphen() (value string) {
	uuidWithHyphen := uuid.New()
	value = strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	return
}
