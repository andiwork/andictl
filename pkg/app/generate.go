package app

import (
	"log"
	"os"

	execute "github.com/alexellis/go-execute/pkg/v1"
	"github.com/andiwork/andictl/configs"
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

}

func ExecShellCommand(bin string, args []string) {
	path, _ := os.Getwd()
	log.Println(" current folder ", path)
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
