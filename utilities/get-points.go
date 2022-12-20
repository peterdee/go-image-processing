package utilities

func GetPoints(current, amount, total int) (int, int) {
	start, end := 0, total

	if current >= amount {
		start = current - amount
	}

	if current < total-amount {
		end = current + amount
	}

	return start, end
}
