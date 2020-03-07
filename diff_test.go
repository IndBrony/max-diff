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
	basicTest(t, []int{5, 5, 1}, []int{0, 1, 1, 2}, -1)
}
func TestMaxDiffString(t *testing.T) {
	basicTest := func(t *testing.T, input string, expected int) {
		if output := MaxDiffString(input); output != expected {
			t.Errorf(`
Test Failed with input %s
Expecting %v, but got %v
					`,
				input, expected, output,
			)
		}
	}
	basicTest(t, "1 2 3 4-1 2 3 4", 8)
	basicTest(t, "10 20 30-0 0 0", 60)
	basicTest(t, "5 5 1-0 1 2", 10)
	basicTest(t, "5 5 1-1 0 1 2", -1)
	basicTest(t, "5 5 1-a 1 2", -1)
}
