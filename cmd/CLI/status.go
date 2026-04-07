package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var StatusCmd = &cobra.Command{
	Use:     "status",
	Aliases: []string{"st"},
	Short:   "shows the current status of the load balancer and the backend servers",
	Long: `it gives an overview of the all the backend server  
	 their status also how many request each backend server has handled 
	 the current load on each backend server and the overall status of the load balancer `,
	RunE: func(cmd *cobra.Command, args []string) error {

		fmt.Println("Analyzing ... ")
		return nil

	},
}

func init() {

	rootCmd.AddCommand(StatusCmd)

}
