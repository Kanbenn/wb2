Что выведет программа? Объяснить вывод программы.

```go
package main

import (
    "fmt"
)

func main() {
    a := [5]int{76, 77, 78, 79, 80}
    var b []int = a[1:4]
    fmt.Println(b)
}
```

Ответ:
```
[77, 78, 79]

```
Левая граница среза слайса является Включительной, а правая -- Исключающей. Это конечно криндж, но не только в голанге так, а почти во всех языках программирования.
