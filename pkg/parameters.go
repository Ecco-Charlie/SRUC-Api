package pkg

import (
	"strings"
)

func GetParameter(path *string, num int) string {
	p := strings.Split(*path, "/")
	return p[num+1]
}
