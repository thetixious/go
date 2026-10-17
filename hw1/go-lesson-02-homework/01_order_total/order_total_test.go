package order_total

import "testing"

func TestCalculateOrder(t *testing.T) {
	tests := []struct {
		name                                  string
		prices                                []int64
		quantities                            []int
		discount                              int
		wantSubtotal, wantDiscount, wantTotal int64
	}{
		{"example", []int64{19990, 4990, 12990}, []int{2, 3, 1}, 10, 67940, 6794, 61146},
		{"no_discount", []int64{101, 200}, []int{1, 2}, 0, 501, 0, 501},
		{"full_discount", []int64{123}, []int{2}, 100, 246, 246, 0},
		{"round_down", []int64{101}, []int{1}, 10, 101, 10, 91},
		{"empty", []int64{}, []int{}, 5, 0, 0, 0},
		{"mismatched_lengths", []int64{100}, []int{}, 10, 0, 0, 0},
		{"negative_price", []int64{-1}, []int{1}, 10, 0, 0, 0},
		{"negative_quantity", []int64{100}, []int{-1}, 10, 0, 0, 0},
		{"negative_discount", []int64{100}, []int{1}, -1, 0, 0, 0},
		{"oversized_discount", []int64{100}, []int{1}, 101, 0, 0, 0},
		{"nil_order", nil, nil, 0, 0, 0, 0},
		{"more_quantities", []int64{100}, []int{1, 2}, 10, 0, 0, 0},
		{"zero_price_and_quantity", []int64{0, 999, 100}, []int{2, 0, 3}, 10, 300, 30, 270},
		{"late_negative_price", []int64{100, -1}, []int{2, 1}, 10, 0, 0, 0},
		{"late_negative_quantity", []int64{100, 200}, []int{2, -1}, 10, 0, 0, 0},
		{"negative_price_zero_quantity", []int64{100, -1}, []int{1, 0}, 10, 0, 0, 0},
		{"negative_quantity_zero_price", []int64{100, 0}, []int{1, -1}, 10, 0, 0, 0},
		{"discount_on_subtotal", []int64{5, 5}, []int{1, 1}, 10, 10, 1, 9},
		{"fraction_below_kopeck", []int64{1}, []int{1}, 99, 1, 0, 1},
		{"large_exact_amount", []int64{9007199254740993}, []int{1}, 1, 9007199254740993, 90071992547409, 8917127262193584},
		{"large_discount_precision", []int64{9007199254740972}, []int{1}, 18, 9007199254740972, 1621295865853374, 7385903388887598},
		{"more_prices_nonempty", []int64{100, 200}, []int{1}, 10, 0, 0, 0},
		{"both_negative", []int64{-100}, []int{-2}, 10, 0, 0, 0},
		{"invalid_price_full_discount", []int64{100, -1}, []int{1, 1}, 100, 0, 0, 0},
		{"invalid_quantity_no_discount", []int64{100, 200}, []int{1, -1}, 0, 0, 0, 0},
		{"discount_after_quantity", []int64{5}, []int{2}, 10, 10, 1, 9},
		{"whole_and_fractional_discount", []int64{199}, []int{1}, 50, 199, 99, 100},
		{"nil_prices_empty_quantities", nil, []int{}, 50, 0, 0, 0},
		{"empty_prices_nil_quantities", []int64{}, nil, 50, 0, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, b, c := CalculateOrder(tt.prices, tt.quantities, tt.discount)
			if a != tt.wantSubtotal || b != tt.wantDiscount || c != tt.wantTotal {
				t.Errorf("CalculateOrder() = (%d,%d,%d), want (%d,%d,%d)", a, b, c, tt.wantSubtotal, tt.wantDiscount, tt.wantTotal)
			}
		})
	}
}
