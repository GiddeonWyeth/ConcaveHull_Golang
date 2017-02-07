package hull

import "math"

var MAX_SEARCH_BBOX_SIZE_PERCENT = 0.6
var MAX_CONCAVE_ANGLE_COS = math.Cos(90 / (180 / math.Pi));

func concave(convex Points, maxSqEdgeLen float64, maxSearchArea [2]float64, grid Grid, edgeSkipList []bool) Points {
	var midPointInserted = false;
	var midPoint Point = nil;
	var bBoxWidth float64;
	var bBoxHeight float64;

	for i := 0; i < len(convex) - 1; i++ {
		edge := Points{convex[i], convex[i + 1]};
		keyInSkipList := string(edge[0][0]) + ',' + string(edge[0][1]) + ',' + string(edge[1][0]) + ',' + string(edge[1][1]);

		if (sqLength(edge[0], edge[1]) < maxSqEdgeLen) || (edgeSkipList[keyInSkipList] == true) {
			continue;
		}

		scaleFactor := 0;
		bBoxAround := bBoxAround(edge);
		for midPoint == nil && (maxSearchArea[0] > bBoxWidth || maxSearchArea[1] > bBoxHeight) {
			bBoxAround = grid.extendBbox(bBoxAround, scaleFactor);
			bBoxWidth = bBoxAround[2] - bBoxAround[0];
			bBoxHeight = bBoxAround[3] - bBoxAround[1];

			midPoint = _midPoint(edge, grid.rangePoints(bBoxAround), convex);
			scaleFactor++;
		}

		if (bBoxWidth >= maxSearchArea[0] && bBoxHeight >= maxSearchArea[1]) {
			edgeSkipList[keyInSkipList] = true;
		}

		if (midPoint != nil) {
			convex.splice(i + 1, 0, midPoint);
			grid.removePoint(midPoint);
			midPointInserted = true;
		}
	}

	if (midPointInserted) {
		return concave(convex, maxSqEdgeLen, maxSearchArea, grid, edgeSkipList);
	}

	return convex;
}

func hull(pointset Points, concavity int) {

	occupiedArea := occupiedArea(pointset);
	maxSearchArea := [2]float64{occupiedArea[0] * MAX_SEARCH_BBOX_SIZE_PERCENT, occupiedArea[1] * MAX_SEARCH_BBOX_SIZE_PERCENT};

	convex := ConvexHull(pointset);
	innerPoints := Filter(pointset)

	cellSize := math.Ceil(1 / (float64(len(pointset)) / (occupiedArea[0] * occupiedArea[1])));

	concave := concave(convex, math.Pow(float64(concavity), 2), maxSearchArea, Grid{innerPoints, cellSize}, []bool{});

	return concave;
}