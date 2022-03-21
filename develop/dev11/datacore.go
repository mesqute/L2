package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

var cache *Cache

func init() {
	// инициализируем кеш
	cache = GetCacheInstance()
}

// InsertData добавляет данные в память
func InsertData(data *Event) error {
	// добавляем данные в кеш
	if err := cache.Insert(data); err != nil {
		return err
	}
	return nil
}

// UpdateData обновляет данные в памяти
func UpdateData(data *Event) error {
	// обновляем данные в кеше
	if err := cache.Update(data); err != nil {
		return err
	}
	return nil

}

// DeleteData удаляет данные из кеша
func DeleteData(data *Event) error {
	// удаляем данные из кеша
	if err := cache.Delete(data); err != nil {
		return err
	}
	return nil
}

// GetData возвращает список событий за указанный период
func GetData(data *Event, days int) (string, error) {
	id := data.Id
	date := data.Date

	// получаем мапу событий
	events, err := cache.Get(id)
	if err != nil {
		return "", err
	}

	// отбираем нужные нам события
	var targetEvents []Event
	for i := 0; i < days; i++ {
		if event, ok := events[date.Add(time.Hour*24*time.Duration(i))]; ok {
			targetEvents = append(targetEvents, event)
		}
	}
	// если события не найдены, возвращаем ошибку
	if len(targetEvents) == 0 {
		err := NotFound.New("[GetData] no data")
		err = AddErrorContext(err, "Совпадений не найдено")
		return "", err
	}

	// формируем возвращаемый список событий
	var result []string
	for _, event := range targetEvents {
		result = append(result, fmt.Sprintf("%v: %v", event.Date.Format("2006-01-02"), event.Description))
	}

	// сортируем список событий
	sort.Strings(result)

	return strings.Join(result, ",\n"), nil
}
