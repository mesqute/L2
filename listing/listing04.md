Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
Вывод: числа от 0 до 9 включительно и после этого дедлок

Писатель в отдельной горутине записывает в небуферизованный канал числа, 
а главная горутина читает данные в range цикле. Когда Писатель завершает свою работу, 
канал остается открыт и range цикл продолжает ждать новые данные, что приводит к дедлоку.

```