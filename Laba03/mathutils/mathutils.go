package mathutils

//факториал

func Faktorial(chislo int) int {
	if chislo < 0 {
		return 0
	}
	rezultat := 1
	for i := 2; i <= chislo; i++{
		rezultat *= i
	}
	return rezultat
}