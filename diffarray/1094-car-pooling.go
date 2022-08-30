package main

func carPooling(trips [][]int, capacity int) bool {
	MAX_TRIP := 1001
	res := make([]int, MAX_TRIP)

	for i := 0; i < len(trips); i++ {
		res[trips[i][1]] += trips[i][0]
		if trips[i][2]+1 < MAX_TRIP {
			res[trips[i][2]+1] -= trips[i][0]
		}
	}

	for i := 1; i < MAX_TRIP; i++ {
		res[i] += res[i-1]
	}
	for i := 0; i < MAX_TRIP; i++ {
		if res[i] > capacity {
			return false
		}
	}
	return true
}
