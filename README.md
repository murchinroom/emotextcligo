# emotextcligo

Emotext Golang SDK: a client to [emotext](https://github.com/murchinroom/emotext).

> 🔑 This is a muli component. You can find the main repository [here](https://github.com/cdfmlr/muvtuber).

## Usage

Install:

```go
go get github.com/murchinroom/emotextcligo
```

Example:

```go
package main

import (
	"fmt"
	"github.com/murchinroom/emotextcligo"
)

func main() {
    # emotextcligo.EmotextServer = "http://localhost:51061" or set env EMOTEXT_SERVER
    text := "抱歉，今天很开心！"

    res, _ := emotextcligo.Query(text)
    # res.Emotions = emotextcligo.Emotions21To7(res.Emotions)

    fmt.Println("Emotions:", res.Emotions)
    fmt.Println("Polarity:", res.Polarity)
}
```

## License

MIT

