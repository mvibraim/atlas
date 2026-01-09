package main

import (
	"atlas/configs"
	"atlas/routes"
)

func main() {
	configs.DatabaseStart()
	routes.Start()
}
