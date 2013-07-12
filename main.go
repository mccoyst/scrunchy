package main

import (
	"flag"
	"io"
	"os"
	"text/tabwriter"
)

var (
	minwidth   = flag.Int("min", 0, "minimum field width")
	tabwidth   = flag.Int("tab", 8, "tab width")
	padding    = flag.Int("pad", 1, "extra padding between fields")
	padchar    = flag.String("padch", " ", "padding character")
	alignright = flag.Bool("right", false, "align fields to the right")
)

func main() {
	flag.Parse()
	if len(*padchar) < 1 {
		os.Stderr.WriteString("padchar must not be empty")
		os.Exit(1)
	}
	flags := uint(0)
	if *alignright {
		flags |= tabwriter.AlignRight
	}
	w := tabwriter.NewWriter(os.Stdout, *minwidth, *tabwidth, *padding, (*padchar)[0], flags)
	defer w.Flush()
	io.Copy(w, os.Stdin)
}
