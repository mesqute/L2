package pattern

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

// State содержит методы состояний
type State interface {
	SetContext(ctx *StateContext)
	DoSomething()
}

// StateContext представляет собой контекст состояний
type StateContext struct {
	state State
}

// NewStateContext конструктор контекста состояний
func NewStateContext(initState State) *StateContext {
	return &StateContext{state: initState}
}

// ChangeState устанавливает новое активное состояние
func (c *StateContext) ChangeState(state State) {
	c.state = state
}

// DoSome делает дела соответствующие активному состоянию
func (c StateContext) DoSome() {
	c.state.DoSomething()
}

// StateA представляет собой пример состояния
type StateA struct {
	context *StateContext
}

// SetContext задает контекст состояния
func (s *StateA) SetContext(ctx *StateContext) {
	s.context = ctx
}

// DoSomething делает дела
func (s *StateA) DoSomething() {
	/*...*/
}

func StateClient() {
	initState := new(StateA)
	context := NewStateContext(initState)
	initState.SetContext(context)
	context.DoSome()
}

/*
Паттерн "стратегия" обычно применяют для следующих задач:
	1) У объекта есть множество влияющих на поведение состояний
	2) Структура содержит множество условных операторов зависящих от текущих значений полей структуры
	3) Необходимо построить машину состояний

Плюсы:
	+ Избавление от множества условных операторов машины состояний
	+ Упрощение кода контекста

Минусы:
	- Неоправданное усложнение кода, если состояний мало и они редко изменяются
*/
