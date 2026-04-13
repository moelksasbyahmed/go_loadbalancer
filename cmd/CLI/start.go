package main

import (
	"errors"
	"fmt"
	"net"

	server "github.com/moelksasbyahmed/go_loadbalancer/internal"
	"github.com/spf13/cobra"
)

var port, configpath string
var actualport string
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "start the load balancer and listen for incoming traffic on a specified port or read from config.yaml file",
	Long:  "start the load balancer and listen for incoming traffic on a specified port or read from config.yaml file and distribute the traffic to the backend servers using multibel algorithims ",
	RunE: func(cmd *cobra.Command, args []string) error {

		fmt.Println("Reading the confguration from Config.yaml file ... ", configpath)
		config, err := server.LoadConfig(configpath)
		if err != nil {
			return err
		}
		actualport, err = handle_port(config)
		if err != nil {
			return err
		}

		fmt.Println("starting the Loadbalancer  on port ", actualport)

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
	StartCmd.Flags().StringVarP(&port, "port", "p", "", "the port the loadbalancer will work on ")
	StartCmd.Flags().StringVarP(&configpath, "config", "c", "", "the path of the config.yaml file to read the configuration from it the Default path is ./config.yaml ")
	rootCmd.AddCommand(StartCmd)

}

func handle_port(config *server.Config) (string, error) {

	if port != "" {
		actualport = port

	} else {
		actualport = config.ProxyConfig.Proxy_port

	}
	if !testport(actualport) {
		return actualport, errors.New("the port is not available you can kill the connection or try another port the retry the connection you entered  " + actualport)
	}
	return actualport, nil
}
