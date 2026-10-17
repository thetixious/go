package order_status

import "testing"

func TestCountStatuses(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  [6]int
	}{
		{"example", []string{"new", "paid", "shipped", "cancelled", "paid", "unknown", "new"}, [6]int{2, 2, 1, 1, 1, 5}},
		{"empty", nil, [6]int{}},
		{"single_new", []string{"new"}, [6]int{1, 0, 0, 0, 0, 1}},
		{"only_cancelled", []string{"cancelled", "cancelled"}, [6]int{0, 0, 0, 2, 0, 0}},
		{"unknown_and_case", []string{"PAID", "", "other", "paid"}, [6]int{0, 1, 0, 0, 3, 1}},
		{"all_active", []string{"new", "paid", "shipped"}, [6]int{1, 1, 1, 0, 0, 3}},
		{"empty_slice", []string{}, [6]int{}},
		{"only_paid", []string{"paid", "paid"}, [6]int{0, 2, 0, 0, 0, 2}},
		{"only_shipped", []string{"shipped", "shipped", "shipped"}, [6]int{0, 0, 3, 0, 0, 3}},
		{"case_sensitive", []string{"NEW", "Paid", "SHIPPED", "Cancelled"}, [6]int{0, 0, 0, 0, 4, 0}},
		{"whitespace_is_unknown", []string{" new", "paid ", "shipped\n", "\tcancelled"}, [6]int{0, 0, 0, 0, 4, 0}},
		{"distinct_counts", []string{"new", "paid", "paid", "shipped", "shipped", "shipped", "cancelled", "cancelled", "cancelled", "cancelled", "", "other", "unknown", "NEW", " paid"}, [6]int{1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, b, c, d, e, f := CountStatuses(tt.input)
			got := [6]int{a, b, c, d, e, f}
			if got != tt.want {
				t.Errorf("CountStatuses() = %v, want %v", got, tt.want)
			}
		})
	}
}
