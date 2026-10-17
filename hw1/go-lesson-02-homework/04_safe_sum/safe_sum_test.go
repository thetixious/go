package safe_sum

import (
	"math"
	"reflect"
	"testing"
)

func TestSafeSum(t *testing.T) {
	tests := []struct {
		name        string
		input       []int64
		wantTotal   int64
		wantSkipped []int
	}{
		{"example", []int64{1000, 2500, math.MaxInt64 - 7, 100}, 3600, []int{2}},
		{"empty", nil, 0, nil},
		{"normal", []int64{2, -5, 10}, 7, nil},
		{"positive_overflow", []int64{math.MaxInt64, 1, -1}, math.MaxInt64 - 1, []int{1}},
		{"negative_overflow", []int64{math.MinInt64, -1, 1}, math.MinInt64 + 1, []int{1}},
		{"multiple_skips", []int64{math.MaxInt64, 1, 2, -10, 5}, math.MaxInt64 - 5, []int{1, 2}},
		{"zero", []int64{0, 0}, 0, nil},
		{"min_and_max", []int64{math.MinInt64, math.MaxInt64}, -1, nil},
		{"empty_slice", []int64{}, 0, nil},
		{"exact_max", []int64{math.MaxInt64 - 1, 1, 0}, math.MaxInt64, nil},
		{"exact_min", []int64{math.MinInt64 + 1, -1, 0}, math.MinInt64, nil},
		{"min_operand_overflow", []int64{-1, math.MinInt64, 2}, 1, []int{1}},
		{"min_operand_safe", []int64{1, math.MinInt64}, math.MinInt64 + 1, nil},
		{"max_operand_overflow", []int64{1, math.MaxInt64, -2}, -1, []int{1}},
		{"max_operand_safe", []int64{-1, math.MaxInt64}, math.MaxInt64 - 1, nil},
		{"both_bounds_and_nonadjacent_skips", []int64{math.MaxInt64, 1, -math.MaxInt64, math.MinInt64, -1, 10}, math.MinInt64 + 10, []int{1, 4}},
		{"negative_multiple_skips", []int64{math.MinInt64, -1, -2, 10, -5}, math.MinInt64 + 5, []int{1, 2}},
		{"negative_total", []int64{-2, -5, 1}, -6, nil},
		{"order_avoids_overflow", []int64{math.MaxInt64, -1, 1}, math.MaxInt64, nil},
		{"order_avoids_underflow", []int64{math.MinInt64, 1, -1}, math.MinInt64, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			total, skipped := SafeSum(tt.input)
			if total != tt.wantTotal {
				t.Errorf("total = %d, want %d", total, tt.wantTotal)
			}
			if len(skipped) == 0 && len(tt.wantSkipped) == 0 {
				return
			}
			if !reflect.DeepEqual(skipped, tt.wantSkipped) {
				t.Errorf("skipped = %v, want %v", skipped, tt.wantSkipped)
			}
		})
	}
}
