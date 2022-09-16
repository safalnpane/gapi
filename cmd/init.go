/*
Copyright © 2022 Safal Neupane <safal.npane@gmail.com>

*/
package cmd

import (
    "errors"
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
	Run: initRun,
    Aliases: []string{"i"},
    Args: func(cmd *cobra.Command, args []string) error {
        if len(args) < 2 {
            return errors.New("Init needs 'Name' and 'Description' of the project")
        }
        return nil
    },
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initRun(cmd *cobra.Command, args []string) {
    cli.ProjectInit(args[0], args[1])
}
