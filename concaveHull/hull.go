package concaveHull

import "math"

var MAX_SEARCH_BBOX_SIZE_PERCENT = 0.6
var MAX_CONCAVE_ANGLE_COS = math.Cos(90 / (180 / math.Pi));

func concave(convex Points, maxSqEdgeLen float64, maxSearchArea Point, grid Grid, edgeSkipList map[string]bool) Points {
	var midPointInserted = false;
	var midPoint Point = Point{};
	var bBoxWidth float64;
	var bBoxHeight float64;
	var i int;

	for i = 0; i < len(convex) - 1; i++ {
		edge := Points{convex[i], convex[i + 1]};
		keyInSkipList := floatToString(edge[0][0]) + "," + floatToString(edge[0][1]) + "," + floatToString(edge[1][0]) + "," + floatToString(edge[1][1]);

		if (sqLength(edge[0], edge[1]) < maxSqEdgeLen) || (edgeSkipList[keyInSkipList] == true) {
			continue;
		}

		scaleFactor := 0;
		midPoint, bBoxWidth, bBoxHeight, scaleFactor = getMidPoint(edge, convex, grid, scaleFactor);
		for midPoint == (Point{}) && (maxSearchArea[0] > bBoxWidth || maxSearchArea[1] > bBoxHeight) {
			midPoint, bBoxWidth, bBoxHeight, scaleFactor = getMidPoint(edge, convex, grid, scaleFactor);
		}

		if (bBoxWidth >= maxSearchArea[0] && bBoxHeight >= maxSearchArea[1]) {
			edgeSkipList[keyInSkipList] = true;
		}

		if (midPoint != (Point{})) {
			convex = splice(convex, i + 1, 0, midPoint);
			grid = grid.removePoint(midPoint);
			midPointInserted = true;
		}
	}

	if (midPointInserted == true) {
		return concave(convex, maxSqEdgeLen, maxSearchArea, grid, edgeSkipList);
	}

	return convex;
}

func Hull(pointset Points, concavity int) Points {

	occupiedArea := occupiedArea(pointset);
	maxSearchArea := Point{occupiedArea[0] * MAX_SEARCH_BBOX_SIZE_PERCENT, occupiedArea[1] * MAX_SEARCH_BBOX_SIZE_PERCENT};

	convex := convexHull(pointset);
	innerPoints := Filter(pointset, convex).reverse()

	cellSize := math.Ceil(1 / (float64(len(pointset)) / (occupiedArea[0] * occupiedArea[1])));
	concave := concave(convex, math.Pow(float64(concavity), 2), maxSearchArea, createGrid(innerPoints, cellSize), map[string]bool{});

	return concave;
}