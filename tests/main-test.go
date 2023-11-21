package main

import (
	"os"

	"github.com/botwayorg/get-latest/api"
)

func main() {
	// Public Repo
	pb_repo := api.LatestWithArgs("denoland/deno", "", false)

	println(pb_repo)

	// With no 'v' character
	no_v := api.LatestWithArgs("cli/cli", "", true)

	println(no_v)

	// Private Repo
	pr_repo := api.LatestWithArgs("user/private_repo", os.Getenv("GITHUB_TOKEN"), false)

	println(pr_repo)
}
