package main

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
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
		Abort_url := Admin_Url + "/abort"

		resp, err := http.Post(Abort_url, "", nil)
		if err != nil {
			return fmt.Errorf("failed to send abort request: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("abort request failed with status: %s", resp.Status)
		}
		defer resp.Body.Close()
		color.Green("Load balancer aborted successfully. Final status and report have been updated in the database.")
		return nil
	},
}

func init() {

	rootCmd.AddCommand(AbortCommand)

}
