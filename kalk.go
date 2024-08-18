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
			40:  "XL",
			41:  "XLI",
			42:  "XLII",
			43:  "XLIII",
			44:  "XLIV",
			45:  "XLV",
			46:  "XLVI",
			47:  "XLVII",
			48:  "XLVIII",
			49:  "XLIX",
			50:  "L",
			51:  "LI",
			52:  "LII",
			53:  "LIII",
			54:  "LIV",
			55:  "LV",
			56:  "LVI",
			57:  "LVII",
			58:  "LVIII",
			59:  "LIX",
			60:  "LX",
			61:  "LXI",
			62:  "LXII",
			63:  "LXIII",
			64:  "LXIV",
			65:  "LXV",
			66:  "LXVI",
			67:  "LXVII",
			68:  "LXVIII",
			69:  "LXIX",
			70:  "LXX",
			71:  "LXXI",
			72:  "LXXII",
			73:  "LXXIII",
			74:  "LXXIV",
			75:  "LXXV",
			76:  "LXXVI",
			77:  "LXXVII",
			78:  "LXXVIII",
			79:  "LXXIX",
			80:  "LXXX",
			81:  "LXXXI",
			82:  "LXXXII",
			83:  "LXXXIII",
			84:  "LXXXIV",
			85:  "LXXXV",
			86:  "LXXXVI",
			87:  "LXXXVII",
			88:  "LXXXVIII",
			89:  "LXXXIX",
			90:  "XC",
			91:  "XCI",
			92:  "XCII",
			93:  "XCIII",
			94:  "XCIV",
			95:  "XCV",
			96:  "XCVI",
			97:  "XCVII",
			98:  "XCVIII",
			99:  "XCIX",
			100: "C",
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
