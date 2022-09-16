/*
Copyright © 2022 Safal Neupane <safal.npane@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
    "github.com/safalnpane/gapi/cli"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise a new gapi project",
	Long: `Initialise a new gapi project.
This creates a folder named "gapi_project" on the current working
directory. This directory contains project realted configurations`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
        cli.ProjectInit("test", "This is a test project")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
