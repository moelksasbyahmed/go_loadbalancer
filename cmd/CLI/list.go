package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"text/tabwriter"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var ListCommand = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "list all the backend servers and their status",
	Long:    "list all the backend servers and their status",
	RunE: func(cmd *cobra.Command, args []string) error {
		List_url := Admin_Url + "/list"
		fmt.Println("Generating the list ... ")
		type response struct {
			Name  string  `json:"name"`
			Alive bool    `json:"alive"`
			Url   url.URL `json:"url"`
		}

		var res []response
		resp, err := http.Get(List_url)
		if err != nil {
			return fmt.Errorf("failed to send list request: %v", err)
		}
		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return fmt.Errorf("failed to decode list response: %v", err)
		}
		fmt.Println("\n---------- Backend Servers ----------")

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "NAME\tURL\tSTATUS")
		for _, server := range res {
			status := color.RedString("DEAD")
			if server.Alive {
				status = color.GreenString("ALIVE")
			}
			fmt.Fprintf(w, "%s\t%s\t%s\n", server.Name, server.Url.String(), status)
		}
		w.Flush()
		fmt.Println("-------------------------------------")
		return nil
	},
}

func init() {

	rootCmd.AddCommand(ListCommand)

}
