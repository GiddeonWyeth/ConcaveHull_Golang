package hull

import "math"

func sqLength(a, b Point) float64 {
	return math.Pow(b[0] - a[0], 2) + math.Pow(b[1] - a[1], 2);
}

func cos(o, a, b Point) float64 {
	aShifted := [2]float64{a[0] - o[0], a[1] - o[1]}
	bShifted := [2]float64{b[0] - o[0], b[1] - o[1]}
	sqALen := sqLength(o, a)
	sqBLen := sqLength(o, b)
	dot := aShifted[0] * bShifted[0] + aShifted[1] * bShifted[1];

	return dot / math.Sqrt(sqALen * sqBLen);
}

func _midPoint(edge, innerPoints, convex Points) Point {
	var point Point = nil
	angle1Cos := MAX_CONCAVE_ANGLE_COS
	angle2Cos := MAX_CONCAVE_ANGLE_COS

	for i := 0; i < len(innerPoints); i++ {
		a1Cos := cos(edge[0], edge[1], innerPoints[i]);
		a2Cos := cos(edge[1], edge[0], innerPoints[i]);

		if (a1Cos > angle1Cos && a2Cos > angle2Cos && !_intersect(Points{edge[0], innerPoints[i]}, convex) && !_intersect(Points{edge[1], innerPoints[i]}, convex)) {

			angle1Cos = a1Cos;
			angle2Cos = a2Cos;
			point = innerPoints[i];
		}
	}

	return point;
}

func _intersect(segment, pointset Points) bool {
	for i := 0; i < len(pointset) - 1; i++ {
		seg := Points{pointset[i], pointset[i + 1]};
		if (segment[0][0] == seg[0][0] && segment[0][1] == seg[0][1]) || (segment[0][0] == seg[1][0] && segment[0][1] == seg[1][1]) {
			continue;
		}
		if (intersect(segment, seg)) {
			return true;
		}
	}
	return false;
}

func occupiedArea(pointset Points) [2]float64 {
	minX := math.Inf(1);
	minY := math.Inf(1);
	maxX := math.Inf(-1);
	maxY := math.Inf(-1);

	for i := len(pointset - 1); i >= 0; i-- {
		if (pointset[i][0] < minX) {
			minX = pointset[i][0];
		}
		if (pointset[i][1] < minY) {
			minY = pointset[i][1];
		}
		if (pointset[i][0] > maxX) {
			maxX = pointset[i][0];
		}
		if (pointset[i][1] > maxY) {
			maxY = pointset[i][1];
		}
	}

	return [2]float64{maxX - minX, maxY - minY};
}

func bBoxAround(edge Points) [4]float64 {
	return [4]float64{math.Min(edge[0][0], edge[1][0]), math.Min(edge[0][1], edge[1][1]), math.Max(edge[0][0], edge[1][0]), math.Max(edge[0][1], edge[1][1])};
}

func cross(o, a, b Point) float64 {
	return (a[0] - o[0]) * (b[1] - o[1]) - (a[1] - o[1]) * (b[0] - o[0]);
}

//func ccw(x1, y1, x2, y2, x3, y3 float64) bool {
//var cw = ((y3 - y1) * (x2 - x1)) - ((y2 - y1) * (x3 - x1));
//return cw > 0 ? true : cw < 0 ? false : true;
//}

//Возможна ошибка!!!
func intersect(seg1, seg2 Points) bool {
	return ((cross(seg1[0], seg2[0], seg2[1]) >= 0) != (cross(seg1[1], seg2[0], seg2[1]) >= 0)) && ((cross(seg1[0], seg1[1], seg2[0]) >= 0) != (cross(seg1[0], seg1[1], seg2[1])))
	//return ccw(x1, y1, x3, y3, x4, y4) != ccw(x2, y2, x3, y3, x4, y4) && ccw(x1, y1, x2, y2, x3, y3) != = ccw(x1, y1, x2, y2, x4, y4);
}