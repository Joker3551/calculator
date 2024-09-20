package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

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

	ArabToRome := map[int]string{
		0:  "Вечная пустота",
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите математическое выражение:")
	math, _ := reader.ReadString('\n')
	math = strings.ReplaceAll(math, " ", "")
	math = strings.ReplaceAll(math, "	", "")

	var oper string
	var mathList []string
	var o1 int
	var o2 int
	var res int
	var rome = false
	var err1 any
	var err2 any

	for i, w := range string(math) {
		if (string(w) == "+" || string(w) == "-" || string(w) == "/" || string(w) == "*") && i != 0 && string(math[i-1]) >= "0" && (string(math[i+1]) >= "0" || string(math[i+1]) == "-") {
			oper = string(w)
			str := string(math[:i]) + " " + string(math[i+1:])
			mathList = strings.Fields(str)

			if mathList[0] <= "9" && mathList[1] <= "9" {
				o1, err1 = strconv.Atoi(mathList[0])
				o2, err2 = strconv.Atoi(mathList[1])
				if err1 != nil || err2 != nil {
					panic("Что-то не так")
				}

			} else if romeToArab[mathList[0]] != 0 && romeToArab[mathList[1]] != 0 {
				o1 = romeToArab[mathList[0]]
				o2 = romeToArab[mathList[1]]
				rome = true

			} else {
				panic("Неправильное выражение!")
			}
			break
		}

	}

	switch string(oper) {
	case "+":
		res = o1 + o2
	case "-":
		res = o1 - o2
	case "/":
		res = o1 / o2
	case "*":
		res = o1 * o2
	default:
		panic("Оператор не найден!")

	}

	if rome {
		if res < 0 {
			panic("В Римской системе нет отрицательных чисел!")
		} else if res > 10 {
			if o1 > o2 {
				fmt.Println(ArabToRome[o1] + ArabToRome[o2])
			} else {
				fmt.Println(ArabToRome[o2] + ArabToRome[o1])
			}
		} else {
			fmt.Println(ArabToRome[res])
		}
	} else {
		fmt.Println(res)
	}

}
