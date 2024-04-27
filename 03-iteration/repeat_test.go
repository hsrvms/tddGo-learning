package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("repeated: %q, expected: %q", repeated, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		Repeat("h")
	}
}

func ExampleRepeat() {
	result := Repeat("st")
	fmt.Println(result)
	// Output: ststststst
}
