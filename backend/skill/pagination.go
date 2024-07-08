package skill

import (
	"strconv"
)

func page(p string) uint {
	num, _ := strconv.Atoi(p)
	if num < 1 {
		return 1
	}
	return uint(num)
}

func itemsPerPage(p string) uint {
	switch num, _ := strconv.Atoi(p); {
	case num <= 0:
		return 10
	case num > 100:
		return 100
	default:
		return uint(num)
	}
}

func orderBy(o string) string {
	if o == "" {
		return "key"
	}

	return o
}
