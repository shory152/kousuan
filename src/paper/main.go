package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Println()
	fmt.Printf("paper -n NumOfOperand -m MaxSum -a MaxAdder -o oprator \n")
	fmt.Printf("       -t total -c caption -f filename.pdf")
	fmt.Println()

	flag.PrintDefaults()
	fmt.Println()
	os.Exit(0)
}

func ErrExit(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Println()
	usage()
}

func main() {
	var nOpd int
	var maxOpd int
	var maxAdder int
	var opr string
	var total int
	var fname string
	var caption string
	var help bool

	flag.BoolVar(&help, "h", false, "show this help")
	flag.IntVar(&nOpd, "n", 2, "number of operands")
	flag.IntVar(&maxOpd, "m", 10, "max sum")
	flag.IntVar(&maxAdder, "a", 10, "max adder")
	flag.StringVar(&opr, "o", "+", "operators, e.g. +++-*/")
	flag.IntVar(&total, "t", 100, "total test case")
	flag.StringVar(&fname, "f", "kousuan.pdf", "output pdf file name")
	flag.StringVar(&caption, "c", "test", "pdf document's caption")

	flag.Parse()

	if help || len(os.Args) <= 1 {
		usage()
	}

	var tm MyCase
	tm.SetMaxOperand(maxOpd)
	tm.SetMaxAdder(maxAdder)
	tm.SetNumberOfOperand(nOpd)
	tm.SetNumberOfCase(total)

	if err := tm.AddOperatorStr(opr); err != nil {
		ErrExit("%v", err)
	}

	result := tm.DoCase()

	printpdf(result, caption, fname)

	for k, v := range tm.opCount {
		fmt.Printf("  %c: %d\n", k, v)
	}
}
