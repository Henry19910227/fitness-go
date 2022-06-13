package preload

type Preload struct {
	Field string
}

type Input struct {
	Preloads []*Preload
}
