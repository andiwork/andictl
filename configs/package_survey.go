package configs

import (
	survey "github.com/AlecAivazis/survey/v2"
)

type AndiPackage struct {
	ModuleName string
	Port       string
}

func PackageSurvey() (answers AndiPackage, err error) {
	// the questions to ask
	qs := []*survey.Question{
		{
			Name: "port",
			Prompt: &survey.Input{
				Message: "Enter your appplication port",
				Default: "8080",
			},
			Validate: survey.Required,
		},
	}
	// the answers will be written to this struct
	answers = AndiPackage{}
	// perform the questions
	err = survey.Ask(qs, &answers)
	return
}
