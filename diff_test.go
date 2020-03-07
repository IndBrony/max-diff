package diff

import "testing"

func TestMaxDiff(t *testing.T) {
	basicTest := func(t *testing.T, first, second []int, expected int) {
		if output := MaxDiff(first, second); output != expected {
			t.Errorf(`
Test Failed with input:
	first = %v
	second = %v
Expecting %v, but got %v
					`,
				first, second, expected, output,
			)
		}
	}
	basicTest(t, []int{1, 2, 3, 4}, []int{1, 2, 3, 4}, 8)
	basicTest(t, []int{10, 20, 30}, []int{0, 0, 0}, 60)
	basicTest(t, []int{5, 5, 1}, []int{0, 1, 2}, 10)
}
