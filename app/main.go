package main

import (
	"fmt"
	configuration "subtitles/configuration"
	domain "subtitles/domain"
)

func main() {
	fmt.Print("Hello")
	subtitleRepo := configuration.NewSubRepository()
	episode := domain.Episode{}
	episode.Number = 1
	episode.Season = 1
	episode.ShowName = "Origins"
	subtitleRepo.Find(episode)

	sub := domain.Subtitle{}

	sub.Save()
}
