Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func WriteChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func FanIn(a, b <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		for {
			select {
			case v := <-a: // case v, okA := <-a
				// if okA { ch <- v }
				ch <- v
			case v := <-b: // case v,okB := <-b
				// if okB { ch <- v }
				ch <- v
			}
			// if !okA && !okB { // Завершаем горутину, если оба канала закрыты.
			// 	return
			// }
		}
	}()
	return ch
}

func main() {

	ch1 := WriteChan(1, 3, 5, 7)
	ch2 := WriteChan(2, 4 ,6, 8)
	
	fan := FanIn(ch1, ch2 )
	for v := range fan {
		fmt.Println(v)
	}
}
```

Ответ:
`1, 2, 3, 4, 5, 6, 7, 8, 0, 0, 0...`

`FanIn` не проверяет закрыты ли каналы, поэтому будут бесконечные нули. Чтение закрытого канала не вызывает ошибку, просто выдаёт дефолтные значения этоготипа. Возможное решение указал в комментариях выше.

