package keyedcache

import (
	"fmt"
	"strings"
)

// Make создает ключ кэша из переданных частей, объединяя их через двоеточие.
// Каждая часть преобразуется в строку с помощью fmt.Sprint.
func Make(parts ...any) string {
	strs := make([]string, len(parts))
	for i, part := range parts {
		strs[i] = fmt.Sprint(part)
	}
	return strings.Join(strs, ":")
}
