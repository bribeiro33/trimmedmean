package trimmedmean

import (
	"errors"
	"sort"
)

// Computes a trimmed mean from a slice of float64 numbers
// Parameters control how much to remove from the low and high ends
//
// Usage:
//
//	TrimmedMean(data, 0.1)         // symmetric 10% trim from both ends
//	TrimmedMean(data, 0.05, 0.2)   // 5% from low end, 20% from high end
//
// Returns an error if trimming removes all data or proportions are invalid
func TrimmedMean(data []float64, trimArgs ...float64) (float64, error) {
	n := len(data)
	if n == 0 {
		return 0, errors.New("input slice is empty")
	}

	var lowTrim, highTrim float64
	switch len(trimArgs) {
	case 0:
		lowTrim, highTrim = 0, 0
	case 1:
		lowTrim, highTrim = trimArgs[0], trimArgs[0]
	case 2:
		lowTrim, highTrim = trimArgs[0], trimArgs[1]
	default:
		return 0, errors.New("too many arguments: expected at most two trimming proportions")
	}

	if lowTrim < 0 || highTrim < 0 || lowTrim >= 1 || highTrim >= 1 {
		return 0, errors.New("trimming proportions must be between 0 and 1")
	}

	sort.Float64s(data)

	lowCut := int(float64(n) * lowTrim)
	highCut := int(float64(n) * (1 - highTrim))

	if lowCut >= highCut {
		return 0, errors.New("trimming removes all data")
	}

	trimmed := data[lowCut:highCut]

	sum := 0.0
	for _, v := range trimmed {
		sum += v
	}

	return sum / float64(len(trimmed)), nil
}

// converts an int slice to a float64 slice.
func ToFloat64(ints []int) []float64 {
	floats := make([]float64, len(ints))
	for i, v := range ints {
		floats[i] = float64(v)
	}
	return floats
}
