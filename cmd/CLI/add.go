package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var name, backendIP string

var AddCommand = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a", "add_backend"},
	Short:   "add a backend server to the load balancer and update the database with the new backend server information use --name <backend server name> --backend-ip <backend server ip address> ",
	Long: `add a backend server to the load balancer and update the database with the new backend server information 
	     also update the status of the load balancer and the backend servers in the database `,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Adding the backend server ... ")
		return nil

	},
}

func init() {
	AddCommand.Flags().StringVarP(&name, "name", "n", "default", "adding name of the backend server for clarity when status or check command occur ")
	AddCommand.Flags().StringVarP(&backendIP, "backend-ip", "i", "", "adding ip address of the backend server for clarity when status or check command occur ")
	rootCmd.AddCommand(AddCommand)

}
