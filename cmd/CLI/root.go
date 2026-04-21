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

func Execute() {
	if err := rootCmd.Execute(); err != nil {

		fmt.Fprintf(os.Stderr, "there was an error lunching the CLI %v ", err)
		os.Exit(1)
	}

}
