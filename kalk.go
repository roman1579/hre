package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		var (
			num1      int    // первый операнд
			num2      int    // второй операнд
			i         int    // счетчик в циклах
			num1_type string //тип первого операнда
			num2_type string //тип второго операнда
			a         string //принимает строку из консоли
		)

		romeToArab := map[string]int{
			"I":    1,
			"II":   2,
			"III":  3,
			"IV":   4,
			"V":    5,
			"VI":   6,
			"VII":  7,
			"VIII": 8,
			"IX":   9,
			"X":    10,
		}
		arabToRome := map[int]string{
			1:   "I",
			2:   "II",
			3:   "III",
			4:   "IV",
			5:   "V",
			6:   "VI",
			7:   "VII",
			8:   "VIII",
			9:   "IX",
			10:  "X",
			11:  "XI",
			12:  "XII",
			13:  "XIII",
			14:  "XIV",
			15:  "XV",
			16:  "XVI",
			17:  "XVII",
			18:  "XVIII",
			19:  "XIX",
			20:  "XX",
			21:  "XXI",
			22:  "XXII",
			23:  "XXIII",
			24:  "XXIV",
			25:  "XXV",
			26:  "XXVI",
			27:  "XXVII",
			28:  "XXVIII",
			29:  "XXIX",
			30:  "XXX",
			31:  "XXXI",
			32:  "XXXII",
			33:  "XXXIII",
			34:  "XXXIV",
			35:  "XXXV",
			36:  "XXXVI",
			37:  "XXXVII",
			38:  "XXXVIII",
			39:  "XXXIX",
			40:  "XC",
			41:  "XCI",
			42:  "XCII",
			43:  "XCIII",
			44:  "XCIV",
			45:  "XCV",
			46:  "XCVI",
			47:  "XCVII",
			48:  "XCVIII",
			49:  "XCIX",
			50:  "C",
			51:  "CI",
			52:  "CII",
			53:  "CIII",
			54:  "CIV",
			55:  "CV",
			56:  "CVI",
			57:  "CVII",
			58:  "CVIII",
			59:  "CIX",
			60:  "CX",
			61:  "CXI",
			62:  "CXII",
			63:  "CXIII",
			64:  "CXIV",
			65:  "CXV",
			66:  "CXVI",
			67:  "CXVII",
			68:  "CXVIII",
			69:  "CXIX",
			70:  "CXX",
			71:  "CXXI",
			72:  "CXXII",
			73:  "CXXIII",
			74:  "CXXIV",
			75:  "CXXV",
			76:  "CXXVI",
			77:  "CXXVII",
			78:  "CXXVIII",
			79:  "CXXIX",
			80:  "CXXX",
			81:  "CXXXI",
			82:  "CXXXII",
			83:  "CXXXIII",
			84:  "CXXXIV",
			85:  "CXXXV",
			86:  "CXXXVI",
			87:  "CXXXVII",
			88:  "CXXXVIII",
			89:  "CXXXIX",
			90:  "XL",
			91:  "XLI",
			92:  "XLII",
			93:  "XLIII",
			94:  "XLIV",
			95:  "XLV",
			96:  "XLVI",
			97:  "XLVII",
			98:  "XLVIII",
			99:  "XLIX",
			100: "L",
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Введите выражение:")

		a, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			return
		}

		// Удаление символа новой строки (если есть)
		a = strings.TrimSuffix(a, "\n")
		a = strings.TrimSuffix(a, "\r")

		// Разделение строки на слова
		aa := strings.Split(a, " ")

		// Проверка, есть ли хотя бы одно слово
		if len(aa) == 1 {
			panic("Нет слов в вводе")
		}
		if len(aa) > 3 {
			panic("неверный формат ввода\n<операнд(число)> <оператор(+ - * /)> <операнд(число)>")
		}

		operand := aa[1]
		// Конвертация первого слова в число
		for i = 1; i < 11; i++ {
			if aa[0] == arabToRome[i] {
				num1 = romeToArab[aa[0]]
				num1_type = "rim"
				break
			} else if i == 10 {
				num1, err = strconv.Atoi(aa[0])
				if err != nil {
					fmt.Println("Ошибка при конвертации первого слова в число:", err)
					return
				}
				num1_type = "arab"
				break
			}
		}
		//конвертация второго слова в число
		for i = 1; i < 11; i++ {
			if aa[2] == arabToRome[i] {
				num2 = romeToArab[aa[2]]
				num2_type = "rim"
				break
			} else if i == 10 {
				num2, err = strconv.Atoi(aa[2])
				if err != nil {
					fmt.Println("Ошибка при конвертации первого слова в число:", err)
					return
				}
				num2_type = "arab"
				break
			}
		}

		//проверка остальных условий
		if 1 > num1 || num1 > 10 || 1 > num2 || num1 > 10 {
			panic("числа должны быть от 1 до 10")
		}

		if num1_type == "arab" {
			if num2_type == "arab" {
				if operand == "+" {
					fmt.Println(num1 + num2)
				} else if operand == "-" {
					fmt.Println(num1 - num2)
				} else if operand == "/" {
					fmt.Println(num1 / num2)
				} else if operand == "*" {
					fmt.Println(num1 * num2)
				} else {
					panic("неправильный оператор")
				}
			} else {
				panic("неверно введенные числа")
			}
		} else if num1_type == "rim" {
			if num2_type == "rim" {
				if operand == "+" {
					fmt.Println(arabToRome[num1+num2])
				} else if operand == "-" {
					res := num1 - num2
					if res < 0 {
						panic("не существует отрицательных римских чисел")
					} else {
						fmt.Println(arabToRome[res])
					}
				} else if operand == "/" {
					fmt.Println(arabToRome[num1/num2])
				} else if operand == "*" {
					fmt.Println(arabToRome[num1*num2])
				} else {
					panic("неправильный оператор")
				}
			} else {
				panic("неверно введенные числа")
			}
		} else {
			panic("неправильный ввод")
		}
	}
}
