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

type MyCase struct {
	op       []rune
	nOperand int
	nBrace   int
	total    int
	maxSum   int
	maxAdder int
	opCount  map[rune]int
}

func (tm *MyCase) AddOperator(op rune) error {
	if !validOp[op] {
		return fmt.Errorf("invalid operator: %c", op)
	}
	tm.op = append(tm.op, op)
	return nil
}

func (tm *MyCase) AddOperatorStr(opstr string) error {
	for _, op := range opstr {
		if err := tm.AddOperator(op); err != nil {
			return err
		}
	}
	return nil
}

func (tm *MyCase) SetNumberOfOperand(n int) { tm.nOperand = n }
func (tm *MyCase) SetNumberOfCase(n int)    { tm.total = n }
func (tm *MyCase) SetMaxOperand(n int)      { tm.maxSum = n }
func (tm *MyCase) SetMaxAdder(n int)        { tm.maxAdder = n }
func (tm *MyCase) CountOp(op rune) {
	if tm.opCount == nil {
		tm.opCount = make(map[rune]int)
	}
	if n, ok := tm.opCount[op]; ok {
		tm.opCount[op] = n + 1
	} else {
		tm.opCount[op] = 1
	}
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

func nDigit(a int) int {
	str := fmt.Sprintf("%d", a)
	return len(str)
}

func opd1fmt(maxSum int) string {
	nd := nDigit(maxSum)
	return fmt.Sprintf("%c-%d%c", '%', nd, 'd')
}

func opd2fmt(maxSum int) string {
	opd1f := opd1fmt(maxSum)
	return "%c " + opd1f
}

func (tm *MyCase) DoCase() ([]string, []string) {
	var result []string
	var result2 []string
	opd1f := opd1fmt(tm.maxSum)
	opd2f := opd2fmt(tm.maxSum)

	fmt.Printf("opd format: %v,%v\n", opd1f, opd2f)

	for i := 0; i < tm.total; i++ {
		row := ""
		maxOpd := tm.maxSum
		allOpd := 0
		lastOpd := 0
		_ = lastOpd
		nextOp := '.'
		for j := 0; j < tm.nOperand; j++ {
			if j == 0 {
				// first number
				nextOp = tm.op[rand.Intn(len(tm.op))]
				opd := rand.Intn(tm.maxSum + 1)
				switch nextOp {
				case OP_ADD, OP_MUL:
					if tm.maxAdder > 0 {
						opd = rand.Intn(tm.maxAdder + 1)
					}
				}
				row += fmt.Sprintf(opd1f, opd)
				allOpd = opd
				lastOpd = opd
				continue
			}

			// next number
			op := nextOp
			tm.CountOp(op)
			opd2 := 0
			switch op {
			case OP_ADD:
				maxOpd = tm.maxSum - allOpd
				if maxOpd < 0 {
					maxOpd = 0
				}
				if tm.maxAdder > 0 && maxOpd > tm.maxAdder {
					maxOpd = tm.maxAdder
				}
				opd2 = rand.Intn(maxOpd + 1)
				allOpd += opd2
				lastOpd = opd2

			case OP_SUB:
				maxOpd = allOpd
				if maxOpd < 0 {
					maxOpd = 0
				}
				if tm.maxAdder > 0 && maxOpd > tm.maxAdder {
					maxOpd = tm.maxAdder
				}
				opd2 = rand.Intn(maxOpd + 1)
				allOpd -= opd2
				lastOpd = opd2

			case OP_MUL:
				if allOpd == 0 {
					maxOpd = tm.maxSum
				} else {
					maxOpd = tm.maxSum / allOpd
				}
				if maxOpd < 0 {
					maxOpd = 0
				}
				if tm.maxAdder > 0 && maxOpd > tm.maxAdder {
					maxOpd = tm.maxAdder
				}

				opd2 = rand.Intn(maxOpd + 1)
				if opd2 != 0 {
					allOpd /= opd2
				}
				lastOpd = opd2

			default:
				panic(fmt.Sprintf("not support operator: %c", op))
			}

			row += fmt.Sprintf(opd2f, op_print[op], opd2)

			nextOp = tm.op[rand.Intn(len(tm.op))]
		}

		r2 := row
		row += fmt.Sprintf(" %c        ", op_print['='])
		r2 += fmt.Sprintf(" %c %-5d", op_print['='], allOpd)

		result = append(result, row)
		result2 = append(result2, r2)
	}

	return result, result2
}

func chuti(tm *MyCase) []string {
	r1, _ := tm.DoCase()
	return r1
}
