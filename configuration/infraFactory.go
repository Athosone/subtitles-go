package configuration

import (
	domain "subtitles/domain"
	repository "subtitles/infrastructure/repository"
)

func NewSubRepository() domain.SubtitleRepository {
	var repo = repository.NewSubRepositoryImp()
	return repo
}
