package iteration

import "testing"

func TestRepeat(t *testing.T) {
	var n int
	n = 5
	repeated := Repeat("a", n)
	expected := "aaaaa"

	if expected != repeated {
		t.Errorf("expected %q got %q", expected, repeated)
	}

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}

}
