package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

type Data struct {
}

// DB это фабрика баз данных
type DB interface {
	Insert(data Data)
	Get(id int) Data
	Delete(id int)
}

// Postgres это один из продуктов фабрики
type Postgres struct {
	/*...*/
}

func (p *Postgres) Insert(data Data) {
	/*...*/
}

func (p *Postgres) Get(id int) Data {
	/*...*/
}

func (p *Postgres) Delete(id int) {
	/*...*/
}

func foo() {
	var db DB
	db = new(Postgres)

	// вызываем метод не конкретного объекта, а метод фабрики
	db.Get(5)
}

/*
Паттерн "фабричный метод" обычно применяют для следующих задач:
	1) Отделение кода производства объектов от кода использования объектов для более удобного расширения

Плюсы:
	+ Избавление от привязки к конкретной структуре
	+ Упрощение добавления новых продуктов
	+ Код добавления продукции расположен в одном месте

*/
