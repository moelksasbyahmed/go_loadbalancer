package main

import (
	"errors"
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

var port, configpath string

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "start the load balancer and listen for incoming traffic on a specified port or read from config.yaml file",
	Long:  "start the load balancer and listen for incoming traffic on a specified port or read from config.yaml file and distribute the traffic to the backend servers using multibel algorithims ",
	RunE: func(cmd *cobra.Command, args []string) error {
		if port != "" {
			fmt.Printf("starting the load balancer on port %s \n", port)
			if !testport(port) {
				return errors.New("the port is not available you can kill the connection or try another port the retry the connection you entered  " + port)
			}

		}
		if configpath != "" && configpath != "./config.yaml" {
			fmt.Println("Reading the confguration from Config.yaml file ... ")
			fmt.Println("starting the Loadbalancer ")

		}

		return nil

	},
}

func testport(port string) bool {

	A, err := net.Listen("tcp", ":"+port)

	if err != nil {
		return false
	}
	defer A.Close()
	return true

}

func init() {
	StartCmd.Flags().StringVarP(&port, "port", "p", "8080", "the port the loadbalancer will work on ")
	StartCmd.Flags().StringVarP(&configpath, "config", "c", "./config.yaml", "the path of the config.yaml file to read the configuration from it ")
	rootCmd.AddCommand(StartCmd)

}
