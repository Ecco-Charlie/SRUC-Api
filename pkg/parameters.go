package pkg

import (
	"fmt"
	"strings"
)

func GetParameter(path *string, num int) string {
	p := strings.Split(*path, "/")
	fmt.Println(p)
	return p[num+1]
}
