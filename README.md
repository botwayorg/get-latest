# Get Latest 🛰️

> Get the latest repository version. Written in Deno Fresh 🦕

## Routes

1. for public repos -> `/:user/:repo`
2. for private repos -> `/:user/:repo?token=TOKEN`
3. you can add `?no-v=true` query var to receive release without `v` char

## Usage

### JS/TS

```js
fetch("https://get-latest.deno.dev/denoland/deno")
  .then(res => res.text())
  .then(data => console.log(data))
  .catch(error => console.error('Error:', error));
```

### Deno

```ts
const res = await fetch("https://get-latest.deno.dev/denoland/deno");

const data = await res.text();

console.log(data);
```

### Golang

```go
package main

import (
	"github.com/botwayorg/get-latest/api"
)

func main() {
	latest := api.LatestWithArgs("denoland/deno", "GITHUB_TOKEN", false)
	//                           ☝ repo name     ☝ github token ☝ remove 'v' character from tag
	println(latest)
}
```

> And of course you can use it with other languages 🤝

## License

Apache-2.0 License - [**License**](https://github.com/botwayorg/get-latest/blob/main/LICENSE)
