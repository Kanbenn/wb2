Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// close(ch) // здесь нужно закрыть канал 
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
0...9
fatal error: all goroutines are asleep - deadlock!
```
Ошибка из-за того, что range продолжает ждать сообщения от незакрытого канала. Идиоматический стандарт в голанге: кто пишет в канал, тот должен его закрыть.
