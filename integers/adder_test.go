package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(3, 4)
	expected := 7
	assertCorrectMessage(t, sum, expected)
}

func ExampleAdd() {
	// This test will be compiled & not executed if output comment is removed
	// no space btw ouput & : , else won't execute
	sum := Add(7, 9)
	fmt.Println(sum)
	// Output: 16
}

func assertCorrectMessage(t testing.TB, sum, expected int) {
	t.Helper()
	if sum != expected {
		t.Errorf("Got %d but expected %d", sum, expected)
	}
}
