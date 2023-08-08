package main

import (
	"github.com/botwayorg/get-latest/api"
)

func main() {
	latest := api.LatestWithArgs("denoland/deno", "", false)
	println(latest)
}
