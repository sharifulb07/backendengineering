package mathUtils

import "testing"

func TestAdd(t *testing.T) {
	result:=Add(40,50)
	expected:=90


	if expected!=result{
		t.Errorf("Expected %d but got %d", expected, result)
	}
}