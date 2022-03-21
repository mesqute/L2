package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

// Strategy содержит общую для всех стратегий функцию выполнения
type Strategy interface {
	Execute(data Data)
}

type StrategyA struct {
}

// Execute выполняет алгоритм соответствующий стратегии А
func (s StrategyA) Execute(data Data) {
	/*...*/
}

// Context позволяет сохранять и использовать определенную стратегию
type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(s Strategy) {
	c.strategy = s
}

func (c *Context) ExecStrategy(data Data) {
	c.strategy.Execute(data)
}

func StrategyClient() {
	context := new(Context)
	str := new(StrategyA)
	context.SetStrategy(str)
	context.ExecStrategy(Data{})
}

/*
Паттерн "стратегия" обычно применяют для следующих задач:
	1) Нужно использовать разные виды одного алгоритма
	2) Приведение похожих структур в единую структуру
	3) Необходимо скрыть детали реализации алгоритмов
	4) Когда есть большое дерево условных операторов, где каждая ветвь представляет собой вариацию алгоритма

Плюсы:
	+ Легко менять используемые алгоритмы
	+ Изолирует код алгоритмов

Минусы:
	- Усложнение кода
	- Необходимость клиенту знать разницу между алгоритмами
*/
