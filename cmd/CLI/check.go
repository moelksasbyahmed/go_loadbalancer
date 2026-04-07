package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var backend string
var all bool
var CheckCommand = &cobra.Command{
	Use:     "check",
	Aliases: []string{"ch"},
	Short:   "check the status of the specific backend server --backend <backend server name or ip address>  or check the status of all the backend servers --all  ",
	Long: `check the status of the specific backend server --backend <backend server name or ip address>  
	check the status of all the backend servers --all and also check the status of the load balancer and the overall status of the load balancer `,
	RunE: func(cmd *cobra.Command, args []string) error {

		fmt.Println("checking backend server  ... ", backend)

		return nil

	},
}

func init() {

	rootCmd.AddCommand(CheckCommand)
	CheckCommand.Flags().StringVarP(&backend, "backend", "b", "", "the name or ip address of the backend server to check its status")
	CheckCommand.Flags().BoolVarP(&all, "all", "a", false, "check the status of all the backend servers and the load balancer")

}
