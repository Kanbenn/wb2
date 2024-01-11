package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type config struct {
	fields      string
	delimiter   string
	isSeparated bool
}

func parseFlags() config {
	cfg := config{}

	flag.StringVar(&cfg.fields, "f", "0", "fields")
	flag.StringVar(&cfg.delimiter, "d", "\t", "delimiter")
	flag.BoolVar(&cfg.isSeparated, "s", false, "separated")
	flag.Parse()

	return cfg
}

func cut(input string, cfg config) string {
	if cfg.isSeparated && !strings.Contains(input, cfg.delimiter) {
		log.Fatal("incorrect string:", input)
	}
	var sb strings.Builder

	splitted := strings.Split(input, cfg.delimiter)
	columns := strings.Split(cfg.fields, ",")
	for i := 0; i < len(columns); i++ {
		column, err := strconv.Atoi(columns[i])
		if err != nil {
			log.Fatal(err)
		}
		sb.WriteString(splitted[column])
	}
	return sb.String()
}

func main() {
	cfg := parseFlags()
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		text := sc.Text()
		fmt.Println("cutting: ", cut(text, cfg))
	}
}
