package sum

import "testing"

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
