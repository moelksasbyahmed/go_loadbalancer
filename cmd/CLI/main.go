package main

import (
	"net"

	"github.com/moelksasbyahmed/go_loadbalancer/internal"
)

var Admin_Url string

func main() {
	Adminconfig, _ := internal.LoadConfig("./config.yaml")
	Admin_Url = "http://" + net.JoinHostPort(Adminconfig.Adminconfig.Host, Adminconfig.Adminconfig.Port)
	Execute()

}
