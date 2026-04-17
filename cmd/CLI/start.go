package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/fatih/color"
	admin "github.com/moelksasbyahmed/go_loadbalancer/cmd/AdminApi"

	config "github.com/moelksasbyahmed/go_loadbalancer/internal"
	server "github.com/moelksasbyahmed/go_loadbalancer/internal/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var LBserver *server.Server
var port, configpath string
var actualport string
var wg sync.WaitGroup
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "start the load balancer and listen for incoming traffic on a specified port or read from config.yaml file",
	Long:  "start the load balancer and listen for incoming traffic on a specified port or read from config.yaml file and distribute the traffic to the backend servers using multibel algorithims ",
	RunE: func(cmd *cobra.Command, args []string) error {
		wg.Add(2)
		fmt.Println("Reading the confguration from Config.yaml file ... ", configpath)
		config, writerconfig, err := config.LoadConfig(configpath)
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
		}, writerconfig)
		config.LoadBalancerConfig.Port, err = handle_port(config)
		if err != nil {
			return err
		}
		if !testport(config.Adminconfig.Port) {
			return errors.New("the admin port is not available you can kill the connection or try another port the retry the connection you entered  " + config.Adminconfig.Port)
		}
		Loadbalancer.Populate_LoadBalancer(config)
		HealthCtx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctx, sigCancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
		defer sigCancel()
		Loadbalancer.StartHealthCheckLoop(HealthCtx, ctx, config.LoadBalancerConfig.HealthCheckInterval)
		LBserver = server.NewServer(config, Loadbalancer)
		go func() {
			defer wg.Done()
			LBserver.Start()

		}()

		adminapi := admin.NewAdminAPI(LBserver, config)
		go func() {
			defer wg.Done()
			adminapi.Start()

		}()
		go func() {
			<-ctx.Done()
			log.Println(color.RedString("\n[System] Ctrl+C received. Initiating shutdown..."))
			shutdown, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer shutdownCancel()
			adminapi.Shutdown(shutdown)

		}()
		wg.Wait()
		vipererr := WriteViper(*writerconfig)
		if vipererr != nil {
			return vipererr
		}
		fmt.Println(color.GreenString("Both servers have safely finished and returned. Exiting program."))

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

func WriteViper(writer viper.Viper) error {
	if err := writer.WriteConfig(); err != nil {
		return err
	}
	fmt.Println(color.GreenString("Successfully synced ServerPool to server.yaml"))
	return nil
}
