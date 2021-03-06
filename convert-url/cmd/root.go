/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"convert-url/lib"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	homedir "github.com/mitchellh/go-homedir"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Short: "show all commands",
	Long: `
 ____                          _   
 / ___|___  _ ____   _____ _ __| |_ 
| |   / _ \| '_ \ \ / / _ \ '__| __|
| |__| (_) | | | \ V /  __/ |  | |_ 
 \____\___/|_| |_|\_/ \___|_|   \__|
                                    by andiahmad
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		listCommand(cmd)
	},
}

func listCommand(cmd *cobra.Command) {
	listFlag := []string{"j", "i", ""}
stop:
	for _, flag := range listFlag {
		switch flag {
		case "j":
			toJson, _ := cmd.Flags().GetString(flag)
			if toJson != "" {
				lib.ConvertToJosn(toJson)
				break stop

			}
		case "i":
			str, _ := cmd.Flags().GetString(flag)

			convert := fmt.Sprintf("`%s`", str)
			fmt.Println(convert)
			if str != "" {
				res, err := lib.PrettyString(convert)
				if err != nil {
					log.Fatal(err.Error())
				}
				fmt.Println(res)

			}
			break stop

		default:
			fmt.Println(cmd.Long)
			fmt.Println("Use: -h ", cmd.Short)

		}

	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().String("j", "", `convert to json (ex: cv --j "a=1")`)
	rootCmd.PersistentFlags().String("i", "", `next feature`)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".convert-url" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".convert-url")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
