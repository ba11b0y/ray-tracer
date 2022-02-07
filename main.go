package main

import (
	"fmt"
	"strings"
)

const width, height = 256, 256

func main() {
	renderFile()
}

func renderFile() {
	var r, g, b float64
	var ir, ig, ib int

	// use string builder for efficient concatenation of strings https://pkg.go.dev/strings#Builder
	var fileContent strings.Builder

	fileContent.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", width, height))
	for j := height - 1; j >= 0; j-- {
		for i := 0; i < width; i++ {
			r = float64(i) / (width - 1)
			g = float64(j) / (height - 1)
			b = 0.25

			ir = int(255.999 * r)
			ig = int(255.999 * g)
			ib = int(255.999 * b)

			fileContent.WriteString(fmt.Sprintf("%d %d %d\n", ir, ig, ib))
		}
	}

	fmt.Println(fileContent.String())

}
