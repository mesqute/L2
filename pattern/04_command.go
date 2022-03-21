package pattern

/*
	Реализовать паттерн «команда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

// Command содержит метод выполнения команды
type Command interface {
	Execute()
}

// SaveCommand выполняет функцию сохранения
type SaveCommand struct {
	// получатель и параметры команды
	/*...*/
}

// NewSaveCommand конструктор команды SaveCommand
func NewSaveCommand( /*получатель и параметры команды*/ ) *SaveCommand {
	return &SaveCommand{ /*получатель и параметры команды*/ }
}

// Execute выполняет команду (запускает процедуру сохранения)
func (s *SaveCommand) Execute() {
	/*...*/
}

// LoadCommand выполняет функцию загрузки
type LoadCommand struct {
	// получатель и параметры команды
	/*...*/
}

// NewLoadCommand конструктор команды LoadCommand
func NewLoadCommand( /*получатель и параметры команды*/ ) *LoadCommand {
	return &LoadCommand{ /*получатель и параметры команды*/ }
}

// Execute выполняет команду (запускает процедуру сохранения)
func (l *LoadCommand) Execute() {
	/*...*/
}

// Handler вызывает заданные команды
type Handler struct {
	command Command
}

// SetCommand задает заранее определенную команду
func (h *Handler) SetCommand(com Command) {
	h.command = com
}

// ExecCommand запускает выполнение команды
func (h *Handler) ExecCommand() {
	h.command.Execute()
}

/*
Паттерн "команда" обычно применяют для следующих задач:
	1) Требуется хранение, логирование, передача команд
	2) Хранение команд для операции undo
	3) Выполнение команд по расписанию, в порядке очереди, передавать команды на удаленный сервер и т.п.

Плюсы:
	+ Убирает прямую зависимость между вызывающими объектами и исполняющими объектами
	+ Позволяет легко реализовать undo и redo
	+ Добавляет возможность отложенного запуска
	+ Возможность сборки комплексных команд
Минусы:
	- Усложняет код
*/
