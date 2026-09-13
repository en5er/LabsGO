package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func average(a, b, c float64) float64 {
	return (a + b + c) / 3
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите три числа (через пробел или каждое с новой строки):")

	line, _ := reader.ReadString('\n')
	fields := strings.Fields(line)

	if len(fields) < 3 {
		fmt.Println("Ошибка: нужно ввести три числа.")
		return
	}

	nums := make([]float64, 3)
	for i := 0; i < 3; i++ {
		v, err := strconv.ParseFloat(fields[i], 64)
		if err != nil {
			fmt.Printf("Ошибка: '%s' не является числом.\n", fields[i])
			return
		}
		nums[i] = v
	}

	avg := average(nums[0], nums[1], nums[2])
	fmt.Printf("Среднее значение %.0f, %.0f и %.0f = %.2f\n",
		nums[0], nums[1], nums[2], avg)
}
