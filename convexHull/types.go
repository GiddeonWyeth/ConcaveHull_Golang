package hull

type Point [2] float64

type Points []Point

type Grid struct {
	Cells    Points
	CellSize float64
}

func (slice Points) include(value int) int {
	for _, v := range slice {
		if (v == value) {
			return true
		}
	}
	return false
}

func Filter(vs Points) Points {
	vsf := make([]string, 0)
	for _, v := range vs {
		if vs.include(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func (points Points) Swap(i, j int) {
	points[i], points[j] = points[j], points[i]
}

func (points Points) Len() int {
	return len(points)
}

// lets sort our Points by x and, if equal, by y
func (points Points) Less(i, j int) bool {
	if points[i][0] == points[j][0] {
		return points[i][1] < points[j][1]
	}
	return points[i][0] < points[j][0]
}

func (grid Grid) rangePoints(bbox [4]float64) Points {
	// (Array) -> Array
	tlCellXY := grid.point2CellXY([2]float64{bbox[0], bbox[1]})
	brCellXY := grid.point2CellXY([2]float64{bbox[2], bbox[3]})
	var points Points;

	for x := tlCellXY[0]; x <= brCellXY[0]; x++ {
		for y := tlCellXY[1]; y <= brCellXY[1]; y++ {
			points = append(points, grid.Cells[x][y]);
		}
	}

	return points;
}
//
//removePoint: func (point) {
//	// (Array) -> Array
//	var cellXY = this.point2CellXY(point),
//		cell = this._cells[cellXY[0]][cellXY[1]],
//		pointIdxInCell;
//
//	for
//	(
//	var i = 0; i < cell.length; i++) {
//	if (cell[i][0] === point[0] && cell[i][1] == = point[1]) {
//	pointIdxInCell = i;
//	break;
//	}
//	}
//
//	cell.splice(pointIdxInCell, 1);
//
//	return cell;
//},
//
func (grid Grid) point2CellXY(point [2]float64) [2]float64 {
	// (Array) -> Array
	x := (point[0] / grid.CellSize)
	y := (point[1] / grid.CellSize);
	return [2]float64{x, y};
}

func (grid Grid) extendBbox(bbox [4]float64, scaleFactor int) [4]float64 {
	return [4]float64{bbox[0] - (scaleFactor * grid.CellSize), bbox[1] - (scaleFactor * grid.CellSize), bbox[2] + (scaleFactor * grid.CellSize), bbox[3] + (scaleFactor * grid.CellSize)};
}