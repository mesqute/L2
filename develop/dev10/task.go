package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {

	var timeout time.Duration

	// считываем ключи
	flag.DurationVar(&timeout, "timeout", 10*time.Second, "таймаут на подключение к серверу")

	flag.Parse()

	// получаем данные для подключения
	host := flag.Arg(0)
	port := flag.Arg(1)

	// формируем строку для подключения
	connString := fmt.Sprintf("%s:%s", host, port)

	// инициализируем переменную для хранения подключения
	var conn net.Conn
	var err error

	// запускаем таймер таймаута
	timer := time.After(timeout)

	// запускаем цикл подключения
loop:
	for {
		select {
		// при истечении времени завершаем программу вызовом паники
		case <-timer:
			panic("Connection timeout!")
		default:
			// пытаемся подключиться
			conn, err = net.Dial("tcp", connString)
			if err != nil {
				// если не удается, то ждем 100 мс и пробуем снова
				time.Sleep(100 * time.Millisecond)
				continue
			}
			// если получается, выходим из цикла
			break loop
		}
	}

	// инициализируем канал для перехвата сообщения об закрытии приложения
	killSignalChan := make(chan os.Signal, 1)
	// перенаправляем сигнал в созданный канал
	signal.Notify(killSignalChan, os.Interrupt)

	// запускаем в отдельной горутине наблюдателя
	go func() {
		select {
		// если прилетает сигнал завершения работы, закрываем сокет
		case <-killSignalChan:
			conn.Close()
		}
	}()

	// бесконечный цикл отправки-получения данных
	for {
		// читаем данные из STDIN
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Сообщение: ")
		text, _ := reader.ReadString('\n')
		// отправляем в socket
		fmt.Fprintf(conn, text+"\n")
		// прослушиваем ответ
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Ответ: " + message)
	}
}
