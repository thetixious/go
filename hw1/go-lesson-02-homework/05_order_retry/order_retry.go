package order_retry

// ProcessOrder обрабатывает результаты попыток по порядку, но делает не
// более maxAttempts попыток. Успех — первое true.
// attempts — число реально просмотренных результатов.
func ProcessOrder(results []bool, maxAttempts int) (success bool, attempts int) {
	for _, value := range results {
		if attempts >= maxAttempts {
			break
		}

		attempts = attempts + 1

		if value {
			return true, attempts
		}
	}
	return false, attempts
}
