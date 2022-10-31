/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "amqC",
	Short: "Module for generating AMQ authz",
	Long:  `Longer description needed`,
	Run:   func(cmd *cobra.Command, args []string) { someshit() },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.amqC.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func someshit() {

	yfile, err := ioutil.ReadFile("/Users/leo/repo/rl/AMQConfig/sample.yml")

	if err != nil {
		log.Fatal(err)
	}

	data := Config{}

	err2 := yaml.Unmarshal(yfile, &data)

	if err2 != nil {

		log.Fatal(err2)
	}

	result := data.generateAllData()

	result.generateAuthorizationEntries()
	/*result.generateUserProperties()
	result.generateGroupProperties()
	result.queueEntries()
	result.topicEntries()*/
}
