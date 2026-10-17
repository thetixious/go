package order_status

// CountStatuses возвращает количество new, paid, shipped, cancelled,
// unknown и active (new + paid + shipped) в указанном порядке.
// Любой статус, не равный одному из четырёх известных, считается unknown.
func CountStatuses(statuses []string) (newCount, paidCount, shippedCount, cancelledCount, unknownCount, activeCount int) {
	// TODO: используйте range и switch.
	return 0, 0, 0, 0, 0, 0
}
