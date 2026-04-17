package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var backend string
var all bool
var CheckCommand = &cobra.Command{
	Use:     "check",
	Aliases: []string{"ch"},
	Short:   "check the status of the specific backend server --backend <backend server name or url >  or check the status of all the backend servers --all  ",
	Long: `check the status of the specific backend server --backend <backend server name or url>  
	check the status of all the backend servers --all and also check the status of the load balancer and the overall status of the load balancer `,
	RunE: func(cmd *cobra.Command, args []string) error {
		Check_url := Admin_Url + "/check"
		if backend == "" && !all {
			return errors.New("please Provide backend name or use --all flag ")
		}
		type payload struct {
			Name string `json:"name"`
			ALL  bool   `json:"all"`
		}
		data := &payload{
			Name: backend,
			ALL:  all,
		}
		bodybuffer := new(bytes.Buffer)
		json.NewEncoder(bodybuffer).Encode(data)
		client, err := http.Post(Check_url, "application/json", bodybuffer)
		if err != nil {
			return err
		}
		defer client.Body.Close()
		states := make(map[string]bool)
		if err := json.NewDecoder(client.Body).Decode(&states); err != nil {
			return err
		}
		if all {
			for backend, alive := range states {
				if alive {
					fmt.Println(color.GreenString("backend %s is ALIVE  ", backend))
				} else {
					fmt.Println(color.RedString("backend %s is DEAD  ", backend))
				}
			}
		} else {
			if alive, ok := states[backend]; ok {
				if alive {
					fmt.Println(color.GreenString("backend %s is ALIVE  ", backend))
				} else {
					fmt.Println(color.RedString("backend %s is DEAD  ", backend))
				}
			}
		}
		return nil

	},
}

func init() {

	rootCmd.AddCommand(CheckCommand)
	CheckCommand.Flags().StringVarP(&backend, "backend", "b", "", "the name or ip address of the backend server to check its status")
	CheckCommand.Flags().BoolVarP(&all, "all", "a", false, "check the status of all the backend servers and the load balancer")

}
