package main

import (
	"io"
	"os"
	"text/tabwriter"
)

func main() {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, ' ', 0)
	defer w.Flush()
	io.Copy(w, os.Stdin)
}
