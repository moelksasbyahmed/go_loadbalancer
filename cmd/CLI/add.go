package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var name, backend_url string
var Max_request int

var AddCommand = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a", "add_backend"},
	Short:   "add a backend server to the load balancer and update the database with the new backend server information use --name <backend server name> --backend-ip <backend server ip address> ",
	Long: `add a backend server to the load balancer and update the database with the new backend server information 
	     also update the status of the load balancer and the backend servers in the database `,
	RunE: func(cmd *cobra.Command, args []string) error {
		var Add_Url = Admin_Url + "/add"

		data := &Backend{
			Name:        name,
			Url:         backend_url,
			Max_request: Max_request,
		}
		bodybuffer := new(bytes.Buffer)
		if err := json.NewEncoder(bodybuffer).Encode(data); err != nil {
			return err
		}
		resp, err := http.Post(Add_Url, "application/json", bodybuffer)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		respBody, _ := io.ReadAll(resp.Body)

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("server error (%d): %s", resp.StatusCode, string(respBody))
		}

		fmt.Println("Success:", string(respBody))

		color.Cyan("success added backend %s with UrL %s to LoadBalancer", name, backend_url)
		return nil

	},
}

func init() {
	AddCommand.Flags().StringVarP(&name, "name", "n", "default", "adding name of the backend server for clarity when status or check command occur ")
	AddCommand.Flags().StringVarP(&backend_url, "backend-url", "u", "", "adding url of the backend server for clarity when status or check command occur ")
	AddCommand.Flags().IntVarP(&Max_request, "max-request", "m", 100, "the maximum number of requests that the backend server can handle before it is marked as unhealthy and removed from the load balancer ")
	rootCmd.AddCommand(AddCommand)

}
