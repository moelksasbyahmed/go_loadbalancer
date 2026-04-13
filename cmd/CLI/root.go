package main

import (
	"fmt"
	"os"
	_ "os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var version = "0.0.1"
var rootCmd = &cobra.Command{
	Use:   "load-balancer",
	Short: "simple demonstration of load balancer in GO",
	Long:  " this is load  balancer in Go that demonstrates how to distribute incoming traffic across multiple backend servers to improve performance and reliability using multibel algorithims .",
	Run: func(cmd *cobra.Command, args []string) {

		cyan := color.New(color.FgCyan).SprintFunc()
		green := color.New(color.FgGreen, color.Bold).SprintFunc()

		asciiArt := `
    __    ____           ____  ____  ________ __
   / /   / __ )         / __ \/ __ \/ ____/ //_/
  / /   / __  |  ______/ /_/ / / / / /   / ,<   
 / /___/ /_/ /  /_____/ _, _/ /_/ / /___/ /| |  
/_____/_____/        /_/ |_|\____/\____/_/ |_|  
`
		fmt.Println(cyan(asciiArt))

		fmt.Println(green("Welcome to my  Load balancer  (LB-ROCK)  please use --help for more information about available command "))

	},
	Version:      version,
	SilenceUsage: true,
}

/*
* we will have 8  commands : start , abort , add , delete , check , list , status  , sync   each command will have its  own functionality and flags
we will store the Backends in postgres database so we will need a migration tool

|-
cmd -- : root.go , main.go , start.go , abort.go , add.go , delete.go , check.go , list.go , status.go , report.go , sync.go
the functionality of each command will defined in the internal package and we will use the cobra library to handle the command line interface
the functionality of each commmand will be as follows

	start : will start the load balancer and listen for incoming traffic on a specified port or read from config.yaml file
	abort : will stop the load balancer and close all connections and sync the state of the load balancer with the database  and store the current state of the load balancer in the database
	add : will add a new backend server to the load balancer and store the backend server in the database
	delete : will delete a backend server from the load balancer and remove the backend server from the database
	check : will check the health of the backend servers and update the status of the backend servers in the database
	list : will list all the backend servers and their status
	status : will show the current status of the load balancer and the backend servers

	sync : will sync the state of the load balancer with the database and update the status of the backend servers in the database
*/
func Execute() {
	if err := rootCmd.Execute(); err != nil {

		fmt.Fprintf(os.Stderr, "there was an error lunching the CLI %v ", err)
		os.Exit(1)
	}

}
