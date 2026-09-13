package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func average(a, b int) float64 {
	return float64(a+b) / 2
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите два целых числа (через пробел):")

	line, _ := reader.ReadString('\n')
	fields := strings.Fields(line)

	if len(fields) < 2 {
		fmt.Println("Ошибка: нужно ввести два числа.")
		return
	}

	a, err1 := strconv.Atoi(fields[0])
	b, err2 := strconv.Atoi(fields[1])

	if err1 != nil || err2 != nil {
		fmt.Println("Ошибка: оба значения должны быть целыми числами.")
		return
	}

	avg := average(a, b)
	fmt.Printf("Среднее значение %d и %d = %.2f\n", a, b, avg)
}