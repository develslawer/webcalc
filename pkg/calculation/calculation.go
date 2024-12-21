package calculation

import (
	"strconv"
	"strings"
	"unicode"
)

func isOp(op rune) bool {
	return op == '-' || op == '+' || op == '*' || op == '/' || op == '(' || op == ')'
}

func getOpPriority(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	default:
		return 0
	}
}

func applyOp(a, b float64, op rune) (float64, error) {
	switch op {
	case '+':
		return a + b, nil
	case '-':
		return a - b, nil
	case '*':
		return a * b, nil
	case '/':
		if b != 0 {
			return a / b, nil
		}
		return 0, ErrDivisionByZero
	default:
		return 0, ErrUnknownOperator
	}
}

func Calc(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, " ", "")
	var stackValues []float64
	var stackOps []rune

	if len(expression) > 0 && expression[0] == '-' {
		expression = "0" + expression
	}

	for i := 0; i < len(expression); i++ {
		char := rune(expression[i])

		if unicode.IsDigit(char) || char == '.' || (char == '-' && (i == 0 || rune(expression[i-1]) == '(')) {
			var number strings.Builder

			if char == '-' {
				number.WriteRune(char)
				i++
			}

			for i < len(expression) && (unicode.IsDigit(rune(expression[i])) || expression[i] == '.') {
				number.WriteRune(rune(expression[i]))
				i++
			}
			i--

			num, err := strconv.ParseFloat(number.String(), 64)
			if err != nil {
				return 0, err
			}
			stackValues = append(stackValues, num)
		} else if isOp(char) {
			if char == '(' {
				stackOps = append(stackOps, char)
			} else if char == ')' {
				for len(stackOps) > 0 && stackOps[len(stackOps)-1] != '(' {
					op := stackOps[len(stackOps)-1]
					stackOps = stackOps[:len(stackOps)-1]
					value1, value2 := stackValues[len(stackValues)-2], stackValues[len(stackValues)-1]
					stackValues = stackValues[:len(stackValues)-2]
					res, err := applyOp(value1, value2, op)
					if err != nil {
						return 0, err
					}
					stackValues = append(stackValues, res)
				}
				stackOps = stackOps[:len(stackOps)-1]
			} else {
				for len(stackOps) > 0 && getOpPriority(stackOps[len(stackOps)-1]) >= getOpPriority(char) {
					op := stackOps[len(stackOps)-1]
					stackOps = stackOps[:len(stackOps)-1]
					value1, value2 := stackValues[len(stackValues)-2], stackValues[len(stackValues)-1]
					stackValues = stackValues[:len(stackValues)-2]
					res, err := applyOp(value1, value2, op)
					if err != nil {
						return 0, err
					}
					stackValues = append(stackValues, res)
				}
				stackOps = append(stackOps, char)
			}
		} else {
			return 0, ErrInvalidExpression
		}
	}

	// Выполняем оставшиеся операции
	for len(stackOps) > 0 {
		op := stackOps[len(stackOps)-1]
		stackOps = stackOps[:len(stackOps)-1]
		if len(stackValues) < 2 {
			return 0, ErrStackOverflow
		}
		value1, value2 := stackValues[len(stackValues)-2], stackValues[len(stackValues)-1]
		stackValues = stackValues[:len(stackValues)-2]
		res, err := applyOp(value1, value2, op)
		if err != nil {
			return 0, err
		}
		stackValues = append(stackValues, res)
	}

	if len(stackValues) == 1 {
		return stackValues[0], nil
	}

	return 0, ErrInvalidExpression
}
