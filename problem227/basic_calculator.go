package problem227

import (
	"strconv"
	"unicode"
)

type node struct {
	Value       string
	Operand     bool
	Left, Right *node
}

func calculate(s string) int {
	var root *node
	for i := 0; i < len(s); {
		c := s[i]
		op := string(c)
		if c >= '0' && c <= '9' {
			j := i + 1
			for ; j < len(s); j++ {
				if s[j] < '0' || s[j] > '9' {
					break
				}
			}
			op = s[i:j]
			i = j - 1
		}
		if op == "+" || op == "-" || op == "*" || op == "/" {
			root = insertOperand(root, op)
		} else if op != " " {
			root = insertOperator(root, op)
		}
		i++
	}
	return evaluateTree(root)
}

func insertOperand(root *node, operand string) *node {
	if root == nil {
		return &node{Value: operand, Operand: true}
	}
	if !root.Operand {
		return &node{Value: operand, Operand: true, Left: root}
	}
	if operand == "+" || operand == "-" {
		return &node{Value: operand, Operand: true, Left: root}
	}
	n := root
	var lastOperand *node
	for n.Right != nil && n.Right.Operand {
		lastOperand = n
		n = n.Right
	}
	if n.Value == "+" || n.Value == "-" {
		n.Right = &node{Value: operand, Operand: true, Left: n.Right}
	} else {
		newNode := &node{Value: operand, Operand: true, Left: n}
		if lastOperand != nil {
			lastOperand.Right = newNode
		} else {
			root = newNode
		}
	}
	return root
}

func insertOperator(root *node, operator string) *node {
	if root == nil {
		return &node{Value: operator}
	}
	if root.Left == nil {
		root.Left = &node{Value: operator}
	} else if root.Right == nil {
		root.Right = &node{Value: operator}
	} else {
		n := root
		for n.Right != nil {
			n = n.Right
		}
		n.Right = &node{Value: operator}
	}
	return root
}

func evaluateTree(root *node) int {
	if root == nil {
		return 0
	}
	if root.Operand {
		left := evaluateTree(root.Left)
		right := evaluateTree(root.Right)
		operand := root.Value
		if operand == "+" {
			return left + right
		} else if operand == "-" {
			return left - right
		} else if operand == "*" {
			return left * right
		} else if operand == "/" {
			return left / right
		}
	}
	value, _ := strconv.Atoi(root.Value)
	return value
}

func calculate2(s string) int {
	num := 0
	op := '+'
	operands := make([]int, 0)
	s += "+"
	for _, r := range s {
		if unicode.IsDigit(r) {
			num = num*10 + int(r-'0')
		} else if r == ' ' {
			continue
		} else {
			switch op {
			case '+':
				operands = append(operands, num)
			case '-':
				operands = append(operands, -num)
			case '*':
				last := operands[len(operands)-1]
				operands = operands[:len(operands)-1]
				operands = append(operands, last*num)
			case '/':
				last := operands[len(operands)-1]
				operands = operands[:len(operands)-1]
				operands = append(operands, last/num)
			}
			op = r
			num = 0
		}
	}
	sum := 0
	for _, num := range operands {
		sum += num
	}
	return sum
}
