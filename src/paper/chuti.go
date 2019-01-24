package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	OP_ADD rune = '+'
	OP_SUB rune = '-'
	OP_MUL rune = '*'
	OP_DIV rune = '/'
)

var validOp map[rune]bool = map[rune]bool{
	OP_ADD: true,
	OP_SUB: true,
	OP_MUL: true,
	OP_DIV: true,
}

var op_print map[rune]rune = map[rune]rune{
	OP_ADD: '＋',
	OP_SUB: '－',
	OP_MUL: '×',
	OP_DIV: '÷',
	'=':    '＝',
}

type timu struct {
	op         []rune
	nOperand   int
	nBrace     int
	total      int
	maxOperand int
	minOperand int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func add(n int, max int) string {
	adder := make([]int, n)
	for i := 0; i < len(adder); i++ {
		adder[i] = rand.Intn(max)
		max -= adder[i]
	}

	result := fmt.Sprintf("%-2d ", adder[0])
	for i := 1; i < len(adder); i++ {
		result += fmt.Sprintf("%c %-2d ", op_print[OP_ADD], adder[i])
	}
	result += fmt.Sprintf("%c      ", op_print['='])
	return result
}

func sub(n int, max int) string {
	adder := make([]int, n)
	for i := 0; i < len(adder); i++ {
		adder[i] = rand.Intn(max)
		if adder[i]*2 < max {

		}
		max = adder[i]
		if max <= 0 {
			max = 1
		}
	}

	result := fmt.Sprintf("%-2d ", adder[0])
	for i := 1; i < len(adder); i++ {
		result += fmt.Sprintf("%c %-2d ", op_print[OP_SUB], adder[i])
	}
	result += fmt.Sprintf("%c      ", op_print['='])
	return result
}

func chuti(tm *timu) []string {
	var result []string
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < tm.total; i++ {
		row := ""
		maxOpd := tm.maxOperand
		allOpd := 0
		for j := 0; j < tm.nOperand; j++ {
			if j == 0 {
				opd := rand.Intn(maxOpd)
				row += fmt.Sprintf("%-2d ", opd)
				allOpd = opd
				continue
			}
			op := tm.op[rand.Intn(len(tm.op))]
			switch op {
			case OP_ADD:
				maxOpd = tm.maxOperand - allOpd
				if maxOpd <= 0 {
					maxOpd = 1
				}
				opd2 := rand.Intn(maxOpd)
				row += fmt.Sprintf("%c %-2d ", op_print[OP_ADD], opd2)
				allOpd += opd2
				continue
			case OP_SUB:
				maxOpd = allOpd
				if maxOpd <= 0 {
					maxOpd = 1
				}
				opd2 := rand.Intn(maxOpd)
				row += fmt.Sprintf("%c %-2d ", op_print[OP_SUB], opd2)
				allOpd -= opd2
			}
		}
		row += fmt.Sprintf(" %c      ", op_print['='])

		result = append(result, row)
	}

	return result
}
