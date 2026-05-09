package mathUtil

import "testing"

func TestSubtraction(t *testing.T) {

	

	tests:=[]struct{
		name string
		a int 
		b int 
		expected int 
	}{
		{"10-3", 10, 3, 7},
		{"20-4", 20, 4, 16},
		{"10-2", 10, 2, 8},
	}


	for _, tt:=range tests{

		t.Run(tt.name, func(t *testing.T) {
			result:=Subtraction(tt.a, tt.b)

			if result !=tt.expected{
				t.Errorf("got %d expected %d", result, tt.expected)
			}
		})
	}





}