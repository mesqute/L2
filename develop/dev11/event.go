package main

import (
	"strconv"
	"time"
)

type Event struct {
	Id          uint64
	Date        time.Time
	Description string
}

// ParseEvent парсит строку с параметрами и возвращает заполненную структуру Event
func ParseEvent(params map[string][]string) (*Event, error) {

	// парсим id
	if _, ok := params["id"]; !ok {
		err := BadRequest.New("Bad Request")
		err = AddErrorContext(err, "в параметрах пропущен id")
		return nil, err
	}
	id, err := strconv.ParseUint(params["id"][0], 10, 64)
	if err != nil {
		err = BadRequest.New(err.Error())
		err = AddErrorContext(err, "невалидный id")

		return nil, err
	}

	// парсим date
	if _, ok := params["date"]; !ok {
		err := BadRequest.New("Bad Request")
		err = AddErrorContext(err, "в параметрах пропущен date")
		return nil, err
	}
	date, err := time.Parse("2006-01-02", params["date"][0])
	if err != nil {
		err = BadRequest.New(err.Error())
		err = AddErrorContext(err, "невалидный date")
		return nil, err
	}

	// парсим description
	descriptions, ok := params["description"]
	var description string
	if ok {
		description = descriptions[0]
	}

	return &Event{Id: id, Date: date, Description: description}, nil
}
