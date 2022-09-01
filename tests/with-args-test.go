package main

import (
	"github.com/botwayorg/get-latest/api"
)

func main() {
	latest := api.LatestWithArgs("railwayapp/cli", "", false)
	println(latest)
}
