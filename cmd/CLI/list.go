package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ListCommand = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "list all the backend servers and their status",
	Long:    "list all the backend servers and their status",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Generating the list ... ")
		return nil
	},
}

func init() {

	rootCmd.AddCommand(ListCommand)

}
