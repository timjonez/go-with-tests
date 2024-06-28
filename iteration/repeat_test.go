package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Repeat 6 times", func(t *testing.T) {
		repeated := Repeat("a", 6)
		expected := "aaaaaa"
		assert(t, repeated, expected)
	})
	t.Run("Repeat 2 times", func(t *testing.T) {
		repeated := Repeat("a", 2)
		expected := "aa"
		assert(t, repeated, expected)
	})
}

func assert(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
