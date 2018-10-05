package object

// Environment

func NewEnvironment() *Environment {
	return &Environment{
		store: make(map[string]Object),
		outer: nil,
	}
}

type Environment struct {
	store map[string]Object
	outer *Environment
}
