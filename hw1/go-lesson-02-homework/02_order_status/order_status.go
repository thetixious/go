package order_status

// CountStatuses возвращает количество new, paid, shipped, cancelled,
// unknown и active (new + paid + shipped) в указанном порядке.
// Любой статус, не равный одному из четырёх известных, считается unknown.
func CountStatuses(statuses []string) (newCount, paidCount, shippedCount, cancelledCount, unknownCount, activeCount int) {
	for _, v := range statuses {
		switch v {
		case "new":
			newCount++
		case "paid":
			paidCount++
		case "shipped":
			shippedCount++
		case "cancelled":
			cancelledCount++
		default:
			unknownCount++
		}
	}
	return newCount, paidCount, shippedCount, cancelledCount, unknownCount, newCount + paidCount + shippedCount
}
