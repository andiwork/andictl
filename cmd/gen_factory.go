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
	"os"

	"github.com/andiwork/andictl/configs"
	"github.com/andiwork/andictl/pkg/factory"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// factoryCmd represents the factory command
var factoryCmd = &cobra.Command{
	Use:   "factory",
	Short: "generate factory in package pkg/<factory_name>",
	Long: `generate factory in package pkg/<factory_name>
	Example:
	andictl generate factory
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("Error ", err)
			os.Exit(0)
		}

		answers, err := configs.GenerateFactorySurvey()
		if err != nil {
			fmt.Printf("Error ", err)
		}
		factory.Generate(answers)
	},
}

func init() {
	generateCmd.AddCommand(factoryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// factoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// factoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
