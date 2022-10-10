package join

type Join struct {
	Query string
	Args []interface{}
}

type Input struct {
	Joins []*Join
}
