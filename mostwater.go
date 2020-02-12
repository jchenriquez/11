package mostwater

import (
	"math"
)

// MaxArea will return the maxAre of a given measured container
func MaxArea(height []int) int {

	start := 0
	end := len(height) - 1

	area := (end - start) * int(math.Min(float64(height[start]), float64(height[end])))

	for end-start >= 1 {
		area = int(math.Max(float64(area), float64((end-start))*math.Min(float64(height[start]), float64(height[end]))))

		changed := false
		var nStart int
		var minIndex *int
		var nEnd int
		var incr int

		if height[start] <= height[end] {
			nStart = start + 1
			nEnd = end
			minIndex = &start
			incr = 1
		} else {
			nStart = end - 1
			minIndex = &end
			nEnd = start
			incr = -1
		}
		for i := nStart; i != nEnd; i += incr {
			if height[i] >= height[*minIndex] {
				*minIndex = i
				changed = true
				break
			}
		}

		if !changed {
			break
		}
	}

	return area
}
