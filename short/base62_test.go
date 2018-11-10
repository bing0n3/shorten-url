package short

import (
	"testing"
)

func TestDecode(t *testing.T) {
	type test struct {
		input string
		want  int
	}
	test_here := make([]test, 100000)
	for i := 0; i < len(test_here); i++ {
		test_here[i] = test{input: Encode(i), want: i}
	}

	for _, a_test := range test_here {
		if got, _ := Decode(a_test.input); got != a_test.want {
			t.Errorf("Decode(%s)=%d, expect: %d", a_test.input, got, a_test.want)
		}
	}
}
