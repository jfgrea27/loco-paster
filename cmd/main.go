package main

import (
	"github.com/jfgrea27/loco-paster/internal/api"
	"github.com/jfgrea27/loco-paster/internal/utils"
)

func main() {
	api.RunServer(utils.BuildEndpoint())
}
