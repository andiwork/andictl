package configs

import (
	survey "github.com/AlecAivazis/survey/v2"
	"github.com/andiwork/andictl/utils"
)

type AndiModel struct {
	Module     string
	Name       string
	Package    string
	ApiVersion string
	AuthType   string
}

func GenerateModelSurvey() (answers AndiModel, err error) {
	pack := utils.PackageList(AppDir + "pkg/")
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
		{
			Name: "apiVersion",
			Prompt: &survey.Input{
				Message: "What is apiVersion?",
				Default: "v1",
			},
		},
	}
	// the answers will be written to this struct
	answers = AndiModel{}
	answers.Module = utils.GetAppModule(AppDir + "go.mod")
	// perform the questions
	err = survey.Ask(qs, &answers)
	return
}
