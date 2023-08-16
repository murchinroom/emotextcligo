package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/murchinroom/emotextcligo"
)

func main() {
    kind21 := flag.Bool("21", false, "use 21 kinds of emotions (default)")
    kind7 := flag.Bool("7", false, "use 7 kinds of emotions (override -21)")
    
    flag.Parse()

    switch {
    case *kind7 && *kind21:
        fmt.Fprintln(os.Stderr, "Expect only one of -21 or -7")
        os.Exit(2)
    case !*kind7 && !*kind21:
        *kind21 = true
    }

    args := flag.Args()
    text := strings.Join(args, " ")

    res, err := emotextcligo.Query(text)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(-1)
    }

    if *kind7 {
        res.Emotions = emotextcligo.Emotions21To7(res.Emotions)
    }

    fmt.Println("Emotions:", res.Emotions)
    fmt.Println("Polarity:", res.Polarity)
}
