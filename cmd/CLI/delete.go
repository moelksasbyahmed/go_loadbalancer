package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DeleteCommand = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d", "del"},
	Short:   "delete a backend server from the load balancer and remove the backend server from the database with name --name <backend server name> or ip address --backend-ip <backend server ip address>  ",
	Long: `delete a backend server from the load balancer and remove the backend server from the database and also update
	 the status of the load balancer and the backend servers in the database `,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Deleting the backend server ... ")
		return nil

	},
}

func init() {

	DeleteCommand.Flags().StringVarP(&name, "name", "n", "", "deleting name of the backend server for clarity when status or check command occur ")
	DeleteCommand.Flags().StringVarP(&backendIP, "backend-ip", "i", "", "deleting ip address of the backend server for clarity when status or check command occur ")
	rootCmd.AddCommand(DeleteCommand)

}
