package main

import (
	"errors"
	"fmt"
	"net"

	admin "github.com/moelksasbyahmed/go_loadbalancer/cmd/AdminApi"

	config "github.com/moelksasbyahmed/go_loadbalancer/internal"
	server "github.com/moelksasbyahmed/go_loadbalancer/internal/server"
	"github.com/spf13/cobra"
)

var LBserver *server.Server
var port, configpath string
var actualport string
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "start the load balancer and listen for incoming traffic on a specified port or read from config.yaml file",
	Long:  "start the load balancer and listen for incoming traffic on a specified port or read from config.yaml file and distribute the traffic to the backend servers using multibel algorithims ",
	RunE: func(cmd *cobra.Command, args []string) error {

		fmt.Println("Reading the confguration from Config.yaml file ... ", configpath)
		config, err := config.LoadConfig(configpath)
		if err != nil {
			return err
		}
		fmt.Println("the configuration is read successfully ")
		Algo, err := server.GetAlgorithim(config.LoadBalancerConfig.Algorithim)

		if err != nil {
			return err
		}

		Loadbalancer := server.NewloadBalancer(&server.LoadBalancerConfig{
			Algorithim: Algo,
		})
		config.LoadBalancerConfig.Port, err = handle_port(config)
		if err != nil {
			return err
		}
		Loadbalancer.Populate_LoadBalancer(config)
		LBserver = server.NewServer(config, Loadbalancer)
		go LBserver.Start()
		adminapi := admin.NewAdminAPI(LBserver, config)
		adminErr := adminapi.Start()
		if adminErr != nil {
			return adminErr
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
	StartCmd.Flags().StringVarP(&port, "port", "p", "", "the port the loadbalancer will work on ")
	StartCmd.Flags().StringVarP(&configpath, "config", "c", "", "the path of the config.yaml file to read the configuration from it the Default path is ./config.yaml ")
	rootCmd.AddCommand(StartCmd)

}

func handle_port(config *config.Config) (string, error) {

	if port != "" {
		actualport = port

	} else {
		actualport = config.LoadBalancerConfig.Port

	}
	if !testport(actualport) {
		return actualport, errors.New("the port is not available you can kill the connection or try another port the retry the connection you entered  " + actualport)
	}
	return actualport, nil
}
