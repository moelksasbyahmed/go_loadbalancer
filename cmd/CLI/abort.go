package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var AbortCommand = &cobra.Command{
	Use:     "abort",
	Aliases: []string{"ab"},
	Short:   "abort the load balancer and stop all the backend servers and also update the status of the load balancer and the backend servers in the database ",
	Long: `Abort the load balancer and gracefully stop all connected backend servers.

This command will also:
  * Sync and update the status of the load balancer in the Postgres database.
  * Update the final status of all backend servers in the database.
  * Generate a final traffic and health report.
  
By default, the report is printed to the console. You can use flags to save it to a file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Aborting the load balancer ... ")
		return nil
	},
}

func init() {

	rootCmd.AddCommand(AbortCommand)

}
