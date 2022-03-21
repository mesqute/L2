package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

// Request содержит поле, которое нужно обработать
type Request struct {
	data interface{}
}

// ChainHandler содержит методы обработчиков
type ChainHandler interface {
	canHandle(request *Request) bool
	SetNext(handler ChainHandler)
	Handle(request *Request)
}

type HandlerA struct {
	next ChainHandler
}

// canHandle проверяет запрос на возможность обработки
func (s *HandlerA) canHandle(request *Request) bool {
	/*...*/
}

// SetNext задает следующий в цепи обработчик
func (s *HandlerA) SetNext(handler ChainHandler) {
	s.next = handler
}

// Handle проверяет возможность обработки запроса и обрабатывает его.
// Если Handle не может обработать запрос, то передает запрос следующему обработчику.
func (s *HandlerA) Handle(request *Request) {
	if s.canHandle(request) {
		/*...*/
		return
	}
	if s.next != nil {
		s.next.Handle(request)
	}
}

type HandlerB struct {
	next ChainHandler
}

// canHandle проверяет запрос на возможность обработки
func (s *HandlerB) canHandle(request *Request) bool {
	/*...*/
}

// SetNext задает следующий в цепи обработчик
func (s *HandlerB) SetNext(handler ChainHandler) {
	s.next = handler
}

// Handle проверяет возможность обработки запроса и обрабатывает его.
// Если Handle не может обработать запрос, то передает запрос следующему обработчику.
func (s *HandlerB) Handle(request *Request) {
	if s.canHandle(request) {
		/*...*/
		return
	}
	if s.next != nil {
		s.next.Handle(request)
	}
}

/*
Паттерн "цепочка вызовов" обычно применяют для следующих задач:
	1) Обработка заранее неизвестных разнообразных запросов различными методами
	2) Необходимо строго последовательное выполнение обработчиков
	3) Требуется возможность динамически изменять цепочку обработчиков

Плюсы:
	+ Реализует принцип единственной ответственности
	+ Уменьшает зависимость между клиентом и обработчиками
Минусы:
	- Запрос может быть не обработан
*/
