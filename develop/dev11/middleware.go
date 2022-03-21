package main

import (
	"log"
	"net/http"
	"time"
)

// Logging middleware логирует запросы
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// запускаем таймер замеряющий время обработки запроса
		start := time.Now()
		// передаем запрос следующему обработчику
		next.ServeHTTP(w, req)
		// после завершения обработки выводим данные о запросе в лог
		log.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
	})
}
