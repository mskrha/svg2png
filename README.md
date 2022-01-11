[![Go Report Card](https://goreportcard.com/badge/github.com/mskrha/svg2png)](https://goreportcard.com/report/github.com/mskrha/svg2png)

## svg2png

### Description
Very simple SVG to PNG converter library using the Inkscape.

### Installation
`go get github.com/mskrha/svg2png`

### Usage
```go
package main

import (
	"fmt"

	"github.com/mskrha/svg2png"
)

func main() {
	var input []byte

	// Fill the input with SVG image

	converter := svg2png.New()

	output, err := converter.Convert(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read the output PNG image
}
```
