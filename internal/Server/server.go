package server

type backend struct {
	name        string
	url         string
	maxLoad     int
	currentLoad int
	status      string
	timeout     int
}

type loadBalancer struct {
	servers       []backend
	overallStatus string
}
