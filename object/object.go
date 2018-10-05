package object

import "fmt"

// Object definition

const (
	INTEGER_OBJ = "INTEGER"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

// Enviroment

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

// Each objects

type Integer struct {
	Value int64
}
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
