package concaveHull

type Point [2]float64

type Points []Point

type Grid struct {
	Cells    map[int]map[int]Points
	CellSize float64
}

var EmptyPoint = Point{}

func createGrid(points Points, cellSize float64) Grid {

	var grid = Grid{Cells: make(map[int]map[int]Points), CellSize: cellSize}

	for _, point := range points {
		var cellXY = grid.point2CellXY(point)
		x := cellXY[0]
		y := cellXY[1]
		if _, ok := grid.Cells[x]; !ok {
			grid.Cells[x] = make(map[int]Points)
		}
		grid.Cells[x][y] = append(grid.Cells[x][y], point)
	}

	return grid
}

func (slice Points) include(value Point) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func Filter(search, points Points) Points {
	vsf := make(Points, 0)
	for _, v := range search {
		if !points.include(v) {
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

func (points Points) Less(i, j int) bool {
	if points[i][0] == points[j][0] {
		return points[i][1] < points[j][1]
	}
	return points[i][0] < points[j][0]
}

func splice(points Points, index, amount int, elements ...Point) Points {
	newslice := make(Points, 0)
	for i := 0; i < index; i++ {
		newslice = append(newslice, points[i])
	}
	for _, el := range elements {
		newslice = append(newslice, el)
	}
	for i := index + amount; i < len(points); i++ {
		newslice = append(newslice, points[i])
	}

	return newslice
}

func (points Points) reverse() Points {
	for i, j := 0, len(points)-1; i < j; i, j = i+1, j-1 {
		points[i], points[j] = points[j], points[i]
	}
	return points
}

func (grid Grid) rangePoints(bbox [4]float64) Points {
	tlCellXY := grid.point2CellXY([2]float64{bbox[0], bbox[1]})
	brCellXY := grid.point2CellXY([2]float64{bbox[2], bbox[3]})
	var points Points

	for x := tlCellXY[0]; x <= brCellXY[0]; x++ {
		for y := tlCellXY[1]; y <= brCellXY[1]; y++ {
			points = append(points, grid.Cells[x][y]...)
		}
	}
	return points
}

func (grid Grid) removePoint(point Point) Grid {
	var cellXY = grid.point2CellXY(point)
	cell := grid.Cells[cellXY[0]][cellXY[1]]
	var pointIdxInCell int

	for i := 0; i < len(cell); i++ {
		if cell[i][0] == point[0] && cell[i][1] == point[1] {
			pointIdxInCell = i
			break
		}
	}
	grid.Cells[cellXY[0]][cellXY[1]] = splice(grid.Cells[cellXY[0]][cellXY[1]], pointIdxInCell, 1)

	return grid
}

func (grid Grid) point2CellXY(point [2]float64) [2]int {
	x := int(point[0] / grid.CellSize)
	y := int(point[1] / grid.CellSize)
	return [2]int{x, y}
}

func (grid Grid) extendBbox(bbox [4]float64, scaleFactor float64) [4]float64 {
	return [4]float64{bbox[0] - (scaleFactor * grid.CellSize), bbox[1] - (scaleFactor * grid.CellSize), bbox[2] + (scaleFactor * grid.CellSize), bbox[3] + (scaleFactor * grid.CellSize)}
}
