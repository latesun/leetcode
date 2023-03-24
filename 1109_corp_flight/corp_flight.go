package corpflight

func corpFlightBookings(bookings [][]int, n int) []int {
	answer := make([]int, n)
	for _, booking := range bookings {
		first, last, seats := booking[0], booking[1], booking[2]
		answer[first-1] += seats
		if last < n {
			answer[last] -= seats
		}
	}

	for i := 1; i < n; i++ {
		answer[i] += answer[i-1]
	}
	return answer
}
