/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/andiwork/andictl/configs"
	"github.com/andiwork/andictl/pkg/model"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// modelCmd represents the model command
var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "generate model in package pkg/<model_name>/models ",
	Long: `generate model in package pkg/<model_name>/models. 
	Example:

andictl generate model --name hello --fields "name:string,age:int"
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		answers, err := configs.GenerateModelSurvey()
		if err != nil {
			fmt.Printf("Error ", err)
		}
		//	answers.Name = slugify.Marshal(answers.Name)
		model.Generate(answers)
	},
}

func init() {
	generateCmd.AddCommand(modelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// modelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// modelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
