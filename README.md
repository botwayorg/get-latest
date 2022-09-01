# get-latest

> Get the latest repository version

## Usage

### routes

1. for public repos -> `/:user/:repo`
2. for private repos -> `/:user/:repo/:token`

### from website

visit [**get-latest website**](https://get-latest.onrender.com) and enter your repository name.

### use api

```go
package main

import (
	"github.com/botwayorg/get-latest/api"
)

func main() {
	latest := api.LatestWithArgs("railwayapp/cli", "GITHUB_TOKEN", false)
	//                           ☝ repo name      ☝ github token  ☝ remove 'v' character from tag
	println(latest)
}
```

## License

Apache-2.0 License - [**License**](https://github.com/botwayorg/get-latest/blob/main/LICENSE)
