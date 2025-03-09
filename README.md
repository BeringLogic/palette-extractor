palette-extractor
=================
[![Go Report Card](https://goreportcard.com/badge/github.com/thomas-bouvier/palette-extractor)](https://goreportcard.com/report/thomas-bouvier/palette-extractor)
[![GoDoc](https://godoc.org/github.com/thomas-bouvier/palette-extractor?status.svg)](https://godoc.org/github.com/thomas-bouvier/palette-extractor)

This program extracts the dominant color or a representative color palette from an image.

## Usage

Here's a simple example, where we build a 5 color palette:

```go
package main

import (
    "os"
    "fmt"
    "image"
    _ "image/gif"
    _ "image/jpeg"
    _ "image/png"

    _ "golang.org/x/image/webp"

    "github.com/BeringLogic/palette-extractor"
)

func main() {
    // opening the image
    file, err := os.Open("image.png"); if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    // decoding the image
    img, _, err := image.Decode(file); if err != nil {
        fmt.Println(err)
        return
    }

    // Creating the extractor object
    extractor := extractor.NewExtractor(img, 10)

    // Displaying the top 5 dominant colors of the image
    fmt.Println(extractor.GetPalette(5))
}
```

You can find [the complete documentation](https://godoc.org/github.com/thomas-bouvier/palette-extractor) on GoDoc.

## Example

The following image has been used for this example:

![Example](image.png)

The program will give the following output when used with the image above:

```
[[234 231 230] [208 24 44] [59 41 37] [158 149 145] [145 126 114]]
```

## Thanks

Many thanks to [Thomas Bouvier](https://github.com/thomas-bouvier) for his [palette-extractor](https://github.com/thomas-bouvier/palette-extractor), from which this has been forked

Many thanks to [Lokesh Dhakar](https://github.com/lokesh) for [his original work](https://github.com/lokesh/color-thief/) and [Shipeng Feng](https://github.com/fengsp) for [his implementation](https://github.com/fengsp/color-thief-py).
