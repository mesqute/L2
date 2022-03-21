package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Page struct {
	Url  string
	Data string
}

func main() {
	/*	nodes, err := goquery.ParseUrl("https://www.google.ru")
		if err != nil {
			fmt.Println(err)
			return
		}

		urls := nodes.Find("url")
		str := nodes.Html()

		i := 0
		fmt.Println(str, i, urls)*/

	err := getFile(`https://www.includehelp.com`)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getFile(ur string) error {
	// парсим полученный URL
	parsedUrl, err := url.Parse(ur)
	if err != nil {
		log.Println(err)
		return err
	}

	// выполняем запрос
	response, err := http.Get(ur)
	if err != nil {
		log.Println(err)
		return err
	}
	defer response.Body.Close()

	// записываем полученный ответ в файл

	// получаем текущую директорию
	currentDir, err := os.Getwd()
	// формируем строку для создания директории расположения файла
	downloadPath := currentDir + "\\download" + filepath.Dir(parsedUrl.Path)
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
	pth := filepath.Base(parsedUrl.Path)
	out, err := os.Create(pth)
	defer out.Close()
	_, err = io.Copy(out, response.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	// возвращаем активную директорию в изначальное положение
	err = os.Chdir(currentDir)
	if err != nil {
		return err
	}
	return nil
}
