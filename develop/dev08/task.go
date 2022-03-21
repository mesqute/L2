package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
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

func execCommand(command string) (error, string, string) {
	// инициализируем переменные, передаваемые как STDOUT и STDERR
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	// создаем команду
	cmd := exec.Command("powershell", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	// запускаем выполнение команды и ждем ее завершения
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}

func main() {
	for {
		fmt.Print("Введите строку: ")
		// инициализируем переменную, хранящую полученную из STDIN строку
		var str string
		// инициализируем сканер STDIN
		scanner := bufio.NewScanner(os.Stdin)
		// ждем пока в STDIN не будет написана строка
		if scanner.Scan() {
			// считываем строку
			str = scanner.Text()
		}

		// выполняем команду
		err, out, errout := execCommand(str)
		if err != nil {
			log.Printf("error: %v\n", err)
		}
		// выводит STDOUT процесса команды
		fmt.Println("--- stdout ---")
		fmt.Println(out)
		// выводит STDERR процесса команды
		fmt.Println("--- stderr ---")
		fmt.Println(errout)
	}
}
