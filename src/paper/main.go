package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Println()
	fmt.Printf("kousuan -n NumOfOperand -m MaxOperand -o oprator -t total -c caption -f filename.pdf")
	fmt.Println()

	flag.PrintDefaults()
	fmt.Println()
	os.Exit(0)
}

func main() {
	var nOpd int
	var maxOpd int
	var opr string
	var total int
	var fname string
	var caption string
	var help bool

	flag.BoolVar(&help, "h", false, "show this help")
	flag.IntVar(&nOpd, "n", 2, "number of operands")
	flag.IntVar(&maxOpd, "m", 10, "max operand")
	flag.StringVar(&opr, "o", "+", "operators, e.g. +-*/")
	flag.IntVar(&total, "t", 100, "total test case")
	flag.StringVar(&fname, "f", "kousuan.pdf", "output pdf file name")
	flag.StringVar(&caption, "c", "test", "pdf document's caption")

	flag.Parse()

	if help || len(os.Args) <= 1 {
		usage()
	}

	var tm timu
	tm.maxOperand = maxOpd
	tm.nOperand = nOpd
	tm.total = total
	for _, c := range opr {
		if validOp[c] {
			tm.op = append(tm.op, c)
		}
	}

	result := chuti(&tm)

	printpdf(result, caption, fname)
}
