package helpers

import (
	"time"
)

func getCurrentDate() time.Time { // получить текущее время
	return time.Now()
}

func ContainsInt(a []int, x int) bool { // проверить существует ли в массиве чисел нужное нам число
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func ContainsString(a []string, x string) bool { // проверить существует ли в массиве строк нужная нам строка
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
