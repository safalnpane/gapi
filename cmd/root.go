/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gapi",
	Short: "GAPI is a go-based API platform for keyboard lovers",
	Long: `GAPI aims to provide a keyboard-friendly API paltform.
You can use gapi to develop small-scale rest-API project. GAPI can work with
multiple servers and changing between them is very easy.

GAPI relies on a directory "gapi_project" which contains all necessary files
related to the project. It is recommended to add this folder to the version
tracking system.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


