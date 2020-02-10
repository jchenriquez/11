package mostwater

import (
	"fmt"
	"math"
)

func toKey(i, j int) string {
	return fmt.Sprintf("%d,%d", i, j)
}

func maxAreaHelper(height []int, memo map[string]int, i, j int) int {
	area := (j - i) * int(math.Min(float64(height[i]), float64(height[j])))
	if j-i == 1 {
		return area
	}

	key := toKey(i, j)

	if val, saw := memo[key]; saw {
		return val
	}

	var area1 int
	var area2 int

	if float64(height[j-1]) > math.Min(float64(height[i]), float64(height[j])) {
		area1 = maxAreaHelper(height, memo, i, j-1)
	}

	if float64(height[i+1]) > math.Min(float64(height[i]), float64(height[j])) {
		area2 = maxAreaHelper(height, memo, i+1, j)
	}

	maxArea := int(math.Max(float64(area1), float64(area2)))
	return int(math.Max(float64(maxArea), float64(area)))

}

func areaHelper(height []int, start, end, incr int) int {
	area := height[start] * height[start+incr]

	if height[start] > height[start+incr] {
		maxPillar = start
	} else {
		maxPillar = start + incr
	}

	for i := start + (incr * 2); i != end; i += incr {
		if (i-maxPillar)*int(math.Min(float64(height[i]), float64(height[maxPillar]))) > area {
			area = (i - maxPillar) * int(math.Min(float64(height[i]), float64(height[maxPillar])))
		}

		if height[i] > height[maxPillar] {
			maxPillar = i
		}
	}

	return area
}

// MaxArea will return the maxAre of a given measured container
func MaxArea(height []int) int {

	var maxPillar int

	area := areaHelper(height, 0, len(height), 1)
	area = int(math.Max(float64()))
}
