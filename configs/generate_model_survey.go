package configs

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"

	survey "github.com/AlecAivazis/survey/v2"
)

type AndiModel struct {
	Module  string
	Name    string
	Package string
}

func GenerateModelSurvey() (answers AndiModel, err error) {
	pack := []string{"new package"}
	files, err := ioutil.ReadDir(AppDir + "pkg/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		name := file.Name()
		if file.IsDir() && name != "middleware" {
			pack = append(pack, file.Name())
		}
	}

	// the questions to ask
	qs := []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Input{
				Message: "What is the name of your model?",
				Default: "",
			},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
		{
			Name: "package",
			Prompt: &survey.Select{
				Message: "Choose the model destination package:",
				Options: pack,
			},
			Validate: survey.Required,
		},
	}
	// the answers will be written to this struct
	answers = AndiModel{}
	answers.Module = GetAppModule()
	// perform the questions
	err = survey.Ask(qs, &answers)
	return
}

func GetAppModule() string {
	file, err := os.Open(AppDir + "go.mod")
	if err != nil {
		panic("Generation can only be done at the root level of the application")
	}
	r := bufio.NewReader(file)
	line, _, err := r.ReadLine()
	if err != nil {
		panic(err)
	}
	moduleName := bytes.TrimPrefix(line, []byte("module "))
	return string(moduleName)
}
