package main

import (
	"github.com/scmn-dev/get-latest/api"
)

func main() {
	latest := api.LatestWithArgs("railwayapp/cli", "")
	println(latest)
}
