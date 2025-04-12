package config

import (
	"fmt"
	"io"
	"os"
)

func ShowBanner() {
	b, err := os.Open("resources/static/banner.txt")
	if err != nil {
		return
	}
	banner, err := io.ReadAll(b)
	if err != nil {
		return
	}
	fmt.Println(string(banner))
}
