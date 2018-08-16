package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/suite911/slurp911"

	"github.com/suite911/term911/vt"
)

func main() {
	outPath, pkgName, varName := "-", "main", "Slurped"
	flag.StringVar(&outPath, "o", "-", "Output path (\"-\" for stdout)")
	flag.StringVar(&pkgName, "p", "main", "Package name")
	flag.StringVar(&varName, "n", "Slurped", "Variable name")
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, usage)
		return
	}

	if err := slurp911.Main(os.Args[0], outPath, args, pkgName, varName); err != nil {
		fmt.Fprintln(os.Stderr, usage)
		log.Fatalln(err)
	}
}

var usage string

func init() {
	usage = "usage: " + os.Args[0] + " " + vt.U("OUTPATH") + " " +
		vt.U("PKGNAME") + " " + vt.U("VARNAME")  " " +
		vt.U("KEY") + ":" + vt.U("PATH") + " " + vt.S("[") +
		vt.U("KEY") + ":" + vt.U("PATH") + " " + vt.S("[...]]")
}
