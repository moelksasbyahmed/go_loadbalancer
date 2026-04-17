package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fatih/color"
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

		fmt.Println(color.GreenString("Analyzing ... "))
		type response struct {
			Name           string `json:"name"`
			Alive          bool   `json:"alive"`
			Url            string `json:"url"`
			CurrentTraffic int    `json:"current_traffic"`
			OverallTraffic int    `json:"overall_traffic"`
		}
		var res []response
		resp, err := http.Get(Admin_Url + "/status")
		if err != nil {
			return fmt.Errorf("failed to send status request: %v", err)
		}
		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return fmt.Errorf("failed to decode status response: %v", err)
		}
		fmt.Println("\n---------- Backend Servers Status ----------")
		for _, server := range res {
			status := color.RedString("DEAD")
			if server.Alive {
				status = color.GreenString("ALIVE")
			}
			fmt.Printf("Name: %s | Status: %s | URL: %s | Current Traffic: %d | Overall Traffic: %d\n", server.Name, status, server.Url, server.CurrentTraffic, server.OverallTraffic)
		}

		return nil

	},
}

func init() {

	rootCmd.AddCommand(StatusCmd)

}
