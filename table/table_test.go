package table

import (
	"fmt"
	"testing"
)

var table = []struct {
	x    int
	y    int
	want int
}{
	{2, 2, 5},
	{5, 3, 8},
	{8, 4, 12},
	{12, 5, 17},
}

// table test (Normal)
func TestSum(t *testing.T) {
	for _, row := range table {
		got := Sum(row.x, row.y)
		if got != row.want {
			t.Errorf("Failed: Got %d want %d", got, row.want)
		}
	}
}

// Table test (with subtest) - de maneira "concorrente", t.Run roda possivelmente de maneira simultânea
func TestSumSubTest(t *testing.T) {
	for _, row := range table {
		t.Run(fmt.Sprintf("Test %d+%d", row.x, row.y), func(t *testing.T) {
			got := Sum(row.x, row.y)
			if got != row.want {
				t.Errorf("Failed: got %d want %d", got, row.want)
			}
		})

	}
}
