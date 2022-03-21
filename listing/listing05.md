Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Вывод: error

test() возвращает пустую ссылку на структуру, реализующую интерфейс. 
Данная пустая ссылка принимается в интерфейс error, что делает его ненулевым.

Для вывода "ok" функция test() должна возвращать нулевой интерфейс error

```
