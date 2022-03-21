Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:

Вывод: nil false

В данном случае содержимое объекта (ошибка) равна nil, но ссылка на этот объект не равна nil.
Так происходит т.к. интерфейс хранит и тип данных и ссылку.

