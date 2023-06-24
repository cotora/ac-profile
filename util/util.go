package util

import(
	"strings"
	"unicode/utf8"
)

func RatingColor(rating int, s string) string {
	if rating >= 2800 {
		return "\033[38;5;196m" + s + "\033[m"
	} else if rating >= 2400 {
		return "\033[38;5;208m" + s + "\033[m"
	} else if rating >= 2000 {
		return "\033[38;5;226m" + s + "\033[m"
	} else if rating >= 1600 {
		return "\033[38;5;27m" + s + "\033[m"
	} else if rating >= 1200 {
		return "\033[38;5;123m" + s + "\033[m"
	} else if rating >= 800 {
		return "\033[38;5;76m" + s + "\033[m"
	} else if rating >= 400 {
		return "\033[38;5;94m" + s + "\033[m"
	} else {
		return "\033[38;5;248m" + s + "\033[m"
	}
}

func Max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func PaddingSpace(s string, n int) string {
	return s + strings.Repeat(" ", n-utf8.RuneCountInString(s))
}