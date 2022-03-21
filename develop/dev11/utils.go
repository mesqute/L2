package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

// CheckMethods проверяет запрос на соответствие обрабатываемым методам
func CheckMethods(r *http.Request, methods ...string) bool {
	for _, method := range methods {
		if r.Method == method {
			return true
		}
	}
	return false
}

// ErrorMethodNotAllowed отправляет код 405 с описанием доступных методов
func ErrorMethodNotAllowed(w http.ResponseWriter, allowMethods ...string) {

	// объединение в строку всех методов переданных в параметрах функции
	allowMethodsString := strings.Join(allowMethods, ", ")

	// передача в заголовок ответа список доступных методов
	w.Header().Set("Allow", allowMethodsString)

	// отправка ответа с кодом и описанием ошибки
	http.Error(w, "Метод запрещен", http.StatusMethodNotAllowed)
}

// Respond формирует и отправляет ответ в формате JSON
func Respond(w http.ResponseWriter, status int, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)

}

// Message формирует тело возвращаемого сообщения
func Message(status string, message string) map[string]interface{} {
	return map[string]interface{}{status: message}
}
