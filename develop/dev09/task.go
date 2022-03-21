package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"regexp"
	"time"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

const (
	NotWorking = StateType(iota)
	Worked
	Saved
	NotVerified
)

type StateType uint8

func main() {

	flag.Parse()

	baseUrl := flag.Arg(0)

	parsedUrl, err := url.Parse(baseUrl)
	if err != nil {
		log.Fatal(parsedUrl)
		return
	}

	// создаем мапу, где ключи - url страниц, а значения - статус
	pages := make(map[url.URL]StateType)
	pages[*parsedUrl] = NotVerified

	// инициализируем канал для перехвата сообщения об закрытии приложения
	killSignalChan := make(chan os.Signal, 1)
	// перенаправляем сигнал в созданный канал
	signal.Notify(killSignalChan, os.Interrupt)

	log.Println("Начинаю работу")
loop:
	for {
		select {
		case <-killSignalChan:
			break loop
		default:
			// считываем кол-во сохраненных страниц до цикла обхода
			countPages := len(pages)
			log.Printf("Сохранено %v страниц\n", countPages)

			for pageUrl, stateType := range pages {
				var page []byte

				if stateType == NotVerified {
					pages[pageUrl], page = getPage(pageUrl)
				}
				if stateType == Worked {
					searchUrls(pages, page, pageUrl.Scheme, parsedUrl.Host)
					err := saveToFile(pageUrl, page)
					if err != nil {
						log.Fatal(err)
					}
				}
			}

			if countPages == len(pages) {
				break loop
			}
		}
	}
	log.Printf("Итого сохранено %v страниц\n", len(pages))

}

func getPage(ur url.URL) (StateType, []byte) {

	resp, err := http.Get(ur.String())
	time.Sleep(100 * time.Millisecond)
	if err != nil || resp.StatusCode != 200 || resp.ContentLength < 10 {
		// если ссылка не рабочая, возвращаем соответствующий статус
		return NotWorking, nil
	}

	var page []byte
	resp.Body.Read(page)
	defer resp.Body.Close()

	return Worked, page
}

func searchUrls(pages map[url.URL]StateType, page []byte, scheme, host string) {
	re := regexp.MustCompile(`href="(.*?)"`)
	hrefPage := re.FindAllSubmatch(page, -1)

	for _, p := range hrefPage {
		var newUrl url.URL
		newUrl.Path = string(p[1])
		newUrl.Scheme = scheme
		newUrl.Host = host

		if _, ok := pages[newUrl]; !ok {
			pages[newUrl] = NotVerified
		}
	}
}

func saveToFile(ur url.URL, body []byte) error {
	// получаем текущую директорию
	currentDir, err := os.Getwd()

	// формируем строку для создания директории расположения файла
	downloadPath := currentDir + "\\download" + filepath.Dir(ur.Path)
	// создаем директорию
	err = os.MkdirAll(downloadPath, os.ModePerm)
	if err != nil {
		return err
	}

	// устанавливаем созданную директорию как текущую
	err = os.Chdir(downloadPath)
	if err != nil {
		return err
	}

	// создаем файл и записываем в него полученные данные
	pth := filepath.Base(ur.Path)
	if len(pth) < 3 {
		pth = "file.html"
	}

	out, err := os.Create(pth)
	defer out.Close()
	_, err = out.Write(body)
	if err != nil {
		return err
	}
	// возвращаем активную директорию в изначальное положение
	err = os.Chdir(currentDir)
	if err != nil {
		return err
	}

	return nil
}
