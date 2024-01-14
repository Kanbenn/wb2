package main

import (
	"fmt"
	"slices"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func SetsOfAnagrams(in []string) *map[string][]string {
	out := make(map[string][]string)

	tempHashes := makeMapOfSortedHashesUsingRunes(in)
	// tempHashes := makeMapOfSortedHashesUsingStrings(in)

	for _, anagrams := range tempHashes {
		if len(anagrams) > 1 {
			firstWord := anagrams[0]
			out[firstWord] = append([]string{}, anagrams...)
		}
	}
	return &out
}

func makeMapOfSortedHashesUsingRunes(in []string) map[string][]string {
	out := make(map[string][]string)
	for i := range in {
		str := strings.ToLower(in[i])
		runes := []rune(str)
		slices.Sort(runes)
		hash := string(runes)
		out[hash] = append(out[hash], in[i])
	}
	return out
}

func makeMapOfSortedHashesUsingStrings(in []string) map[string][]string {
	out := make(map[string][]string)
	for i := range in {
		str := strings.ToLower(in[i])
		letters := strings.Split(str, "")
		slices.Sort(letters)
		hash := strings.Join(letters, "")
		out[hash] = append(out[hash], in[i])
	}
	return out
}

func main() {
	in := []string{"Пятак", "пЯтка", "тяпкА", "Листок", "слИток", "стОлик", "одувАн"}

	fmt.Println(SetsOfAnagrams(in))
}
