/*
Copyright © 2022 Safal Neupane <safal.npane@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run individual request file againts a server.",
	Long: `Run individual request file againts a default server set on your project.
You can specify the server to use by specifying --with-server flag followed by the name of
the server. This command uses the API-client built inside of gapi to make a web request to
the specified server and try to parse the response in JSON format.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
