package mbr

import "github.com/intdxdt/math"

func distanceDxdy(self *MBR, other *MBR) (float64, float64) {
	dx, dy := 0.0, 0.0
	//find closest edge by x
	if self[x2] < other[x1] {
		dx = other[x1] - self[x2]
	} else if self[x1] > other[x2] {
		dx = self[x1] - other[x2]
	}
	//find closest edge by y
	if self[y2] < other[y1] {
		dy = other[y1] - self[y2]
	} else if self[y1] > other[y2] {
		dy = self[y1] - other[y2]
	}
	return dx, dy
}

//Distance computes the distance between two mbrs
func (mbr *MBR) Distance(other *MBR) float64 {
	if mbr.Intersects(other) {
		return 0.0
	}
	dx, dy := distanceDxdy(mbr, other)
	return math.Hypot(dx, dy)
}

//DistanceSquare computes the distance squared between mbrs
func (mbr *MBR) DistanceSquare(other *MBR) float64 {
	if mbr.Intersects(other) {
		return 0.0
	}
	dx, dy := distanceDxdy(mbr, other)
	return (dx * dx) + (dy * dy)
}
