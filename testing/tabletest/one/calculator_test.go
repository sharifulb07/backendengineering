package mathUtil

import "testing"

func TestMultiply(t *testing.T) {

	result := Multiply(10, 20)
	expected := 20

	if expected != result {
		t.Errorf("Multiply (a*b) expected %d but go %d", expected, result)
	}
}