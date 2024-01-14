package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===

# Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

const intro = `Welcome to the dev08 - a bash emulator
Usage: <command> <args>
requires bash or Git-bash/WSL to work
Ctrl+C to exit`

func main() {
	fmt.Println(intro)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		cmd := exec.Command("bash", "-c", input)

		cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
		cmd.Run()
	}
}
