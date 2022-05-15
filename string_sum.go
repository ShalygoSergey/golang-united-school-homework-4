package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {

	fields := strings.Fields(input)

	if len(fields) == 0 {
		return "", fmt.Errorf("StringSum: %w", errorEmptyInput)
	}

	in := strings.Join(fields, "")

	//m := regexp.MustCompile(`-?\+?\[0-9\]-?\+?`)
	//if len(m.ReplaceAllString(in, "")) != 0 {
	//	return "", fmt.Errorf("StringSum: %w", ?)
	//}

	op := make([]string, 0)
	tmp := ""
	runes := []rune(in)
	for i := 0; i < len(runes); i = i + 1 {
		if string(runes[i]) == "+" || string(runes[i]) == "-" {
			if len(tmp) != 0 {
				op = append(op, tmp)
				tmp = ""
			}
			op = append(op, string(runes[i]))
		} else {
			tmp = tmp + string(runes[i])
		}
	}
	if len(tmp) != 0 {
		op = append(op, tmp)
	}

	count := 0
	operator := "+"
	value := 0
	for i := 0; i < len(op); i = i + 1 {

		if op[i] == "+" {
			operator = "+"
		} else if op[i] == "-" {
			operator = "-"
		} else {
			val, err := strconv.Atoi(op[i])
			if err != nil {
				return "", fmt.Errorf("StringSum: %w", err)
			}

			if operator == "+" {
				value = value + val
			} else {
				value = value - val
			}
			count = count + 1
		}
	}

	if count != 2 {
		return "", fmt.Errorf("StringSum: %w", errorNotTwoOperands)
	}

	return strconv.Itoa(value), nil
}
