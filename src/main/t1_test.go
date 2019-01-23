package main

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	var tm timu
	tm.maxOperand = 20
	tm.nOperand = 2
	tm.total = 1000
	tm.op = append(tm.op, OP_ADD)
	tm.op = append(tm.op, OP_SUB)
	//tm.op = append(tm.op, OP_MUL)
	//tm.op = append(tm.op, OP_DIV)

	result := chuti(&tm)
	fmt.Println(result)

	printpdf(result, "hello.pdf")
}

func Test2(t *testing.T) {
	var tm timu
	tm.maxOperand = 20
	tm.nOperand = 3
	tm.total = 1000
	tm.op = append(tm.op, OP_ADD)
	tm.op = append(tm.op, OP_SUB)

	result := chuti(&tm)
	fmt.Println(result)

	printpdf(result, "hello2.pdf")
}

func Test3(t *testing.T) {
	var tm timu
	tm.maxOperand = 20
	tm.nOperand = 4
	tm.total = 200
	tm.op = append(tm.op, OP_ADD)
	tm.op = append(tm.op, OP_SUB)

	result := chuti(&tm)
	fmt.Println(result)

	printpdf(result, "hello3.pdf")
}
