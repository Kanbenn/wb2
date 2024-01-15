package main

import (
	"errors"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
  - "a4bc2d5e" => "aaaabccddddde"
  - "abcd" => "abcd"
  - "45" => "" (некорректная строка)
  - "" => ""

Дополнительное задание: поддержка escape - последовательностей
  - qwe\4\5 => qwe45 (*)
  - qwe\45 => qwe44444 (*)
  - qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var ErrIncorrectString = errors.New("некорректная строка")

func UnzipStr(s string) (string, error) {
	var prevRune rune
	var sb strings.Builder
	var isEscaped bool

	backSlash := '\\'

	for i, r := range s {
		isDigit := unicode.IsDigit(r)
		isBackSlash := (r == backSlash)

		// если первый символ это число -- ошибка.
		// если два числа подряд и предыдущее число не эскейплено -- ошибка.
		firstSymbolIsDigit := (isDigit && (i == 0))
		isUnescapedPrevDigit := (unicode.IsDigit(prevRune) && !isEscaped)
		twoUnescapedDigitsInARow := (isDigit && isUnescapedPrevDigit)
		if firstSymbolIsDigit || twoUnescapedDigitsInARow {
			return s, ErrIncorrectString
		}

		{ // обработка слэшей
			if isBackSlash && prevRune != backSlash {
				prevRune = r
				continue
			}
			if isBackSlash && prevRune == backSlash { // если два слэйша подряд
				sb.WriteRune(r) // пишу только один слэш из двух
				prevRune = -1   // и обнуляю предыдущую руну чтобы следующая итерация игнорила слэш.
				continue
			}
		}

		if !isDigit || prevRune == backSlash {
			sb.WriteRune(r)
			if isDigit {
				prevRune = r
				isEscaped = true // эскейпим это число на следующую итерацию чтобы не сработала ошибка.
				continue
			}
		}
		// если это число, берём предыдущую руну и повторяем её столько раз.
		if isDigit {
			repeat := int(r - '1')
			// log.Println("repeat:", repeat)
			if repeat < 0 {
				return s, ErrIncorrectString
			}
			if prevRune == -1 { // если prevRune == -1 значит предыдущий символ бэк-слэш
				prevRune = backSlash // возвращаю бэкслэш, чтобы код ниже мог сделать распаковку слэша.
			}
			// str := strings.Repeat(string(prevRune), repeat)
			// sb.WriteString(str)
			for i := 0; i < repeat; i++ {
				sb.WriteRune(prevRune)
			}

		}
		prevRune = r
		isEscaped = false
	}

	return sb.String(), nil
}

// вариант решения без дополнительного задания по обработке слэшей
func UnzipStrNoBackSlashes(s string) (string, error) {
	var prevRune rune
	var sb strings.Builder

	for i, r := range s {
		isDigit := unicode.IsDigit(r)

		if (isDigit && i == 0) || (isDigit && unicode.IsDigit(prevRune)) {
			return s, ErrIncorrectString
		}
		if !isDigit {
			sb.WriteRune(r)
		}
		// если это число, берём предыдущую руну и повторяем её столько раз.
		if isDigit {
			repeat := int(r - '1')
			str := strings.Repeat(string(prevRune), repeat)
			sb.WriteString(str)
		}
		prevRune = r
	}
	return sb.String(), nil
}
