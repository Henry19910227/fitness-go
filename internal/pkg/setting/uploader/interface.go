package uploader

type Setting interface {
	AllowExts() []string
	MaxSize() int
	FilePath() string
}
