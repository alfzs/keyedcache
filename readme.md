# Keyedcache

Простая утилита для генерации ключей кэша из произвольных частей. Объединяет значения в строку с разделителем `:` (двоеточие), что удобно для использования в in-memory и распределённых кэшах (например, Redis).

## Установка

```bash
go get github.com/alfzs/keyedcache
```

## Использование

```go
package main

import (
	"fmt"
	"github.com/alfzs/keyedcache"
)

func main() {
	key := keyedcache.Make("user", 42, "profile")
	fmt.Println(key) // Вывод: user:42:profile
}
```

## Функции

### `Make(parts ...any) string`

Создаёт ключ кэша путём объединения переданных частей в строку через `:`. Каждая часть конвертируется в строку с помощью `fmt.Sprint`.

#### Примеры

```go
keyedcache.Make("tenant", 123, "session", "active")
// => "tenant:123:session:active"

keyedcache.Make("user", nil, 99)
// => "user:<nil>:99"
```

## Применение

- Генерация ключей для Redis или Memcached
- Построение ключей для in-memory map-кэшей
- Формирование унифицированных идентификаторов с пространством имён

## Лицензия

MIT
