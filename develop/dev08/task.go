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

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	fmt.Println("Welcome to the Dev08")
	fmt.Println("Usage: <command> <args>")
	fmt.Println("requires bash or Git-bash/WSL to work")
	fmt.Println("Ctrl+C to exit")

	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Println("wb-shell$")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		args := strings.Fields(input)

		var cmd *exec.Cmd
		if len(args) <= 1 {
			cmd = exec.Command(input)
		} else {
			cmd = exec.Command(args[0], args[1:]...)
		}

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
}
