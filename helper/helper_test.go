package helper

import "testing"

func TestHelper(t *testing.T) {
	{
		arr := []string{"A", "B", "C"}
		pos, got := InArrayStr(arr, "B")
		want := true
		wantPos := 1

		if got != want {
			t.Errorf("Want %t, got %t", want, got)
		}
		if pos != wantPos {
			t.Errorf("Want %d, got %d", wantPos, pos)
		}
	}
	{
		arr := []string{"A", "B", "C"}
		pos, got := InArrayStr(arr, "X")
		want := false
		wantPos := -1

		if got != want {
			t.Errorf("Want %t, got %t", want, got)
		}
		if pos != wantPos {
			t.Errorf("Want %d, got %d", wantPos, pos)
		}
	}
}
