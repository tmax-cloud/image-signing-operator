package utils

import "strings"

func ParseImage(image string) (name, tag string) {
	res := strings.Split(image, ":")

	if len(res) == 0 {
		return
	} else if len(res) == 1 {
		name = res[0]
		return
	}
	name = res[0]
	tag = res[1]

	return
}
