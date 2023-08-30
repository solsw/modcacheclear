package main

import (
	"flag"
	"io"
	"log/slog"

	"github.com/solsw/modhelper"
)

var (
	pattern   string
	toKeep    int
	removeAll bool
	skipped   bool
	noOutput  bool
)

func initFlags() {
	flag.StringVar(&pattern, "p", "", "regular expression `pattern` to match module path")
	flag.IntVar(&toKeep, "v", 0, "number of module versions to keep")
	flag.BoolVar(&removeAll, "a", false, "remove all module versions")
	flag.BoolVar(&skipped, "printskipped", false, "print skipped module paths")
	flag.BoolVar(&noOutput, "nooutput", false, "produce no output except errors")
	flag.Parse()
}

func main() {
	initFlags()
	var w io.Writer
	if !noOutput {
		w = flag.CommandLine.Output()
	}
	err := modhelper.ModCacheClear(pattern, toKeep, removeAll, skipped, w)
	if err != nil {
		slog.Error(err.Error())
	}
}
