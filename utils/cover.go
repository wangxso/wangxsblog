package utils

import "strconv"

func PageInt(page string) int {
	n, _ := strconv.Atoi(page)
	return n
}

func SizeInt(size string) int {
	n, _ := strconv.Atoi(size)
	return n
}
