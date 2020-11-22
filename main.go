package main

import (
	"github.com/w3-server/routes"
)

func main() {
	r := routes.InitRouter()
	r.Run(":2020")
}
