package trimmedmean

import "testing"

func TestTrimmed_meanSymmetric(t *testing.T) {
	data := []float64{1, 2, 3, 4, 100}
	got, err := TrimmedMean(data, 0.2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := 3.0
	if got != want {
		t.Errorf("expected %.2f, got %.2f", want, got)
	}
}

func TestTrimmed_meanAsymmetric(t *testing.T) {
	data := []float64{1, 2, 3, 4, 100}
	got, err := TrimmedMean(data, 0.2, 0.0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := 27.25
	if got != want {
		t.Errorf("expected %.2f, got %.2f", want, got)
	}
}

func TestTrimmed_meanEmpty(t *testing.T) {
	_, err := TrimmedMean([]float64{}, 0.1)
	if err == nil {
		t.Error("expected error for empty slice, got nil")
	}
}

func TestToFloat64(t *testing.T) {
	ints := []int{1, 2, 3}
	floats := ToFloat64(ints)
	for i, v := range ints {
		if floats[i] != float64(v) {
			t.Errorf("expected %v, got %v", float64(v), floats[i])
		}
	}
}
