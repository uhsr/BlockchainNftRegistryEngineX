// cmd/blockchainnftregistryenginex/main.go
package main

import (
"flag"
"log"
"os"

"blockchainnftregistryenginex/internal/blockchainnftregistryenginex"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := blockchainnftregistryenginex.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
