package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	//для своих пакетов нужно добавлять путь в импорте. Я в шоке
	"labsgo/Laba03/mathutils"
	"labsgo/Laba03/stringutils"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Задание 1
	fmt.Println("Задание 1.")
	fmt.Println("Пакет mathutils с функцией Faktorial создан.")

	// Задание 2
	fmt.Println("\n Задание 2.")
	fmt.Print("Введите число для факториала: ")
	stroka1, _ := reader.ReadString('\n')
	chislo1, oshibka1 := strconv.Atoi(strings.TrimSpace(stroka1))
	if oshibka1 != nil || chislo1 < 0 {
		fmt.Println("Нужно неотрицательное целое число")
	} else {
		fmt.Printf("Факториал %d = %d\n", chislo1, mathutils.Faktorial(chislo1))
	}

	// Задание 3
	fmt.Println("\n Задание 3.")
	fmt.Print("Введите строку для переворота: ")
	stroka3, _ := reader.ReadString('\n')
	stroka3 = strings.TrimSpace(stroka3)
	fmt.Println("Перевёрнутая строка:", stringutils.Perevernut(stroka3))

	// Задание 4
	fmt.Println("\n Задание 4.")
	massiv := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Массив:", massiv)

	// Задание 5
	fmt.Println("\n Задание 5.")
	srez := massiv[:]
	fmt.Println("Срез из массива:", srez)

	srez = append(srez, 60, 70)
	fmt.Println("После добавления 60 и 70:", srez)

	srez = append(srez[:2], srez[3:]...)
	fmt.Println("После удаления элемента с индексом 2:", srez)

	// Задание 6
	fmt.Println("\n Задание 6.")
	stroki := []string{"Enser", "программирование", "плохо", "Go", "знать"}
	fmt.Println("Срез строк:", stroki)

	samayaDlinnaya := ""
	for _, s := range stroki {
		if len(s) > len(samayaDlinnaya) {
			samayaDlinnaya = s
		}
	}
	fmt.Println("Самая длинная строка:", samayaDlinnaya)
}