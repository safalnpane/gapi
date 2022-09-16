/*
Copyright © 2022 Safal Neupane <safal.npane@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var requestCmd = &cobra.Command{
	Use:   "request",
	Short: "API request information",
	Long: `This handles API request related information.
The request in gapi is the information describing an REST-API request.
This includes request path, method, headers and request body. Use this
command to list, describe, update and delete the API requests on your project.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("request called")
	},
}

func init() {
	rootCmd.AddCommand(requestCmd)
}
