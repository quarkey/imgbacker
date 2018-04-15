package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := new(Stack)

	s.Push("string1")
	s.Push("string2")

	expected := 2

	if s.Len() != expected {
		t.Errorf("expected %d items in stack\n", expected)
	}
	expected = 1
	s.Pop()
	if s.Len() != expected {
		t.Errorf("expected %d items in stack\n", expected)
	}

	expectedstr := "string1"
	if s.Pop().(string) != expectedstr {
		t.Errorf("expected %s, got %s", s.Pop().(string), expectedstr)
	}
}
