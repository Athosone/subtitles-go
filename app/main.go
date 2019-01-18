package main

import (
	"fmt"
	domain "subtitles/domain"
)

func main() {
	fmt.Print("Hello")
	sub := domain.Subtitle{}

	sub.Save()
}
