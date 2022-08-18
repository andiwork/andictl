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
	"github.com/google/uuid"
	slugify "github.com/metal3d/go-slugify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//var cfgFile string
var andictlVersion bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "andictl",
	Short: "andictl is a  tool for managing your Andi Web Application.",
	Long:  `andictl is a  tool for managing your Andi Web Application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if andictlVersion {
			fmt.Println("version: 1.0.0")
		} else {

			// If a config file is found, read it in.
			if err := viper.ReadInConfig(); err != nil {
				// perform the questions
				_, err := os.Create(configs.AppDir + ".andi.yaml")
				if err != nil {
					fmt.Println(err)
					os.Exit(0)
				}
				answers, err := configs.InitSurvey()
				if err != nil {
					fmt.Println(err)
					os.Exit(0)
				} else {
					appName := slugify.Marshal(answers.Name)
					viper.Set("application.name", appName)
					viper.Set("application.type", answers.Type)
					viper.Set("application.port", answers.Port)
					viper.Set("application.app-id", uuid.New().String())
					viper.Set("application.database-type", answers.DatabaseType)
					if answers.AuthType != "none" {
						viper.Set("application.auth", true)
						viper.Set("application.authType", answers.AuthType)
					}
					viper.Set("models", "")

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
		}

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
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolVarP(&andictlVersion, "version", "v", false, "show andictl version")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Find home directory.
	//home, err := homedir.Dir()
	//cobra.CheckErr(err)

	// Search config in home directory with name ".andictl" (without extension).
	viper.AddConfigPath(configs.AppDir)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".andi")
	viper.AutomaticEnv() // read in environment variables that match
}

func IsKeyInConfFile(getKey string, searchKey string, searchValue string) (exist bool, entries []interface{}) {
	if err := viper.ReadInConfig(); err == nil {
		fromFile := viper.Get(getKey)
		if fromFile != nil {
			entries = fromFile.([]interface{})
			//fmt.Println("get model 0 ", models[0].(map[interface{}]interface{})["package"])
			for _, v := range entries {
				if v.(map[interface{}]interface{})[searchKey] == searchValue {
					exist = true
					break
				}
			}
		}

	}
	return
}
