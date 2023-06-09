package main

import "fmt"

func main() {
	// fmt.Printf("123\f456\f789")
	fmt.Printf("\vf\v\a")
	var s string = "Это строка"

	// Однако длина строки len(s) составит 19 байт, т.к. использованные кирилические символы
	// занимают 2 байта, а пробел занимает 1 байт.
	fmt.Printf("Длина строки: %d байт\n", len(s))

	// Посмотрим как взять подстроку
	fmt.Printf("Напечатаем только второе слово в кавычках: \"%v\"\n", s[7:])

	/*
		Попробуем изменить что-то встроке:
		s[3] = 12
		Ошибка компиляции: cannot assign to s[3], потому что строки - неизменяемые последовательности.
	*/

	// "Изменим строку", создав новую
	s = s + " Новая строка"
	fmt.Printf("%v\n", s)

	// А теперь проитерируемся по этой строке
	for _, b := range s {
		fmt.Printf("%v ", b)
	}
}
