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
	"errors"
	"os"

	"github.com/andiwork/andictl/utils"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Live reload for Go apps",
	Long:  `Live reload for Go apps based on Air`,
	Run: func(cmd *cobra.Command, args []string) {
		home, err := homedir.Dir()
		andictlHome := home + "/.andictl"
		if err == nil {
			if _, err := os.Stat(andictlHome + "/air"); errors.Is(err, os.ErrNotExist) {
				os.MkdirAll(andictlHome, os.ModePerm)
				air, _ := utils.DownloadFile("air-install.sh", "https://raw.githubusercontent.com/cosmtrek/air/master/install.sh")
				utils.ExecShellCommand("sh "+air+" -b "+andictlHome, []string{}, false)
			}

		}
		utils.ExecShellCommand(andictlHome+"/air", nil, true)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
