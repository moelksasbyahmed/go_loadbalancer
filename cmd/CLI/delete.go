package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var DeleteCommand = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d", "del"},
	Short:   "delete a backend server from the load balancer and remove the backend server from the database with name --name <backend server name> or ip address --backend-ip <backend server ip address>  ",
	Long: `delete a backend server from the load balancer and remove the backend server from the database and also update
	 the status of the load balancer and the backend servers in the database `,
	RunE: func(cmd *cobra.Command, args []string) error {
		if name == "" && backend_url == "" {
			return fmt.Errorf("you must provide either the name of the backend server or the url of the backend server to delete it from the load balancer and the database ")
		}
		Del_url := Admin_Url + "/remove"
		payload := &Backend{
			Name: name,
			Url:  backend_url,
		}
		body_bytes := new(bytes.Buffer)
		err := json.NewEncoder(body_bytes).Encode(&payload)
		if err != nil {
			return fmt.Errorf("failed to encode payload: %v", err)
		}
		resp, err := http.NewRequest(http.MethodDelete, Del_url, body_bytes)
		if err != nil {
			return fmt.Errorf("failed to create request: %v", err)
		}
		client := &http.Client{}
		response, err := client.Do(resp)
		if err != nil {
			return err
		}
		defer response.Body.Close()
		if err != nil {
			return fmt.Errorf("failed to send request: %v", err)
		}
		if response.StatusCode != http.StatusOK {
			return fmt.Errorf("server error (%d): %s couldnt delete the backend server ", response.StatusCode, response.Status)
		}
		fmt.Println(color.GreenString("successfully removed backend %s ", name))
		return nil

	},
}

func init() {

	DeleteCommand.Flags().StringVarP(&name, "name", "n", "", "deleting name of the backend server for clarity when status or check command occur ")
	DeleteCommand.Flags().StringVarP(&backend_url, "backend-url", "u", "", "deleting url of the backend server for clarity when status or check command occur ")
	rootCmd.AddCommand(DeleteCommand)

}
