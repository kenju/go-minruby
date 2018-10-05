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
		{"-5", -5},
		{"-10", -10},
		// +
		{"1 + 1", 2},
		{"5 + 5", 10},
		// -
		{"4 - 3", 1},
		{"3 - 4", -1},
		// *
		{"4 * 3", 12},
		{"3 * 4", 12},
		{"-4 * 3", -12},
		{"-3 * -4", 12},
		// /
		{"10 / 2", 5},
		{"10 / 3", 3},
		{"-10 / 2", -5},
		{"-10 / -3", 3},
		// %
		{"10 % 3", 1},
		{"10 % 2", 0},
		// mixture of operators
		{"1 + 2 + 3 % 4 * 5 * 6 + 7 - 8 / 9", 100},
		// with parenthesis
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestVariableStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{
			input: `
				x = 5
				x
			`,
			expected: 5,
		},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestEvalBooleanExpressoins(t *testing.T) {
	tests := []struct {
		input string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
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

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t", result.Value, expected)
		return false
	}

	return true
}
