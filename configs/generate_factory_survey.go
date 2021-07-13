package configs

import (
	survey "github.com/AlecAivazis/survey/v2"
	"github.com/andiwork/andictl/utils"
)

type AndiFactory struct {
	Name    string
	Package string
	Module  string
}

func GenerateFactorySurvey() (answers AndiFactory, err error) {
	pack := utils.PackageList(AppDir + "pkg/")
	// the questions to ask
	qs := []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Input{
				Message: "What is the name of your factory?",
				Default: "",
			},
			Validate: survey.Required,
		},
		{
			Name: "package",
			Prompt: &survey.Select{
				Message: "Choose the factory destination package:",
				Options: pack,
			},
			Validate: survey.Required,
		},
	}
	// the answers will be written to this struct
	answers = AndiFactory{}
	answers.Module = utils.GetAppModule(AppDir + "go.mod")
	// perform the questions
	err = survey.Ask(qs, &answers)
	return
}
