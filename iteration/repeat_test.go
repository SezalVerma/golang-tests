package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 5)
	want := "aaaaa"
	assertCorrectMessage(t, got, want)
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got : %q , want : %q", got, want)
	}
}

// like test, appears in docs
func ExampleRepeat() {
	got := Repeat("b", 6)
	fmt.Println(got)
	// Output: bbbbbb
}

// to run it -> go test -bench="."
// benchmarks - like tests ; measures how long it takes to execute
func BenchmarkRepeat(b *testing.B) {
	// b.N times - decided by framework , to get decent results
	for i := 0; i < b.N; i++ {
		Repeat("a", 7)
	}
}
