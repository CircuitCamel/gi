package main

import (
	"gitignore/internal/server"
	"gitignore/internal/utils"
)

func main() {
	server.Start(utils.LoadConfig())
}
