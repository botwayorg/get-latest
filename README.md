# get-latest

> Get the latest repository version

## Usage

### routes

1. for public repos -> `/:user/:repo`
2. for private repos -> `/:user/:repo/:token`

### from website

go to **https://get-latest.secman.dev** and enter your repository name.

![preview](https://user-images.githubusercontent.com/64256993/148659957-57b9f6f4-fb8d-4e41-bd11-25098a428a4f.png)

### use api

```go
package main

import (
	"github.com/scmn-dev/get-latest/api"
)

func main() {
	latest := api.LatestWithArgs("railwayapp/cli", "")
	//                           ☝ repo name      ☝ github token
	println(latest)
}
```

## License

Apache-2.0 License - [**License**](https://github.com/scmn-dev/get-latest/blob/main/LICENSE)
