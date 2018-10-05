package evaluator

import (
	"testing"

	"github.com/kenju/go-minruby/lexer"
	"github.com/kenju/go-minruby/object"
	"github.com/kenju/go-minruby/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		// integer literal
		{"5", 5},
		{"10", 10},
		// +
		{"1 + 1", 2},
		{"5 + 5", 10},
		// -
		{"4 - 3", 1},
		{"3 - 4", -1},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	return Eval(program, env)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
		return false
	}

	return true
}
