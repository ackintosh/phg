package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Evaluate(e Expression) (string, error) {
	switch e.(type) {
	case BinOpExpr:
		left, err := Evaluate(e.(BinOpExpr).left)
		if err != nil {
			panic(err)
		}

		right, err := Evaluate(e.(BinOpExpr).right)
		if err != nil {
			panic(err)
		}

		switch e.(BinOpExpr).operator {
		case '+':
			lefti, err := strconv.Atoi(left)
			if err != nil {
				panic(err)
			}
			righti, err := strconv.Atoi(right)
			if err != nil {
				panic(err)
			}
			result := lefti + righti
			evaluated := strconv.Itoa(result)
			return evaluated, nil
		}
	case NumExpr:
		return e.(NumExpr).literal, nil
	case StringExpr:
		return strings.Trim(e.(StringExpr).literal, "\""), nil
	case EchoStatement:
		result, err := Evaluate(e.(EchoStatement).expr)
		if err != nil {
			panic(err)
		}
		fmt.Print(result)
		return "", nil
	}

	return "", nil
}
