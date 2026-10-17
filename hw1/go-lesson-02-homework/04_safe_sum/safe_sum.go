package safe_sum

import "math"

// SafeSum последовательно суммирует транзакции, пропуская те, добавление
// которых привело бы к переполнению int64. skipped — индексы пропущенных
// элементов исходного среза, в порядке появления.
func SafeSum(transactions []int64) (total int64, skipped []int) {
	// Min < total + v < Max
	// Min -v < total < Max - v
	for i, v := range transactions {
		if v > 0 && total > math.MaxInt64-v || v < 0 && total < math.MinInt64-v {
			skipped = append(skipped, i)
		} else {
			total += v
		}
	}
	return total, skipped
}
