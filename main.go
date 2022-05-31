package main

import (
	"fmt"
	"os"

	"github.com/brutalmethod/musicalgarbage/server"
	"github.com/brutalmethod/musicalgarbage/utils"
)

func main() {

	utils.EnvInit()
	fmt.Println(os.Getenv("GIN_ADDR"))
	server.Init()

}
