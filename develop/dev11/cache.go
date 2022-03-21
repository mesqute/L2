package main

import (
	"sync"
	"time"
)

var (
	mtxInstance sync.Mutex
	instance    *Cache = nil
)

type Cache struct {
	data map[uint64]map[time.Time]Event
	mtx  sync.RWMutex
}

// initCacheInstance инициализирует единственный экземпляр структуры Cache
func initCacheInstance() {
	mtxInstance.Lock()
	defer mtxInstance.Unlock()
	if instance == nil {
		cache := new(Cache)
		cache.data = make(map[uint64]map[time.Time]Event)
		instance = cache
	}
}

// GetCacheInstance возвращает указатель на структуру Cache
func GetCacheInstance() *Cache {
	if instance == nil {
		initCacheInstance()
	}
	return instance
}

// Get возвращает данные из кэша по id.
// Если не находит данные, то возвращает ошибку.
func (c *Cache) Get(id uint64) (map[time.Time]Event, error) {
	//проверяем, инициализирован ли кеш
	if c.data == nil {
		err := New("[Cache.Get] кеш не инициализирован")
		return nil, err
	}

	// используем блокировку для чтения (не блокирует чтение для остальных, но блокирует запись)
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	// считываем данные и их наличие
	val, ok := c.data[id]
	// если данных нет, то возвращаем ошибку
	if !ok {
		err := NotFound.New("[Cache.Get] data not found")
		err = AddErrorContext(err, "Объект с таким id не найден")
		return nil, err
	}
	// если данные есть, то возвращаем их
	return val, nil
}

func (c *Cache) Insert(data *Event) error {
	id := data.Id
	date := data.Date

	//проверяем, инициализирован ли кеш
	if c.data == nil {
		err := New("[Cache.Get] кеш не инициализирован")
		return err
	}

	c.mtx.Lock()
	defer c.mtx.Unlock()

	// проверка, существуют ли в кэше данные с таким же id.
	// если не существуют, то инициализируем
	if _, ok := c.data[id]; !ok {
		c.data[id] = make(map[time.Time]Event)
	}

	// записываем данные в кэш
	c.data[id][date] = *data

	return nil
}

func (c *Cache) Update(data *Event) error {
	id := data.Id
	date := data.Date

	//проверяем, инициализирован ли кеш
	if c.data == nil {
		err := New("[Cache.Get] кеш не инициализирован")
		return err
	}

	c.mtx.Lock()
	defer c.mtx.Unlock()

	// проверка, существуют ли в кэше данные с таким же id.
	// если не существуют, то возвращаем ошибку
	if _, ok := c.data[id]; !ok {
		err := NotFound.New("[Cache.Update] data not found")
		err = AddErrorContext(err, "Объект с таким id не найден")
		return err
	}
	if _, ok := c.data[id][date]; !ok {
		err := NotFound.New("[Cache.Update] data not found")
		err = AddErrorContext(err, "Объект с таким date не найден")
		return err
	}

	// обновляем данные в кэше
	c.data[id][date] = *data

	return nil

}

func (c *Cache) Delete(data *Event) error {
	id := data.Id
	date := data.Date
	// проверяем, инициализирован ли кеш
	if c.data == nil {
		err := New("[Cache.Get] кеш не инициализирован")
		return err
	}

	c.mtx.Lock()
	defer c.mtx.Unlock()

	// считываем данные и их наличие

	// если данных нет, то возвращаем ошибку
	if _, ok := c.data[id]; !ok {
		err := NotFound.New("[Cache.Delete] data not found")
		err = AddErrorContext(err, "Объект с таким id не найден")
		return err
	}
	if _, ok := c.data[id][date]; !ok {
		err := NotFound.New("[Cache.Delete] data not found")
		err = AddErrorContext(err, "Объект с таким date не найден")
		return err
	}

	delete(c.data[id], date)

	return nil
}
