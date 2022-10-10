package where

type Where struct {
	Query string
	Args []interface{}
}

type Input struct {
	Wheres []*Where
}
