# get-latest

> Get the latest repository version

## Usage

### from website

go to **https://get-latest.secman.dev** and enter your repository name.

![preview](https://user-images.githubusercontent.com/64256993/148659264-fb715223-17f2-4602-a9e9-9bbfd60f4df7.png)

### use api

```go
package main

import (
	"github.com/scmn-dev/get-latest/api"
)

func main() {
	latest := api.LatestWithArgs("railwayapp/cli", "")
	println(latest)
}
```

## License

MIT - [**License**](https://github.com/scmn-dev/get-latest/blob/main/LICENSE)
