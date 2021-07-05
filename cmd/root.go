/*
Copyright © 2021 James Kokou Gagglo <freemanpolys@gmail.com>

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
	"github.com/andiwork/andictl/pkg/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "andictl",
	Short: "Andictl is a  tool for managing your Andi Web Application.",
	Long:  `Andictl is a  tool for managing your Andi Web Application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err != nil {
			// perform the questions
			os.Create(".andictl.yaml")
			answers, err := configs.InitSurvey()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("responseeee", answers)
				viper.Set("application.name", answers.Name)
				viper.Set("application.type", answers.Type)
				viper.Set("application.port", answers.Port)
				viper.Set("application.database-type", answers.DatabaseType)
				viper.Set("application.auth", answers.Auth)
				err := viper.WriteConfig()
				if err != nil {
					fmt.Println("Eror while writing config file:", err)
				}
			}
		}
		if err := viper.Unmarshal(&configs.AppConfs); err != nil {
			fmt.Printf("couldn't read config: %s", err)
		}
		//Generate app skeleton
		app.Generate()

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.andictl.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Find home directory.
	//home, err := homedir.Dir()
	//cobra.CheckErr(err)

	// Search config in home directory with name ".andictl" (without extension).
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName(".andictl")
	viper.AutomaticEnv() // read in environment variables that match
}
