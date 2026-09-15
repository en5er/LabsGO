package stringutils

//переворот строки

func Perevernut(stroka string) string {
	runy := []rune(stroka)
	for i, j:= 0, len(runy)-1; i < j; i, j = i+1, j-1 {
		runy[i], runy[j] = runy[j], runy[i]
	}
	return string(runy)
}