package hull

import (
	"sort"
)

func ConvexHull(points Points) Points {
	sort.Sort(points);
	var lower Points;
	lowerLen := 0;
	for i := 0; i < len(points); i++ {
		lowerLen = len(lower);
		for lowerLen >= 2 && cross(lower[lowerLen - 2], lower[lowerLen - 1], points[i]) <= 0 {
			lower = lower[:lowerLen - 1]
			lowerLen--
		}
		lower = append(lower, points[i])
	}

	var upper Points;
	var upperLen int;
	for i := len(points) - 1; i >= 0; i-- {
		upperLen = len(upper)
		for upperLen >= 2 && cross(upper[upperLen - 2], upper[upperLen - 1], points[i]) <= 0 {
			upper = upper[:upperLen - 1]
			upperLen--
		}
		upper = append(upper, points[i])
	}
	lower = lower[:lowerLen - 1]
	upper = upper[:upperLen - 1]

	return append(lower, upper...)

	return points
}