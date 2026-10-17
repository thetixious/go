package order_retry

import "testing"

func TestProcessOrder(t *testing.T) {
	tests := []struct {
		name         string
		results      []bool
		limit        int
		wantSuccess  bool
		wantAttempts int
	}{
		{"example", []bool{false, false, true, false}, 3, true, 3},
		{"first_success", []bool{true, false}, 2, true, 1},
		{"no_success", []bool{false, false, false}, 3, false, 3},
		{"limit_shorter", []bool{false, false, true}, 2, false, 2},
		{"results_shorter", []bool{false, true}, 5, true, 2},
		{"results_exhausted", []bool{false, false}, 5, false, 2},
		{"zero_limit", []bool{true}, 0, false, 0},
		{"negative_limit", []bool{true}, -4, false, 0},
		{"empty", nil, 5, false, 0},
		{"success_at_limit", []bool{false, true, true}, 2, true, 2},
		{"empty_slice", []bool{}, 5, false, 0},
		{"one_allowed_success", []bool{true, false}, 1, true, 1},
		{"one_allowed_failure", []bool{false, true}, 1, false, 1},
		{"stop_at_first_of_multiple_successes", []bool{false, true, false, true}, 4, true, 2},
		{"empty_zero_limit", []bool{}, 0, false, 0},
		{"empty_negative_limit", nil, -1, false, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			success, attempts := ProcessOrder(tt.results, tt.limit)
			if success != tt.wantSuccess || attempts != tt.wantAttempts {
				t.Errorf("ProcessOrder() = (%t,%d), want (%t,%d)", success, attempts, tt.wantSuccess, tt.wantAttempts)
			}
		})
	}
}
