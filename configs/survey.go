package configs

import (
	survey "github.com/AlecAivazis/survey/v2"
)

func InitSurvey() (answers Application, err error) {
	// the questions to ask
	qs := []*survey.Question{
		{
			Name: "type",
			Prompt: &survey.Select{
				Message: "Which type of application would you like to create",
				Options: []string{"api"},
				Default: "api",
			},
			Validate: survey.Required,
		},
		{
			Name: "name",
			Prompt: &survey.Input{
				Message: "What is the name of your application?",
				Default: "andictl",
			},
			Validate: survey.Required,
		},
		{
			Name: "port",
			Prompt: &survey.Input{
				Message: "On which port would like your server to run?",
				Default: "9090",
			},
			Validate: survey.Required,
		},
		{
			Name: "auth",
			Prompt: &survey.Select{
				Message: "Which type of authentication would you like to use?",
				Options: []string{"none", "jwt", "oidc"},
				Default: "none",
			},
			Validate: survey.Required,
		},
		{
			Name: "databaseType",
			Prompt: &survey.Select{
				Message: "Which type of database would you like to use?",
				Options: []string{"non", "postgres", "mysql"},
				Default: "postgres",
			},
			Validate: survey.Required,
		},
	}
	// the answers will be written to this struct
	answers = Application{}
	// perform the questions
	err = survey.Ask(qs, &answers)
	return
}
