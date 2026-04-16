package main

import (
	"fmt"
	"net"

	"github.com/moelksasbyahmed/go_loadbalancer/internal"
)

var Admin_Url string

func main() {
	Adminconfig, _ := internal.LoadConfig("")
	if Adminconfig == nil {
		fmt.Println("Failed to load configuration, using default admin URL")
		return
	}
	Admin_Url = "http://" + net.JoinHostPort(Adminconfig.Adminconfig.Host, Adminconfig.Adminconfig.Port)
	Execute()

}
