package main

import (
  "os"
  "fmt"
  "flag"
  "image"
  _ "image/gif"
  _ "image/jpeg"
  _ "image/png"

  _ "golang.org/x/image/webp"

  "github.com/BeringLogic/palette-extractor"
)

func main() {
  quality := flag.Int("q", 1, "quality")
  nbColors := flag.Int("n", 3, "number of colors to extract")
  flag.Parse()

  tail := flag.Args()
  if len(tail) == 0 {
    fmt.Println("Usage: palette-extractor [-q <quality>] [-n <nbColors>] <filename>")
    return
  }

  filename := tail[0]

  file, err := os.Open(filename); if err != nil {
    fmt.Println(err)
    return
  }
	defer file.Close()

  img, _, err := image.Decode(file); if err != nil {
    fmt.Println(err)
    return
  }

  extractor, err := extractor.NewExtractor(img, *quality); if err != nil {
    fmt.Println(err)
    return
  }

  palette := extractor.GetPalette(*nbColors)
  for _, color := range palette {
    fmt.Println(color)
  }
}
