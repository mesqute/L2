package main

import (
	"net/http"
)

func createHandler(w http.ResponseWriter, r *http.Request) {
	// проверяем соответствие методу
	if !CheckMethods(r, http.MethodPost) {
		err := MethodNotAllowed.New("method not allowed")
		HandleError(w, err)
		return
	}

	// проверяем соответствие Media Type
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/x-www-form-urlencoded" {
		err := UnsupportedMediaType.New(http.StatusText(http.StatusUnsupportedMediaType))
		err = AddErrorContext(err, http.StatusText(http.StatusUnsupportedMediaType))
		HandleError(w, err)
		return
	}

	// парсим строку параметров
	if err := r.ParseForm(); err != nil {
		HandleError(w, err)
		return
	}

	// проводим валидацию и заполняем структуру Event
	event, err := ParseEvent(r.PostForm)
	if err != nil {
		HandleError(w, err)
		return
	}

	// добавляем событие в БД
	err = InsertData(event)
	if err != nil {
		HandleError(w, err)
		return
	}

	// отправляем ответ об успешном завершении операции
	Respond(w, http.StatusCreated, Message("result", "created"))
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	// проверяем соответствие методу
	if !CheckMethods(r, http.MethodPost) {
		err := MethodNotAllowed.New("method not allowed")
		HandleError(w, err)
		return
	}

	// проверяем соответствие Media Type
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/x-www-form-urlencoded" {
		err := UnsupportedMediaType.New(http.StatusText(http.StatusUnsupportedMediaType))
		err = AddErrorContext(err, http.StatusText(http.StatusUnsupportedMediaType))
		HandleError(w, err)
		return
	}

	// считываем строку
	err := r.ParseForm()
	if err != nil {
		HandleError(w, err)
		return
	}

	// парсим строку параметров
	event, err := ParseEvent(r.PostForm)
	if err != nil {
		HandleError(w, err)
		return
	}

	// обновляем событие в БД
	err = UpdateData(event)
	if err != nil {
		HandleError(w, err)
		return
	}

	// отправляем ответ об успешном завершении операции
	Respond(w, http.StatusOK, Message("result", "updated"))

}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	// проверяем соответствие методу
	if !CheckMethods(r, http.MethodPost) {
		err := MethodNotAllowed.New("method not allowed")
		HandleError(w, err)
		return
	}

	// проверяем соответствие Media Type
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/x-www-form-urlencoded" {
		err := UnsupportedMediaType.New(http.StatusText(http.StatusUnsupportedMediaType))
		err = AddErrorContext(err, http.StatusText(http.StatusUnsupportedMediaType))
		HandleError(w, err)
		return
	}

	// парсим строку параметров
	err := r.ParseForm()
	if err != nil {
		HandleError(w, err)
		return
	}

	// проводим валидацию и заполняем структуру Event
	event, err := ParseEvent(r.PostForm)
	if err != nil {
		HandleError(w, err)
		return
	}

	// удаляем событие из БД
	err = DeleteData(event)
	if err != nil {
		HandleError(w, err)
		return
	}

	// отправляем ответ об успешном завершении операции
	Respond(w, http.StatusOK, Message("result", "deleted"))

}

func dayEventsHandler(w http.ResponseWriter, r *http.Request) {
	// проверяем соответствие методу
	if !CheckMethods(r, http.MethodGet) {
		err := MethodNotAllowed.New("method not allowed")
		HandleError(w, err)
		return
	}

	// парсим строку параметров
	err := r.ParseForm()
	if err != nil {
		HandleError(w, err)
		return
	}

	// проводим валидацию и заполняем структуру Event
	event, err := ParseEvent(r.Form)
	if err != nil {
		HandleError(w, err)
		return
	}

	// получаем список событий за один день с указанной даты
	data, err := GetData(event, 1)
	if err != nil {
		HandleError(w, err)
		return
	}

	// отправляем ответ
	Respond(w, http.StatusOK, Message("result", data))
}

func weekEventsHandler(w http.ResponseWriter, r *http.Request) {
	// проверяем соответствие методу
	if !CheckMethods(r, http.MethodGet) {
		err := MethodNotAllowed.New("method not allowed")
		HandleError(w, err)
		return
	}

	// парсим строку параметров
	err := r.ParseForm()
	if err != nil {
		HandleError(w, err)
		return
	}

	// проводим валидацию и заполняем структуру Event
	event, err := ParseEvent(r.Form)
	if err != nil {
		HandleError(w, err)
		return
	}

	// получаем список событий за семь дней (неделю) с указанной даты
	data, err := GetData(event, 7)
	if err != nil {
		HandleError(w, err)
		return
	}

	// отправляем ответ
	Respond(w, http.StatusOK, Message("result", data))

}

func monthEventsHandler(w http.ResponseWriter, r *http.Request) {
	// проверяем соответствие методу
	if !CheckMethods(r, http.MethodGet) {
		err := MethodNotAllowed.New("method not allowed")
		HandleError(w, err)
		return
	}

	// парсим строку параметров
	err := r.ParseForm()
	if err != nil {
		HandleError(w, err)
		return
	}

	// проводим валидацию и заполняем структуру Event
	event, err := ParseEvent(r.Form)
	if err != nil {
		HandleError(w, err)
		return
	}

	// получаем список событий за тридцать дней (месяц) с указанной даты
	data, err := GetData(event, 30)
	if err != nil {
		HandleError(w, err)
		return
	}

	// отправляем ответ
	Respond(w, http.StatusOK, Message("result", data))
}
