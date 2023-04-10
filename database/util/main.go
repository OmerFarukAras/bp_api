package util

import (
	"strings"
)

func CreateUID() string {
	data := "sadasd"
	println(data)
	return data
}

func CheckString(str string) bool {
	if len(strings.TrimSpace(str)) == 0 {
		return false
	}
	return true
}
