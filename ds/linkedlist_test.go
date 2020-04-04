package ds_test

import (
	"errors"
	"testing"

	"github.com/urbanishimwe/challenges/ds"
)

func TestInsert(t *testing.T) {
	l := ds.LinkedList()
	l.Insert(1)
	if l.Get(1).Index != 1 {
		t.Errorf("expected %v, got %v", 1, l.Get(1).Index)
	}
}

func TestFailInsert(t *testing.T) {
	l := ds.LinkedList()
	_, got := l.Insert(func() int { return 1 })
	expected := errors.New("can not insert un-comparable value")
	if !errors.As(got, &expected) {
		t.Errorf("expected error %v, got %v", expected, got)
	}
}
