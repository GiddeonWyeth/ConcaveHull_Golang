package concaveHull

import (
	"sort"
)

func convexHull(points Points) Points {
	sort.Sort(points);

	pointsCopy := append(Points(nil), points...)
	result := append(getConvexHullHalf(pointsCopy), getConvexHullHalf(pointsCopy.reverse())...)

	return append(result, points[len(points) - 1])
}