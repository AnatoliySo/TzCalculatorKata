package main

import (
	"errors"
	"fmt"
	"strings"
)

func RimskNum(s string) bool {
	romanNumerals := map[string]bool{
		"I":    true,
		"II":   true,
		"III":  true,
		"IV":   true,
		"V":    true,
		"VI":   true,
		"VII":  true,
		"VIII": true,
		"IX":   true,
		"X":    true,
	}
	return romanNumerals[s]
}

// Только ли из цифр строка
func AllString(s string) bool {
	for _, char := range s {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

func ArabicNum(s string) int {
	arabicNumbers := map[string]int{
		"1":  1,
		"2":  2,
		"3":  3,
		"4":  4,
		"5":  5,
		"6":  6,
		"7":  7,
		"8":  8,
		"9":  9,
		"10": 10,
	}
	return arabicNumbers[s]
}

func main() {
	fmt.Print("Введите пример но только без пробела) : ")
	var input string
	fmt.Scanln(&input)
	input = strings.ReplaceAll(input, " ", "")

	deystvia := strings.Split(input, "+")
	if len(deystvia) != 2 {
		deystvia = strings.Split(input, "-")
		if len(deystvia) != 2 {
			deystvia = strings.Split(input, "*")
			if len(deystvia) != 2 {
				deystvia = strings.Split(input, "/")
				if len(deystvia) != 2 {
					fmt.Println("Ошибка: Неправильный формат ввода")
					return
				}
			}
		}
	}

	operand1 := deystvia[0]
	operator := ""
	operand2 := deystvia[1]

	if strings.Contains(input, "+") {
		operator = "+"
	} else if strings.Contains(input, "-") {
		operator = "-"
	} else if strings.Contains(input, "*") {
		operator = "*"
	} else if strings.Contains(input, "/") {
		operator = "/"
	}

	// арабские числами
	ArabOperand1 := isArabicNumber(operand1)
	ArabOperand2 := isArabicNumber(operand2)

	// римские числами
	RimskOperand1 := RimskNum(operand1)
	RimskOperand2 := RimskNum(operand2)

	var result string
	if ArabOperand1 && ArabOperand2 {
		result = RunArabicOperation(operand1, operator, operand2)
	} else if RimskOperand1 && RimskOperand2 {
		result, _ = RunRomskiOperation(operand1, operator, operand2)
	} else {
		fmt.Println("Ошибка: Неправильный формат операндов")
		return
	}

	fmt.Println("Результат:", result)
}

func isArabicNumber(s string) bool {
	// Арабские ли числа
	return AllString(s)
}

func toRimsk(num int) string {
	if num <= 0 || num > 10 {
		return ""
	}
	romanNumerals := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	return romanNumerals[num-1]
}

func romanToArabic(roman string) int {
	romanNumerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
	}
	result := 0
	prevValue := 0
	for i := len(roman) - 1; i >= 0; i-- {
		value := romanNumerals[rune(roman[i])]
		if value < prevValue {
			result -= value
		} else {
			result += value
		}
		prevValue = value
	}
	return result
}

// Выполнение с арабскими
func RunArabicOperation(operand1, operator, operand2 string) string {
	num1 := ArabicNum(operand1)
	num2 := ArabicNum(operand2)
	switch operator {
	case "+":
		return fmt.Sprintf("%d", num1+num2)
	case "-":
		return fmt.Sprintf("%d", num1-num2)
	case "*":
		return fmt.Sprintf("%d", num1*num2)
	case "/":
		if num2 == 0 {
			return "Ошибка: Деление на ноль"
		}
		return fmt.Sprintf("%d", num1/num2)
	default:
		return "Ошибка: Недопустимая операция"
	}
}

// Выполнение с римскими
func RunRomskiOperation(operand1, operator, operand2 string) (string, error) {
	num1 := romanToArabic(operand1)
	num2 := romanToArabic(operand2)
	result := 0
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			return "", errors.New("Деление на ноль")
		}
		result = num1 / num2
	default:
		return "", errors.New("Недопустимая операция")
	}
	if result < 1 {
		return "", errors.New("Результат операции не может быть представлен римским числом")
	}
	return toRimsk(result), nil
}
