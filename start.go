package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Input:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// Разделение входной строки на части
	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		fmt.Println("Неверный формат ввода")
		return
	}

	isRoman1 := strings.ContainsAny(parts[0], "IVXLCDM")
	isRoman2 := strings.ContainsAny(parts[2], "IVXLCDM")
	if isRoman1 != isRoman2 {
		panic("Числа разных систем исчисления")
	}

	if isRoman1 {
		num1 := fromRoman(parts[0])
		num2 := fromRoman(parts[2])
		// Выполнение операции
		switch parts[1] {
		case "+":
			result := num1 + num2
			fmt.Println("Output:\n", toRoman(result))
		case "-":
			result := num1 - num2
			fmt.Println("Output:\n", toRoman(result))
			if num1-num2 <= 0 {
				panic("Результат меньше или равен нулю")
			}
		case "*":
			result := num1 * num2
			fmt.Println("Output:\n", toRoman(result))
		case "/":
			if num2 == 0 {
				fmt.Println("Деление на ноль")
				return
			}
			result := num1 / num2
			fmt.Println("Output:\n", toRoman(result))
		default:
			panic("Неверная операция")
		}
	} else {
		// Преобразование строк в числа
		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[2])

		if num1 > 10 || num2 > 10 || num1 < 1 || num2 < 1 {
			fmt.Println("Введите числа от 1 до 10 включительно")
			return
		}
		// Выполнение операции
		switch parts[1] {
		case "+":
			fmt.Println("Output:\n", num1+num2)
		case "-":
			fmt.Println("Output:\n", num1-num2)
		case "*":
			fmt.Println("Output:\n", num1*num2)
		case "/":
			if num2 == 0 {
				fmt.Println("Деление на ноль")
				return
			}
			fmt.Println("Output:\n", num1/num2)
		default:
			panic("Неверная операция")
		}
	}
}

// Преобразование римских чисел в арабские
func fromRoman(s string) int {
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	total := 0
	prev := 0
	for _, r := range s {
		value := romanMap[string(r)]
		if value > prev {
			total += value - 2*prev
		} else {
			total += value
		}
		prev = value
	}
	return total
}

func toRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var roman strings.Builder
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			num -= values[i]
			roman.WriteString(symbols[i])
		}
	}
	return roman.String()
}
