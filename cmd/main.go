package main

import (
	"github.com/jonnay101/ginServer/pkg/ginserver"
)

func main() {
	svc := ginserver.New("/people", "", ":3000")

	svc.StartServer()
}
