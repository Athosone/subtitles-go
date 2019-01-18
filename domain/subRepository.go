package domain

type SubtitleRepository interface {
	Find(episode Episode) (Subtitle, error)
}
