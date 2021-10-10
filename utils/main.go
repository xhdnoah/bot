package utils

import "unicode"

func IsChinese(s string) (zh bool) {
	for _, r := range s {
		if unicode.Is(unicode.Scripts["Han"], r) {
			zh = true
			break
		}
	}
	return
}
