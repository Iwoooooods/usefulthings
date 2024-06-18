package set

import (
	"testing"
)

func Test(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	if !s.Contains(1) {
		t.Errorf("Expected set to contain 1")
	}
	s.Remove(1)
	if s.Contains(1) {
		t.Errorf("Not expected set to contain 1")
	}
	s = NewSet(1, 2, 3, 2, 1)
	if s.Size() != 3 {
		t.Errorf("Expected set to have size 3")
	}
}
