package rect

type Rect struct {
	X1, Y1, X2, Y2 int
}

func NewRect(x1, y1, x2, y2 int) *Rect {
	return &Rect{
		X1: x1,
		Y1: y1,
		X2: x2,
		Y2: y2,
	}
}

// GetExtremes returns the minimum and maximum X and Y coordinates of the rectangle.
func (r *Rect) GetExtremes() (minX, maxX, minY, maxY int) {
	minX, maxX = min(r.X1, r.X2), max(r.X1, r.X2)
	minY, maxY = min(r.Y1, r.Y2), max(r.Y1, r.Y2)

	return minX, maxX, minY, maxY
}

// AreaGrid calculates the area of the rectangle in grid units.
func (r *Rect) AreaGrid() int {
	minX, maxX, minY, maxY := r.GetExtremes()
	width := maxX - minX + 1
	height := maxY - minY + 1
	return height * width
}

// GetCorners returns the four corners of the rectangle as a slice of (x, y) coordinate pairs.
//
// The corners are returned in in clockwise order starting from the top-left corner (assuming 0,0 is top-left).
func (r *Rect) GetCorners() [][2]int {
	minX, maxX, minY, maxY := r.GetExtremes()
	return [][2]int{
		{minX, minY},
		{maxX, minY},
		{maxX, maxY},
		{minX, maxY},
	}
}
