package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

// House собираемая структура
type House struct {
	/*...*/
}

// HouseBuilder строитель структуры House
type HouseBuilder struct {
	result House
}

// Reset инициализирует новый экземпляр структуры House
func (h *HouseBuilder) Reset() {
	h.result = House{}
}

// SetWalls добавляет указанное кол-во стен
func (h *HouseBuilder) SetWalls(num int) {
	/*...*/
}

// SetRoof добавляет крышу
func (h *HouseBuilder) SetRoof() {
	/*...*/
}

// SetDoors добавляет указанное кол-во дверей
func (h *HouseBuilder) SetDoors(num int) {
	/*...*/
}

// SetWindows добавляет указанное кол-во окон
func (h *HouseBuilder) SetWindows(num int) {
	/*...*/
}

// GetResult возвращает собранную структуру Home
func (h *HouseBuilder) GetResult() House {
	return h.result
}

/*
Паттерн "строитель" обычно применяют для следующих задач:
	1) Пошаговая реализация сложной структуры,
	2) Избавление от конструктора с большим кол-вом параметров,
	3) Компоновка методов инициализации структуры в одном месте,
	4) Создание разных представлений одной структуры

Плюсы:
	+ Пошаговая сборка
	+ Переиспользуемость кода
	+ Изоляция кода сборки структуры от остальной бизнес-логики

Минусы:
	- Усложнение кода
*/
