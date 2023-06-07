package shared

import (
	"strconv"
	"strings"
)

func ParseInt(s string) int {
	striped := strings.Trim(s, " ")
	res, err := strconv.ParseInt(striped, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(res)
}
