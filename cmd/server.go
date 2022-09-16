/*
Copyright © 2022 Safal Neupane <safal.npane@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/safalnpane/gapi/cli"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Manage servers in your project",
	Long: `GAPI can handle multiple servers in a project. You can use this command to
manage the API server on your project.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
		if cli.IsProject() {
			project := cli.Project{}
			project.ReadProject()
			fmt.Println(project)
		} else {
            fmt.Println("[-] Project not found.")
        }
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
