package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// positive / negative / zero
func opredelitZnak(chislo int) string {
	if chislo > 0 {
		return "Positive"
	} else if chislo < 0 {
		return "Negative"
	}
	return "Zero"
}

// длина строки
func dlinaStroki(stroka string) int {
	return len(stroka)
}

type Rectangle struct {
	shirina float64
	vysota  float64
}

// площадь прямоугольника
func (p Rectangle) Ploshad() float64 {
	return p.shirina * p.vysota
}

// среднее двух целых
func sredneeDvuh(chislo1, chislo2 int) float64 {
	return float64(chislo1+chislo2) / 2
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Задание 1
	fmt.Println("Задание 1.")
	fmt.Print("Введите число: ")
	stroka1, _ := reader.ReadString('\n')
	chislo1, oshibka1 := strconv.Atoi(strings.TrimSpace(stroka1))
	if oshibka1 != nil {
		fmt.Println("Нужно целое число")
	} else if chislo1%2 == 0 {
		fmt.Printf("Число %d чётное\n", chislo1)
	} else {
		fmt.Printf("Число %d нечётное\n", chislo1)
	}

	// Задание 2
	fmt.Println("\n Задание 2.")
	fmt.Print("Введите число: ")
	stroka2, _ := reader.ReadString('\n')
	chislo2, oshibka2 := strconv.Atoi(strings.TrimSpace(stroka2))
	if oshibka2 != nil {
		fmt.Println("Нужно целое число")
	} else {
		fmt.Println("Результат:", opredelitZnak(chislo2))
	}

	// Задание 3
	fmt.Println("\n Задание 3.")
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Задание 4
	fmt.Println("\n Задание 4.")
	fmt.Print("Введите строку: ")
	stroka4, _ := reader.ReadString('\n')
	stroka4 = strings.TrimSpace(stroka4)
	fmt.Println("Длина строки:", dlinaStroki(stroka4))

	// Задание 5
	fmt.Println("\n Задание 5.")
	priamougolnik := Rectangle{shirina: 5.5, vysota: 3.2}
	fmt.Printf("Ширина = %.1f, высота = %.1f\n", priamougolnik.shirina, priamougolnik.vysota)
	fmt.Printf("Площадь = %.2f\n", priamougolnik.Ploshad())

	// Задание 6
	fmt.Println("\n Задание 6.")
	fmt.Print("Введите два целых числа через пробел: ")
	stroka6, _ := reader.ReadString('\n')
	chasti6 := strings.Fields(stroka6)

	if len(chasti6) < 2 {
		fmt.Println("Нужно два числа")
		return
	}

	chislo6a, oshibka6a := strconv.Atoi(chasti6[0])
	chislo6b, oshibka6b := strconv.Atoi(chasti6[1])

	if oshibka6a != nil || oshibka6b != nil {
		fmt.Println("Оба числа должны быть целыми")
		return
	}

	fmt.Printf("Среднее значение = %.2f\n", sredneeDvuh(chislo6a, chislo6b))
}