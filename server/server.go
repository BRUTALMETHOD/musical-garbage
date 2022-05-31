package server

import (
	"os"

	"github.com/brutalmethod/musicalgarbage/routers"
)

func Init() {
	serverName := os.Getenv("GIN_ADDR")
	serverPort := os.Getenv("GIN_PORT")
	r := routers.NewRouter()
	r.Run(serverName + ":" + serverPort)
}
