package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// перевод из римских в арабские
func RotoAr(ro string) int {
	maprotoar := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var tempNum, sum int
	for i := len(ro) - 1; i >= 0; i-- {
		ro1 := ro[i]
		map1 := maprotoar[ro1]
		if map1 < tempNum {
			sum -= map1
		} else {
			sum += map1
			tempNum = map1
		}
	}
	return sum
}

// перевод из арабских в римские
func ArtoRo(ar int) string {
	mapartoro := []struct {
		key   int
		value string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	itog := ""
	for _, val := range mapartoro {
		for ar >= val.key {
			itog += val.value
			ar -= val.key
		}
	}
	return itog
}

func main() {
	for {
		fmt.Println("INPUT:")
		//считывание и распределение в одномерный массив символов строки
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.ReplaceAll(text, " ", "")
		n := len(text) - 1
		simv := []rune(text)
		oper := [4]int{42, 43, 45, 47} //определение массива с операторами
		//var i, j int
		var s int       //число операторов
		var op int      //оператор выражения
		var a, b string //операнды до и после оператора соответсвенно, тип - строка
		var an, bn int  //операнды до и после оператора соответсвенно, тип - целое число
		var at, bt int  // вид чисел операнды до и после оператора соответсвенно, арабское или римское
		s = 0
		op = 0
		at = 0
		bt = 0
		for i := 0; i < n-1; i++ {
			for j := 0; j < 4; j++ {
				if int(simv[i]) == oper[j] {
					s += 1
					op = oper[j]
				}
			}
			//определение принадлежности диапазону от 1 до 10 и определение вида чисел операнд
			if s == 0 {
				a += string(simv[i])
				if simv[i] > 47 && simv[i] < 58 {
					at = 0
				} else {
					at = 1
				}
			} else {
				b += string(simv[i+1])
				if simv[i+1] > 47 && simv[i+1] < 58 {
					bt = 0
				} else {
					bt = 1
				}
			}
		}
		fmt.Println("OUTPUT:")
		var cn int //результат, целое число
		if a == "" || b == "" {
			fmt.Println("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
			goto LOOP
		}
		if s > 1 {
			fmt.Println("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
			goto LOOP
		}
		if at == bt {
			if at == 0 {
				an, _ = strconv.Atoi(a)
				bn, _ = strconv.Atoi(b)
			} else {
				an = RotoAr(a)
				bn = RotoAr(b)
			}
		} else {
			fmt.Println("Выдача паники, так как используются одновременно разные системы счисления.")
			break
		}
		if s == 0 {
			fmt.Println("Выдача паники, так как строка не является математической операцией.")
			break
		}
		if !(an > 0 && an < 11) || !(bn > 0 && bn < 11) {
			fmt.Println("Калькулятор должен принимать на вход целые числа от 1 до 10 включительно.")
			break
		}
		switch op {
		case 42:
			cn = an * bn
		case 43:
			cn = an + bn
		case 45:
			cn = an - bn
		case 47:
			cn = an / bn
		}
		//вывод результата, если не было паники
		if at == 1 {
			if cn < 1 {
				fmt.Println("Выдача паники, так как в римской системе могут быть только положительные числа.")
				goto LOOP
			} else {
				fmt.Println(ArtoRo(cn))
			}
		} else {
			fmt.Println(cn)
		}
	LOOP:
	}
}
